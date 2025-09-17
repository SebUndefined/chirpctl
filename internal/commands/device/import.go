package device

import (
	"encoding/csv"
	"fmt"
	"strings"

	"github.com/sebundefined/chirpctl/internal/app"
	"github.com/sebundefined/chirpctl/internal/commands/cmdutils"
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
			r, closer, err := cmdutils.OpenReader(importFlags.File) // stdin/fichier
			if err != nil {
				return err
			}
			defer closer()
			csvReader := csv.NewReader(r)
			csvReader.Comma = ';'
			records, err := csvReader.ReadAll()
			if err != nil {
				return fmt.Errorf("Unable to parse file as CSV for "+importFlags.File, err)
			}
			start := 0
			if !importFlags.SkipHeader && len(records) > 0 {
				// Vérifie si la première ligne ressemble à un header
				if strings.EqualFold(records[0][0], "name") {
					start = 1
				}
			}
			var devices []device.CreateParams
			for _, record := range records[start:] {
				d := device.CreateParams{
					Device: device.Device{
						Name:            record[0],
						Description:     record[1],
						DevEUI:          record[2],
						JoinEUI:         record[3],
						ApplicationID:   record[4],
						DeviceProfileID: record[5],
						SkipFcntCheck:   record[6] == "true",
						IsDisabled:      record[7] == "true",
						Variables:       make(map[string]string),
					},
				}
				varsString := strings.TrimSpace(record[8])
				if len(varsString) > 0 {
					vars := strings.Split(varsString, ",")
					for _, v := range vars {
						keyValue := strings.Split(v, "=")
						if len(keyValue) != 2 {
							continue
						}
						d.Device.Variables[keyValue[0]] = keyValue[1]
					}
				}
				devices = append(devices, d)
			}
			fmt.Println(devices)
			fmt.Printf("Importing %d devices...\n", len(devices))
			return nil
		},
	}

	cmdutils.AttachImportFlags(cmd, &importFlags)
	return cmd
}
