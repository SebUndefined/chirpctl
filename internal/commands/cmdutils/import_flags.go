package cmdutils

import (
	"strings"

	"github.com/spf13/cobra"
)

type FileFormat string

const (
	FileFormatJSON    FileFormat = "JSON"
	FileFormatYAML    FileFormat = "YAML"
	FileFormatCsv     FileFormat = "CSV"
	FileFormatUnknown FileFormat = "UNKNOWN"
)

func ParseFileFormat(s string) FileFormat {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "csv":
		return FileFormatCsv
	case "json":
		return FileFormatJSON
	case "yaml", "yml":
		return FileFormatYAML
	default:
		return FileFormatUnknown
	}
}

func (f FileFormat) String() string {
	if len(strings.TrimSpace(string(f))) == 0 {
		return FileFormatUnknown.String()
	}
	return string(f)
}

type ImportOptions struct {
	File         string
	Format       string
	DryRun       bool
	UpSert       bool
	Batch        int32
	SkipHeader   bool
	CsvDelimiter string
}

func AttachImportFlags(cmd *cobra.Command, opts *ImportOptions) {
	cmd.Flags().StringVar(
		&opts.File, "file", "",
		"Input file path (- for stdin)",
	)
	cmd.Flags().StringVar(
		&opts.Format, "format", "csv",
		"Input format: csv (JSON and YAML planned)",
	)
	cmd.Flags().BoolVar(
		&opts.DryRun, "dry-run", false,
		"Validate the input without creating resources",
	)
	cmd.Flags().BoolVar(
		&opts.UpSert, "upsert", false,
		"Update existing resources (not yet implemented)",
	)
	cmd.Flags().Int32Var(
		&opts.Batch, "batch", 100,
		"Batch size (not yet implemented)",
	)
	cmd.Flags().BoolVar(
		&opts.SkipHeader, "skip-header", false,
		"Skip the CSV header (unsupported for devices)",
	)
	cmd.Flags().StringVar(
		&opts.CsvDelimiter, "csv-delim", ";",
		"CSV field delimiter",
	)
}

type ExportFlags struct {
	Format FileFormat
	Output string // fichier, "-" = stdout
	// filtres communs
	TenantID string
	Search   string
	Limit    uint32
	Offset   uint32
}

func AttachExportFlags(cmd *cobra.Command, f *ExportFlags) { /* add flags */ }

type ApplyFlags struct {
	File   string // yaml/json, peut contenir plusieurs docs ---
	DryRun bool
	Prune  bool // supprimer ce qui n'est pas dans le fichier (scope/label)
	Force  bool
}

func AttachApplyFlags(cmd *cobra.Command, f *ApplyFlags) { /* add flags */ }
