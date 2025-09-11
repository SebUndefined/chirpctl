package device

import (
	"encoding/csv"
	"fmt"

	"github.com/sebundefined/chirpctl/internal/app"
	"github.com/sebundefined/chirpctl/internal/commands/cmdutils"
	"github.com/spf13/cobra"
)

func NewImportCommand(a *app.App) *cobra.Command {
	var importFlags cmdutils.ImportOptions
	var cmd = &cobra.Command{
		Use:   "import",
		Short: "Import devices",
		Long:  `import allows you to import devices from a given file (csv, json or yaml). `,
		RunE: func(cmd *cobra.Command, args []string) error {

			fmt.Println(importFlags.Format)
			r, closer, err := cmdutils.OpenReader(importFlags.File) // stdin/fichier
			if err != nil {
				return err
			}
			defer closer()
			csvReader := csv.NewReader(r)
			records, err := csvReader.ReadAll()
			if err != nil {
				return fmt.Errorf("Unable to parse file as CSV for "+importFlags.File, err)
			}
			fmt.Println(records)
			//l := do.MustInvokeNamed[logger.Interface]("logger")
			return nil
		},
	}

	cmdutils.AttachImportFlags(cmd, &importFlags)
	return cmd
}
