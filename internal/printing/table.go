package printing

import "fmt"

type Row []any

func (r Row) toStrings() []string {
	out := make([]string, len(r))
	for i, c := range r {
		out[i] = fmt.Sprint(c)
	}
	return out
}

type TableData struct {
	Headers Row
	Rows    []Row
}
