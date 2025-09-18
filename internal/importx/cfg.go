package importx

import (
	"runtime"
	"time"
)

// RunnerConfig -
type RunnerConfig struct {
	// Max number of concurrent workers.
	// For CPU-bound tasks: keep ≈ runtime.NumCPU().
	// For IO-bound tasks: higher values can help, but too high may cause contention and memory overhead.
	Workers int
	// Maximum number of jobs buffered in the queue
	// before Submit() blocks or returns an error.
	// Acts as backpressure to protect the system.
	QueueSize int

	// Maximum duration allowed per job execution.
	// Each job runs with its own context.WithTimeout().
	Timeout time.Duration

	// Number of times a failed job should be retried
	// before being marked as permanently failed.
	// Retries are attempted immediately.
	MaxRetries int

	// Target rate limit, expressed in jobs per second.
	// If <= 0, no rate limiting is applied.
	RatePerSecond int

	// Maximum burst of jobs allowed at once (token bucket capacity).
	// If <= 0, defaults to min(RatePerSecond, 2*Workers).
	Burst int

	// Minimum time interval between token emissions.
	// Prevents extremely small periods (e.g. >1k/sec).
	// Defaults to 1ms.
	MinPeriod time.Duration
}

func (rc *RunnerConfig) normalize() {
	if rc.Workers <= 0 {
		rc.Workers = 1
	}
	if rc.QueueSize <= 0 {
		rc.QueueSize = 256
	}
	if rc.Timeout <= 0 {
		rc.Timeout = 3 * time.Second
	}
	if rc.MaxRetries < 0 {
		rc.MaxRetries = 0
	}
	if rc.MinPeriod <= 0 {
		rc.MinPeriod = time.Millisecond
	}
	if rc.Burst <= 0 {
		if rc.RatePerSecond > 0 {
			rc.Burst = min(rc.RatePerSecond, 2*rc.Workers)
		} else {
			rc.Burst = 2 * rc.Workers
		}
	}

}

var DefaultRunnerConfig = RunnerConfig{
	Workers:       runtime.NumCPU(),
	QueueSize:     256,
	Timeout:       10 * time.Second,
	MaxRetries:    3,
	RatePerSecond: 10,
}

type RunnerOption func(*RunnerConfig)

func WithWorkers(workers int) RunnerOption {
	return func(cfg *RunnerConfig) {
		cfg.Workers = workers
	}
}

func WithQueueSize(queueSize int) RunnerOption {
	return func(cfg *RunnerConfig) {
		cfg.QueueSize = queueSize
	}
}

func WithTimeout(timeout time.Duration) RunnerOption {
	return func(cfg *RunnerConfig) {
		cfg.Timeout = timeout
	}
}
func WithMaxRetries(maxRetries int) RunnerOption {
	return func(cfg *RunnerConfig) {
		cfg.MaxRetries = maxRetries
	}
}

func WithRatePerSecond(ratePerSecond int) RunnerOption {
	return func(cfg *RunnerConfig) {
		cfg.RatePerSecond = ratePerSecond
	}
}

func WithBurst(burst int) RunnerOption {
	return func(cfg *RunnerConfig) {
		cfg.Burst = burst
	}
}

func WithMinPeriod(minPeriod time.Duration) RunnerOption {
	return func(cfg *RunnerConfig) {
		cfg.MinPeriod = minPeriod
	}
}
