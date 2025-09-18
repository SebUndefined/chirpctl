package device

import (
	"fmt"
	"io"

	"github.com/sebundefined/chirpctl/internal/app"
	"github.com/sebundefined/chirpctl/internal/commands/cmdutils"
	"github.com/sebundefined/chirpctl/internal/csvx"
	"github.com/sebundefined/chirpctl/internal/iox"
	"github.com/sebundefined/chirpctl/internal/service/device"
	"github.com/spf13/cobra"
)

func NewImportCommand(a *app.App) *cobra.Command {
	var importFlags cmdutils.ImportOptions
	var cmd = &cobra.Command{
		Use:   "import",
		Short: "Import devices",
		Long:  `import allows you to import devices from a given file (csv, json or yaml). `,
		RunE: func(cmd *cobra.Command, args []string) error {
			r, closer, err := iox.OpenReader(importFlags.File) // stdin/fichier
			if err != nil {
				return err
			}
			defer closer()
			csvReader := csvx.NewReader(r,
				csvx.WithHasHeader(true),
				csvx.WithComma(';'),
				csvx.WithTrimSpace(true))
			_, err = csvReader.Header()
			if err != nil {
				return err
			}
			indexes, err := csvReader.BuildColumnIndex()
			if err != nil {
				return err
			}
			specs, err := device.CompileDeviceCsvSpecs(indexes)
			if err != nil {
				return err
			}
			var devices []device.CreateParams
			for {
				row, idx, err := csvReader.Next()
				if err == io.EOF {
					break
				} else if err != nil {
					return err
				}
				csvSpecs, err := device.DecodeDeviceCsvSpecs(row, specs)
				if err != nil {
					return fmt.Errorf("error decoding CSV specs at line %d: %w", idx, err)
				}
				devices = append(devices, device.CreateParams{
					Device: csvSpecs,
				})
			}
			fmt.Println(devices)
			fmt.Printf("Importing %d devices...\n", len(devices))
			return nil
		},
	}

	cmdutils.AttachImportFlags(cmd, &importFlags)
	return cmd
}
