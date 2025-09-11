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
	switch s {
	case string(FileFormatJSON):
		return FileFormatJSON
	case string(FileFormatYAML):
		return FileFormatYAML
	case string(FileFormatCsv):
		return FileFormatCsv
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
	File   string
	Format string
	DryRun bool
	UpSert bool
	Batch  int32
}

func AttachImportFlags(cmd *cobra.Command, opts *ImportOptions) {
	cmd.Flags().StringVar(&opts.File, "file", "", "The file to import.")
	cmd.Flags().StringVar(&opts.Format, "format", "csv", "The format to use for importing.")
	cmd.Flags().BoolVar(&opts.DryRun, "dry-run", false, "Whether or not to actually import the files.")
	cmd.Flags().BoolVar(&opts.UpSert, "upsert", false, "Whether or not to upsert the files.")
	cmd.Flags().Int32Var(&opts.Batch, "batch", 100, "Size of each import to import.")
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
