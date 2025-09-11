package printing

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"gopkg.in/yaml.v3"
)

type Row []any

type TableData struct {
	Headers Row
	Rows    []Row
}

type Options struct {
	Format  Format
	Wide    bool
	NoColor bool
}

func RenderTableStyled(td TableData, noColor bool) {
	if noColor {
		text.DisableColors()
	}
	tw := table.NewWriter()
	tw.SetOutputMirror(os.Stdout)
	tw.SetStyle(table.StyleLight)
	tw.Style().Color.Header = text.Colors{text.Bold}
	tw.AppendHeader(table.Row(td.Headers))
	for _, r := range td.Rows {
		tw.AppendRow(table.Row(r))
	}
	tw.Render()
}

func RenderTableRaw(td TableData) {
	widths := make([]int, len(td.Headers))
	rows := append([]Row{td.Headers}, td.Rows...)
	for _, r := range rows {
		for i, cell := range r {
			w := len(fmt.Sprint(cell))
			if w > widths[i] {
				widths[i] = w
			}
		}
	}
	for i, h := range td.Headers {
		fmt.Printf("%-*s", widths[i], fmt.Sprint(h))
		if i < len(td.Headers)-1 {
			fmt.Print("  ")
		}
	}
	fmt.Println()
	for _, r := range td.Rows {
		for i, cell := range r {
			fmt.Printf("%-*s", widths[i], fmt.Sprint(cell))
			if i < len(r)-1 {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

func RenderJSON(v any) error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}

func RenderYAML(v any) error {
	b, err := yaml.Marshal(v)
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write(b)
	return err
}

func RenderCSV(td TableData) error {
	w := csv.NewWriter(os.Stdout)
	// headers
	_ = w.Write(toStrings(td.Headers))
	for _, r := range td.Rows {
		_ = w.Write(toStrings(r))
	}
	w.Flush()
	return w.Error()
}

func toStrings(r Row) []string {
	out := make([]string, len(r))
	for i, c := range r {
		out[i] = fmt.Sprint(c)
	}
	return out
}

func RenderList(format Format, td TableData, list any, noColor bool) error {
	var err error
	switch format {
	case FormatTableStyled:
		RenderTableStyled(td, noColor)
	case FormatTable:
		RenderTableRaw(td)
	case FormatJSON:
		err = RenderJSON(list)
	case FormatYAML:
		err = RenderYAML(list)
	case FormatCSV:
		err = RenderCSV(td)
	}
	return err
}
