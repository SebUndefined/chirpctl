package commands

import (
	"fmt"
	"sort"
	"strings"

	"github.com/sebundefined/chirpctl/internal/commands/shared"
	"github.com/spf13/cobra"
)

func NewVerbFirstRouter(verb string, targets map[shared.Resource]*cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:   fmt.Sprintf("%s [resource]", verb),
		Short: fmt.Sprintf("%s a resource", strings.Title(verb)),
		Args:  cobra.MinimumNArgs(1),
		RunE: func(c *cobra.Command, args []string) error {
			res := shared.Resource(args[0])
			if t, ok := targets[res]; ok {
				t.SetArgs(args[1:])
				return t.Execute()
			}
			// clé inconnue → message clair
			keys := make([]string, 0, len(targets))
			for k := range targets {
				keys = append(keys, string(k))
			}
			sort.Strings(keys)
			return fmt.Errorf("unknown resource %q for verb %q (expected one of %v)", res, verb, keys)
		},
	}

	// auto-completion stable
	keys := make([]string, 0, len(targets))
	for k := range targets {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)
	cmd.ValidArgs = keys
	return cmd
}
