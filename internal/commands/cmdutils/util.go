package cmdutils

import (
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func FlagChanged(cmd *cobra.Command, name string) bool {
	f := cmd.Flags().Lookup(name)
	if f == nil {
		f = cmd.InheritedFlags().Lookup(name)
	}
	return f != nil && f.Changed
}

func SetIfString(cmd *cobra.Command, name string, apply func(string)) {
	if FlagChanged(cmd, name) {
		v, _ := cmd.Flags().GetString(name)
		apply(v)
	}
}

func SetIfBool(cmd *cobra.Command, name string, apply func(bool)) {
	if FlagChanged(cmd, name) {
		v, _ := cmd.Flags().GetBool(name)
		apply(v)
	}
}

func SetIfDuration(cmd *cobra.Command, name string, apply func(time.Duration)) {
	if FlagChanged(cmd, name) {
		v, _ := cmd.Flags().GetDuration(name)
		apply(v)
	}
}

func IsEmptyString(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
