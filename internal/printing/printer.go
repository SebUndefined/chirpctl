package printing

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"gopkg.in/yaml.v3"
)

type Options struct {
	Format  Format
	Wide    bool
	NoColor bool
}

// Printer renders CLI results using the selected output format.
type Printer struct {
	out  io.Writer
	opts Options
}

// NewPrinter creates a printer with an explicit output destination.
func NewPrinter(out io.Writer, opts Options) *Printer {
	return &Printer{
		out:  out,
		opts: opts,
	}
}

func (p *Printer) Format() Format {
	return p.opts.Format
}

// UseColor reports whether colored output is enabled.
func (p *Printer) UseColor() bool {
	return p.opts.Format == FormatTableStyled && !p.opts.NoColor
}

// Print renders structured data or its tabular representation.
func (p *Printer) Print(data any, td TableData) error {
	var err error
	switch p.opts.Format {
	case FormatTableStyled:
		return p.renderTableStyled(td)
	case FormatTable:
		return p.renderTableRaw(td)
	case FormatJSON:
		err = p.renderJSON(data)
	case FormatYAML:
		err = p.renderYAML(data)
	case FormatCSV:
		err = p.renderCSV(td)
	default:
		return fmt.Errorf("unsupported output format: %s", p.opts.Format)
	}
	return err
}

func (p *Printer) renderJSON(data any) error {
	enc := json.NewEncoder(p.out)
	enc.SetIndent("", "  ")
	return enc.Encode(data)
}
func (p *Printer) renderYAML(data any) error {
	enc := yaml.NewEncoder(p.out)
	enc.SetIndent(2)
	if err := enc.Encode(data); err != nil {
		return err
	}
	return enc.Close()
}
func (p *Printer) renderCSV(td TableData) error {
	w := csv.NewWriter(p.out)
	// Write the CSV header and propagate encoding errors.
	if err := w.Write(td.Headers.toStrings()); err != nil {
		return err
	}
	for _, r := range td.Rows {
		if err := w.Write(r.toStrings()); err != nil {
			return err
		}
	}
	w.Flush()
	return w.Error()
}

func (p *Printer) renderTableStyled(td TableData) error {
	// Create an independent table writer.
	tw := table.NewWriter()
	tw.SetStyle(table.StyleLight)

	// Configure colors for this table only.
	tw.Style().Color = table.ColorOptionsDefault

	if !p.opts.NoColor {
		tw.Style().Color.Header = text.Colors{text.Bold}
	}
	tw.AppendHeader(table.Row(td.Headers))
	for _, r := range td.Rows {
		tw.AppendRow(table.Row(r))
	}
	// Write the rendered table to the configured destination.
	_, err := io.WriteString(p.out, tw.Render()+"\n")
	return err
}

// renderTableRaw renders a plain-text table without borders or ANSI styling.
// All output is written to the Printer's configured writer.
func (p *Printer) renderTableRaw(td TableData) error {
	// Calculate column widths based on headers and rows.
	widths := make([]int, len(td.Headers))
	rows := append([]Row{td.Headers}, td.Rows...)
	for _, row := range rows {
		for i, cell := range row {
			width := len(fmt.Sprint(cell))
			if width > widths[i] {
				widths[i] = width
			}
		}
	}
	// Write a row using the calculated column widths.
	writeRow := func(row Row) error {
		for i, cell := range row {
			if i > 0 {
				if _, err := fmt.Fprint(p.out, "  "); err != nil {
					return err
				}
			}
			if _, err := fmt.Fprintf(p.out, "%-*s", widths[i], fmt.Sprint(cell)); err != nil {
				return err
			}
		}
		_, err := fmt.Fprintln(p.out)
		return err
	}
	// Write table headers.
	if err := writeRow(td.Headers); err != nil {
		return err
	}
	// Write table rows.
	for _, row := range td.Rows {
		if err := writeRow(row); err != nil {
			return err
		}
	}
	return nil
}
