package printing

import (
	"fmt"
	"slices"
	"strings"

	"github.com/samber/lo"
)

// FormatStringMap returns sorted key-value pairs for terminal output.
func FormatStringMap(values map[string]string) string {
	if len(values) == 0 {
		return "-"
	}
	keys := lo.Keys(values)
	slices.Sort(keys)
	strValues := make([]string, len(keys))
	for i, key := range keys {
		strValues[i] = fmt.Sprintf("%s=%s", key, values[key])
	}
	return strings.Join(strValues, ", ")
}
