package device

import (
	"encoding/csv"
	"fmt"

	"github.com/sebundefined/chirpctl/internal/iox"
)

// writeDeviceImportExample writes three fictional devices.
// Existing files must never be overwritten.
func writeDeviceImportExample(path string) error {
	writer, closer, err := iox.OpenWriter(
		path,
		iox.WithNoClobber(),
	)
	if err != nil {
		return err
	}

	csvWriter := csv.NewWriter(writer)
	csvWriter.Comma = ';'

	rows := [][]string{
		{
			"devEUI", "name", "description",
			"applicationID", "deviceProfileID",
			"skipFcntCheck", "isDisabled",
			"tags", "variables",
		},
		{
			"0000000000000001",
			"example-device-01",
			"Example water meter 1",
			"11111111-1111-4111-8111-111111111111",
			"22222222-2222-4222-8222-222222222222",
			"false", "false",
			`{"type":"water-meter","site":"demo"}`,
			`{"environment":"example","label":"device-01"}`,
		},
		{
			"0000000000000002",
			"example-device-02",
			"Example water meter 2",
			"11111111-1111-4111-8111-111111111111",
			"22222222-2222-4222-8222-222222222222",
			"false", "false",
			`{"type":"water-meter","site":"demo"}`,
			`{"environment":"example","label":"device-02"}`,
		},
		{
			"0000000000000003",
			"example-device-03",
			"Example water meter 3",
			"11111111-1111-4111-8111-111111111111",
			"22222222-2222-4222-8222-222222222222",
			"false", "true",
			`{"type":"water-meter","site":"demo"}`,
			`{"environment":"example","label":"device-03"}`,
		},
	}

	writeErr := csvWriter.WriteAll(rows)
	closeErr := closer()

	if writeErr != nil {
		return fmt.Errorf("write example CSV: %w", writeErr)
	}
	if closeErr != nil {
		return fmt.Errorf("close example CSV: %w", closeErr)
	}

	return nil
}
