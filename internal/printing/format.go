package printing

import (
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

func ParseFormat(s string) Format {
	s = strings.ToUpper(strings.TrimSpace(s))
	switch s {
	case string(FormatTable):
		return FormatTable
	case string(FormatJSON):
		return FormatJSON
	case string(FormatYAML):
		return FormatYAML
	case string(FormatCSV):
		return FormatCSV
	default:
		return FormatTableStyled
	}

}
