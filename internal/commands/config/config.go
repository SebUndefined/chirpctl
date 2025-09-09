package config

import (
	"errors"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sebundefined/chirpctl/internal/app"
	"github.com/sebundefined/chirpctl/internal/config"
	configWizard "github.com/sebundefined/chirpctl/internal/tui/config"
	"github.com/spf13/cobra"
)

func newProfileCommand(a *app.App) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "profile",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	}
	cmd.AddCommand(
		newInitConfigCommand(a),
		newGetCurrentProfileCommand(a),
	)
	return cmd
}

func newGetCurrentProfileCommand(a *app.App) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "view",
		Short: "A brief description of your command",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			/*ctx := cmd.Context()
			svc := configctx.From(ctx)
			name, prof, err := svc.CurrentProfile(ctx)
			if err != nil {
				return err
			}

			p := tea.NewProgram(NewProfileView(name, prof))
			res, err := p.Run()
			return err
			a.CfgSvc.Init(cmd.Context())
			fmt.Println(a.Cfg)*/
			return nil
		},
	}

	return cmd
}

func newInitConfigCommand(a *app.App) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "init",
		Short: "Init the profile",
		Long:  `Guide you into the steps for creating a new profile.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			get, err := a.CfgSvc.Get(cmd.Context())
			if err != nil && !errors.Is(err, os.ErrNotExist) {
				return err
			}
			if get != nil && len(get.Contexts) > 0 {
				return errors.New("config Already Exists")
			}
			res, err := configWizard.Run()
			if err != nil {
				return err
			}

			init, err := a.CfgSvc.Init(cmd.Context(), config.File{
				Version:        1,
				DefaultContext: res.ServerName,
				Contexts: map[string]config.Context{
					res.Name: {
						Addr:       res.Addr,
						Insecure:   res.Insecure,
						CACert:     res.CACert,
						ServerName: res.ServerName,
						Token:      res.Token,
						Timeout:    res.Timeout,
						MTls:       config.MTls{},
						ApiVersion: res.ApiVersion,
					},
				},
			})
			if err != nil {
				return err
			}
			fmt.Println(res)
			fmt.Println(get)
			fmt.Println(init)
			return nil
		},
	}

	return cmd
}

// modèle
type profileViewModel struct {
	name  string
	prof  config.Context
	ready bool
}

func NewProfileView(name string, prof config.Context) tea.Model {
	return &profileViewModel{name: name, prof: prof, ready: true}
}

func (m *profileViewModel) Init() tea.Cmd {
	return nil // pas d’action asynchrone
}

func (m *profileViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// fermer sur "q" ou "esc"
		if msg.String() == "q" || msg.String() == "esc" {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m *profileViewModel) View() string {
	if !m.ready {
		return "Chargement du profil..."
	}

	return fmt.Sprintf(
		`Profil courant: %s

Adresse   : %s
Insecure  : %v
Token     : %s
Timeout   : %s
`,
		m.name,
		m.prof.Addr,
		m.prof.Insecure,
		truncateToken(m.prof.Token),
		m.prof.Timeout,
	)
}

func truncateToken(t string) string {
	if len(t) > 10 {
		return t[:4] + "..." + t[len(t)-4:]
	}
	return t
}
