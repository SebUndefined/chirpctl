package csvx

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strings"
)

type ColumnIndex map[string]int

type Config struct {
	Comma     rune
	HasHeader bool
	TrimSpace bool
}

var defaultConfig = Config{
	Comma:     ';',
	HasHeader: true,
	TrimSpace: true,
}

type Option func(*Config)

func WithComma(c rune) Option {
	return func(config *Config) {
		config.Comma = c
	}
}

func WithTrimSpace(trimSpace bool) Option {
	return func(config *Config) {
		config.TrimSpace = trimSpace
	}
}

func WithHasHeader(hasHeader bool) Option {
	return func(config *Config) {
		config.HasHeader = hasHeader
	}
}

type Reader struct {
	reader  *csv.Reader
	config  Config
	headers []string
	line    int
}

func NewReader(reader io.Reader, opts ...Option) *Reader {
	cfg := defaultConfig
	for _, opt := range opts {
		opt(&cfg)
	}
	csvReader := csv.NewReader(reader)
	csvReader.Comma = cfg.Comma
	csvReader.TrimLeadingSpace = cfg.TrimSpace
	return &Reader{
		reader: csvReader,
		config: cfg,
	}
}

// Header get the header form csv file
func (r *Reader) Header() ([]string, error) {
	if r.headers != nil {
		return r.headers, nil
	}
	hdr, err := r.reader.Read()
	if err != nil {
		return nil, err
	}
	if len(hdr) > 0 {
		hdr[0] = strings.TrimPrefix(hdr[0], "\uFEFF")
	}
	r.headers = hdr
	r.line++
	r.reader.FieldsPerRecord = len(hdr)
	return hdr, nil
}

func (r *Reader) Next() ([]string, int, error) {
	if len(r.headers) == 0 {
		return nil, 0, errors.New("call Header() first (header line is required)")
	}
	row, err := r.reader.Read()
	if err != nil {
		if err == io.EOF {
			return nil, 0, io.EOF
		}
		return nil, 0, fmt.Errorf("line %d: %w", r.line+1, err)
	}
	r.line++
	return row, r.line, nil
}

func (r *Reader) BuildColumnIndex() (ColumnIndex, error) {
	if len(r.headers) == 0 {
		return nil, errors.New("call Header() first (header line is required)")
	}
	idx := make(ColumnIndex, len(r.headers))
	for i, name := range r.headers {
		if len(strings.TrimSpace(name)) == 0 {
			return nil, fmt.Errorf("column index must have a non-empty column name")
		}
		if _, ok := idx[name]; ok {
			return nil, fmt.Errorf("column index: duplicate column name: %s", name)
		}
		idx[name] = i
	}
	return idx, nil
}
