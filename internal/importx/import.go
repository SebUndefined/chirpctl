package importx

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Job represent a job that needs to be done
type Job[T any] struct {
	Line int
	Item T
}

// DoFunc f
type DoFunc[T any] func(ctx context.Context, j Job[T]) error

// RunnerResult is a part of the result
type RunnerResult struct {
	Err  error
	Line int
}

type Runner[T any] struct {
	cfg     RunnerConfig
	do      DoFunc[T]
	jobs    chan Job[T]
	results chan RunnerResult
	wg      sync.WaitGroup
	ctx     context.Context
	cancel  context.CancelFunc
	tokens  chan struct{}
	ticker  *time.Ticker
	// Once
	startOnce        sync.Once
	closeJobsOnce    sync.Once
	closeResultsOnce sync.Once
	stopOnce         sync.Once
}

func NewRunner[T any](parentCtx context.Context, do DoFunc[T], opts ...RunnerOption) *Runner[T] {
	cfg := DefaultRunnerConfig
	for _, opt := range opts {
		opt(&cfg)
	}
	cfg.normalize()
	ctx, cancel := context.WithCancel(parentCtx)
	return &Runner[T]{
		cfg:     cfg,
		do:      do,
		jobs:    make(chan Job[T], cfg.QueueSize),
		results: make(chan RunnerResult, min(cfg.QueueSize, 1024)),
		wg:      sync.WaitGroup{},
		ctx:     ctx,
		cancel:  cancel,
	}
}

func (r *Runner[T]) Start() {
	r.startOnce.Do(func() {
		// Add all worker to waigroup
		r.wg.Add(r.cfg.Workers)
		// In case the rate limiter is enabled, we start it. See startRateLimiter for more details
		if r.cfg.RatePerSecond > 0 {
			r.startRateLimiter(r.ctx)

		}
		for i := 0; i < r.cfg.Workers; i++ {
			go r.startWorker()
		}
	})
}

func (r *Runner[T]) startRateLimiter(ctx context.Context) {
	rate := r.cfg.RatePerSecond
	if rate <= 0 {
		return
	}
	period := time.Second / time.Duration(rate)
	if period < r.cfg.MinPeriod {
		period = r.cfg.MinPeriod
	}

	r.tokens = make(chan struct{}, r.cfg.Burst)
	r.ticker = time.NewTicker(period)
	accumulator := 0.0
	tokensPerSecond := float64(rate)
	maxAccumulator := float64(cap(r.tokens))
	go func() {
		defer r.ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-r.ticker.C:
				// How many tokens per tick ?
				accumulator += tokensPerSecond * period.Seconds()
				// Max accumulator cannot be upper that tokens capacity
				if accumulator > maxAccumulator {
					accumulator = maxAccumulator
				}
				// emit only the integer part
				emit := int(accumulator)
				if emit <= 0 {
					continue
				}
				emitted := 0
			emitLoop:
				for i := 0; i < emit; i++ {
					select {
					case r.tokens <- struct{}{}:
						emitted++
					default:
						break emitLoop
					}
				}
				accumulator -= float64(emitted)
				if accumulator < 0 {
					accumulator = 0
				}
			}
		}
	}()
}

func (r *Runner[T]) startWorker() {
	defer r.wg.Done()
	for {
		select {
		case job, ok := <-r.jobs:
			if !ok {
				return // Job closed, worker ended
			}
			result, canceled := r.tryDo(r.ctx, job)
			if canceled {
				return
			}
			select {
			case r.results <- result: // Try to send result on channel
			case <-r.ctx.Done(): // if stop has been called
				return
			}
		case <-r.ctx.Done():
			return
		}
	}
}

func (r *Runner[T]) tryDo(ctx context.Context, job Job[T]) (RunnerResult, bool) {
	result := RunnerResult{
		Line: job.Line,
	}

	defer func() {
		if rec := recover(); rec != nil {
			result.Err = fmt.Errorf("panic: %v", rec)
		}
	}()

	// rate-limit (with ctx)
	if r.cfg.RatePerSecond > 0 {
		select {
		case <-r.tokens:
		case <-ctx.Done():
			result.Err = ctx.Err()
			return result, true
		}
	}

	// timeout per job
	ctxJob, cancel := context.WithTimeout(ctx, r.cfg.Timeout)
	err := r.do(ctxJob, job)
	cancel()

	result.Err = err
	return result, false
}

func (r *Runner[T]) Submit(j Job[T]) error {
	select {
	case r.jobs <- j:
		return nil
	case <-r.ctx.Done():
		return r.ctx.Err()
	}
}

func (r *Runner[T]) Emit(res RunnerResult) {
	select {
	case r.results <- res:
	case <-r.ctx.Done():
	}
}

func (r *Runner[T]) CloseJobs() {
	r.closeJobsOnce.Do(func() {
		close(r.jobs)
	})
}

func (r *Runner[T]) Results() <-chan RunnerResult {
	return r.results
}

func (r *Runner[T]) Wait() {
	r.wg.Wait()
	r.closeResultsOnce.Do(func() { close(r.results) })
}

func (r *Runner[T]) Stop() {
	r.stopOnce.Do(func() {
		r.cancel()
	})
}
