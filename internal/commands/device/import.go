package device

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/sebundefined/chirpctl/internal/app"
	"github.com/sebundefined/chirpctl/internal/commands/cmdutils"
	"github.com/sebundefined/chirpctl/internal/csvx"
	"github.com/sebundefined/chirpctl/internal/importx"
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
			if importFlags.SkipHeader {
				return fmt.Errorf("--skip-header is not supported for device CSV import: header row is required")
			}
			csvReader := csvx.NewReader(r,
				csvx.WithHasHeader(true),
				csvx.WithComma(';'),
				csvx.WithTrimSpace(true))
			hdr, err := csvReader.Header()
			if err != nil {
				return err
			}
			fmt.Println("HEADER =", hdr)
			indexes, err := csvReader.BuildColumnIndex()
			if err != nil {
				return err
			}
			specs, err := device.CompileDeviceCsvSpecs(indexes)
			if err != nil {
				return err
			}
			svc, err := a.Factory.Device()

			if err != nil {
				return err
			}
			var submittedCount int
			runner := importx.NewRunner[device.ImportParams](cmd.Context(), func(ctx context.Context, j importx.Job[device.ImportParams]) error {
				// CHANGED: on ne swallow plus l'erreur Create.
				// Si la création échoue, il faut remonter l'erreur au runner.
				_, errCreate := svc.Create(ctx, device.CreateParams{
					Device: j.Item.Device,
				})
				if errCreate != nil {
					return fmt.Errorf("create device %s: %w", j.Item.Device.DevEUI, errCreate)
				}

				// CHANGED: si des clés existent, on tente de les créer.
				if j.Item.Keys != nil {
					if err := svc.CreateKeys(ctx, *j.Item.Keys); err != nil {
						return fmt.Errorf("create keys for %s: %w", j.Item.Device.DevEUI, err)
					}
				}

				return nil
			}, importx.WithWorkers(8),
				importx.WithQueueSize(512),
				importx.WithTimeout(10*time.Second),
				importx.WithRatePerSecond(100), // commence prudent; monte si OK
				importx.WithBurst(32),
				importx.WithMinPeriod(1*time.Millisecond))
			runner.Start()
			defer runner.Stop()
			var (
				okCount  int
				errCount int
			)
			done := make(chan struct{})
			go func() {
				for res := range runner.Results() {
					if res.Err != nil {
						errCount++
						// log court et utile (ligne CSV + err)
						fmt.Fprintf(cmd.ErrOrStderr(), "ligne %d: %v\n", res.Line, res.Err)
					} else {
						okCount++
					}
				}
				close(done)
			}()
			for {
				row, idx, err := csvReader.Next()
				if err == io.EOF {
					fmt.Println("EOF reached")
					break
				} else if err != nil {
					return err
				}
				fmt.Println("ROW", idx, row)
				csvSpecs, err := device.DecodeDeviceImportCsvSpecs(row, specs)
				if err != nil {
					return fmt.Errorf("error decoding CSV specs at line %d: %w", idx, err)
				}
				fmt.Println("Submitting", idx, csvSpecs.Device.DevEUI)
				if err := runner.Submit(importx.Job[device.ImportParams]{
					Line: idx,
					Item: csvSpecs,
				}); err != nil {
					return fmt.Errorf("submit line %d: %w", idx, err)
				}
				// CHANGED: on compte les jobs réellement soumis
				submittedCount++
			}
			runner.CloseJobs() // plus de jobs à venir
			runner.Wait()      // attend la fin des workers + ferme results
			<-done             // le collecteur a fini de vider results

			fmt.Fprintf(
				cmd.OutOrStdout(),
				"Import terminé: %d soumis, %d ok, %d erreurs\n",
				submittedCount,
				okCount,
				errCount,
			)
			return nil
		},
	}

	cmdutils.AttachImportFlags(cmd, &importFlags)
	return cmd
}
