package device

import (
	"context"
	"fmt"
	"io"
	"strings"
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
	var examplePath string

	var cmd = &cobra.Command{
		Use:     "devices",
		Aliases: []string{"device"},
		Short:   "Import devices",
		Long:    `import allows you to import devices from a given file (csv, json or yaml). `,
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			format := cmdutils.ParseFileFormat(importFlags.Format)

			if examplePath != "" {
				if cmd.Flags().Changed("file") || importFlags.DryRun {
					return fmt.Errorf(
						"--generate-example cannot be combined with --file or --dry-run",
					)
				}

				if format != cmdutils.FileFormatCsv {
					return fmt.Errorf(
						"example generation currently supports CSV only",
					)
				}

				if err := writeDeviceImportExample(examplePath); err != nil {
					return err
				}

				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Example CSV created: %s\n", examplePath)
				return nil
			}

			if importFlags.UpSert {
				return fmt.Errorf("--upsert is not implemented yet")
			}

			if cmd.Flags().Changed("batch") {
				return fmt.Errorf("--batch is not implemented yet")
			}

			switch format {
			case cmdutils.FileFormatCsv:
				return importDevices(cmd, a, importFlags)

			case cmdutils.FileFormatJSON, cmdutils.FileFormatYAML:
				return fmt.Errorf(
					"%s device import is not implemented yet",
					strings.ToLower(format.String()),
				)

			default:
				return fmt.Errorf(
					"unsupported import format %q (expected csv, json or yaml)",
					importFlags.Format,
				)
			}
		},
	}

	cmdutils.AttachImportFlags(cmd, &importFlags)
	cmd.Flags().StringVar(&examplePath, "generate-example", "", "Write an example CSV with three fictional devices")
	return cmd
}

func importDevices(cmd *cobra.Command, a *app.App, flags cmdutils.ImportOptions) error {
	if flags.SkipHeader {
		return fmt.Errorf("device import requires a CSV header")
	}

	if !strings.EqualFold(flags.Format, "csv") {
		return fmt.Errorf("unsupported import format %q", flags.Format)
	}

	delimiter := []rune(flags.CsvDelimiter)
	if len(delimiter) != 1 {
		return fmt.Errorf("--csv-delim must contain exactly one character")
	}

	r, closer, err := iox.OpenReader(flags.File)
	if err != nil {
		return err
	}
	defer closer()

	reader := csvx.NewReader(r, csvx.WithHasHeader(true), csvx.WithComma(delimiter[0]), csvx.WithTrimSpace(true))

	if _, err := reader.Header(); err != nil {
		return err
	}

	indexes, err := reader.BuildColumnIndex()
	if err != nil {
		return err
	}

	specs, err := device.CompileDeviceCsvSpecs(indexes)
	if err != nil {
		return err
	}

	var jobs []importx.Job[device.ImportParams]
	invalid := 0
	seen := make(map[string]int)

	for {
		row, line, err := reader.Next()

		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("read CSV: %w", err)
		}

		item, err := device.DecodeDeviceImportCsvSpecs(row, specs)
		if err != nil {
			invalid++
			fmt.Fprintf(cmd.ErrOrStderr(), "Line %d: %v\n", line, err)
			continue
		}

		devEUI := strings.ToLower(item.Device.DevEUI)

		if previous, exists := seen[devEUI]; exists {
			invalid++
			fmt.Fprintf(cmd.ErrOrStderr(), "Line %d: duplicate DevEUI (first seen at line %d)\n", line, previous)
			continue
		}

		seen[devEUI] = line

		jobs = append(jobs, importx.Job[device.ImportParams]{
			Line: line,
			Item: item,
		})
	}
	if len(jobs) == 0 {
		fmt.Fprintln(cmd.OutOrStdout(), "No devices to import")
		return nil
	}

	fmt.Fprintf(
		cmd.OutOrStdout(),
		"CSV validation: %d valid, %d invalid\n",
		len(jobs),
		invalid,
	)

	if flags.DryRun {
		fmt.Fprintln(
			cmd.OutOrStdout(),
			"DRY RUN: no devices have been created",
		)

		if invalid > 0 {
			return fmt.Errorf("dry-run failed: %d invalid rows", invalid)
		}
		return nil
	}

	if invalid > 0 {
		return fmt.Errorf(
			"import aborted: %d invalid rows; no devices created",
			invalid,
		)
	}

	return executeDeviceCSVImport(cmd, a, jobs)
}

func executeDeviceCSVImport(cmd *cobra.Command, a *app.App, jobs []importx.Job[device.ImportParams]) error {
	svc, err := a.Factory.Device()
	if err != nil {
		return err
	}
	runner := importx.NewRunner[device.ImportParams](
		cmd.Context(),
		func(ctx context.Context, j importx.Job[device.ImportParams]) error {
			_, err := svc.Create(ctx, device.CreateParams{
				Device: j.Item.Device,
			})
			if err != nil {
				return fmt.Errorf(
					"create device %s: %w",
					j.Item.Device.DevEUI,
					err,
				)
			}

			if j.Item.Keys != nil {
				if err := svc.CreateKeys(ctx, *j.Item.Keys); err != nil {
					return fmt.Errorf(
						"create keys for %s: %w",
						j.Item.Device.DevEUI,
						err,
					)
				}
			}

			return nil
		},
		importx.WithWorkers(8),
		importx.WithQueueSize(512),
		importx.WithTimeout(10*time.Second),
		importx.WithRatePerSecond(100),
		importx.WithBurst(32),
		importx.WithMinPeriod(time.Millisecond),
	)
	runner.Start()
	defer runner.Stop()

	type summary struct {
		success int
		failed  int
	}

	collected := make(chan summary, 1)

	// Collect results concurrently to avoid blocking workers.
	go func() {
		var result summary

		for res := range runner.Results() {
			if res.Err != nil {
				result.failed++
				fmt.Fprintf(
					cmd.ErrOrStderr(),
					"Line %d: %v\n",
					res.Line,
					res.Err,
				)
				continue
			}

			result.success++
		}

		collected <- result
	}()

	submitted := 0
	var submitErr error

	for _, job := range jobs {
		if err := runner.Submit(job); err != nil {
			submitErr = fmt.Errorf(
				"submit line %d: %w",
				job.Line,
				err,
			)
			break
		}

		submitted++
	}

	// Signal that no more jobs will arrive.
	runner.CloseJobs()

	// Wait for workers to finish and close the results channel.
	runner.Wait()

	// Wait for the result collector to finish.
	result := <-collected

	fmt.Fprintf(
		cmd.OutOrStdout(),
		"Import finished: %d submitted, %d successful, %d failed\n",
		submitted,
		result.success,
		result.failed,
	)

	if submitErr != nil {
		return submitErr
	}

	if err := cmd.Context().Err(); err != nil {
		return err
	}

	if result.success+result.failed != submitted {
		return fmt.Errorf(
			"import incomplete: %d submitted, %d results",
			submitted,
			result.success+result.failed,
		)
	}

	if result.failed > 0 {
		return fmt.Errorf(
			"device import completed with %d failures",
			result.failed,
		)
	}

	return nil
}
