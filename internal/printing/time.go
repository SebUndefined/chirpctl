package printing

import "time"

// FormatTime formats an optional timestamp using the given layout.
// Nil and zero timestamps are displayed as "-".
// An empty layout uses Go's standard date-time format.
func FormatTime(value *time.Time, layout string) string {
	if value == nil || value.IsZero() {
		return "-"
	}

	if layout == "" {
		layout = time.DateTime
	}

	return value.Format(layout)
}
