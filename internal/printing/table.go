package printing

import (
	"fmt"
	"strings"
)

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

// AppendSection adds a section heading to a detail table.
func AppendSection(td *TableData, name string, styled, useColor bool) {
	label := strings.ToUpper(strings.TrimSpace(name))

	if styled && useColor {
		label = StatusSymbol(label, ColorBlue)
	}

	td.Rows = append(td.Rows, Row{label, ""})
}
