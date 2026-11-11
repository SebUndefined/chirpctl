package entityimport

import (
	"github.com/sebundefined/chirpctl/internal/app"
	"github.com/spf13/cobra"
)

func NewImportCmd(a *app.App) *cobra.Command {
	var file, typ, format string
	var upsert, cont bool
	var batch int
	var dryRun bool

	cmd := &cobra.Command{
		Use:   "import",
		Short: "Import resources from CSV/JSON/YAML",
		RunE: func(cmd *cobra.Command, args []string) error {
			// 1) ouvrir/decoder (stdin si file == "-")
			// 2) déterminer type (flag, header, ou contenu)
			// 3) valider & batcher
			// 4) appeler les services via a.Factory (Device/Gateway…)
			// 5) afficher un résumé + code de sortie approprié
			return nil
		},
	}
	cmd.Flags().StringVarP(&file, "file", "f", "", "Input file ('-' for stdin)")
	cmd.Flags().StringVar(&typ, "type", "", "Resource type (device|gateway|...)")
	cmd.Flags().StringVar(&format, "format", "", "Input format (csv|json|yaml)")
	cmd.Flags().BoolVar(&upsert, "upsert", false, "Create or update if exists")
	cmd.Flags().BoolVar(&cont, "continue-on-error", false, "Do not stop on first error")
	cmd.Flags().IntVar(&batch, "batch-size", 100, "Batch size")
	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "Validate only")
	_ = cmd.MarkFlagRequired("file")
	return cmd
}
