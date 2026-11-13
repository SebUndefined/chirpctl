package printing

import (
	"fmt"
	"maps"
	"slices"
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

// AppendStringMapRows appends map entries in alphabetical order.
func AppendStringMapRows(td *TableData, prefix string, values map[string]string) {
	if len(values) == 0 {
		td.Rows = append(td.Rows, Row{prefix, "-"})
		return
	}

	for _, key := range slices.Sorted(maps.Keys(values)) {
		td.Rows = append(td.Rows, Row{
			prefix + "." + key,
			values[key],
		})
	}
}
