package printing

import (
	"fmt"
	"slices"
	"strings"
)

type Format string

const (
	FormatTableStyled Format = "TABLE_STYLED"
	FormatTable       Format = "TABLE"
	FormatJSON        Format = "JSON"
	FormatYAML        Format = "YAML"
	FormatCSV         Format = "CSV"
	FormatUnknown     Format = "UNKNOWN"
)

func (enum Format) String() string {
	if len(strings.TrimSpace(string(enum))) == 0 {
		return string(FormatUnknown)
	}
	return string(enum)
}

func FormatValues() []Format {
	return []Format{
		FormatTableStyled,
		FormatTable,
		FormatJSON,
		FormatYAML,
		FormatCSV,
	}
}

func (enum Format) IsValid() bool {
	if slices.Contains(FormatValues(), enum) {
		return true
	}
	return false
}

func ParseFormat(value string) (Format, error) {
	normalized := strings.ToUpper(
		strings.ReplaceAll(strings.TrimSpace(value), "-", "_"),
	)
	switch normalized {
	case string(FormatTable):
		return FormatTable, nil
	case string(FormatJSON):
		return FormatJSON, nil
	case string(FormatYAML):
		return FormatYAML, nil
	case string(FormatCSV):
		return FormatCSV, nil
	case string(FormatTableStyled):
		return FormatTableStyled, nil
	default:
		return FormatUnknown, fmt.Errorf("unsupported output format %q", value)
	}

}
