package commands

import (
	"fmt"
	"sort"
	"strings"

	"github.com/sebundefined/chirpctl/internal/commands/shared"
	"github.com/spf13/cobra"
)

func NewVerbFirstRouter(verb string, targets map[shared.Resource][]string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   fmt.Sprintf("%s [resource]", verb),
		Short: fmt.Sprintf("%s a resource", strings.Title(verb)),
		Args:  cobra.MinimumNArgs(1),
		RunE: func(c *cobra.Command, args []string) error {
			res := shared.Resource(args[0])
			path, ok := targets[res]
			if !ok {
				// message clair
				keys := make([]string, 0, len(targets))
				for k := range targets {
					keys = append(keys, string(k))
				}
				sort.Strings(keys)
				return fmt.Errorf("unknown resource %q for verb %q (expected one of %v)", res, verb, keys)
			}

			// newArgs = chemin resource-first + args restants (flags etc.)
			newArgs := append(append([]string{}, path...), args[1:]...)

			root := c.Root()
			root.SetArgs(newArgs)
			// Repartir depuis la racine pour bénéficier des PersistentFlags/PreRun
			return root.Execute()
		},
	}
	cmd.DisableFlagParsing = true
	// complétion stable
	keys := make([]string, 0, len(targets))
	for k := range targets {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)
	cmd.ValidArgs = keys
	return cmd
}
