package config

import (
	"errors"
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sebundefined/chirpctl/internal/config"
	"github.com/sebundefined/chirpctl/internal/tui/components"
	"github.com/sebundefined/chirpctl/internal/tui/styles"
)

// ---- Modèle Bubble Tea
type model struct {
	step       int
	focusIndex int
	err        error

	// inputs
	name       components.TextField
	apiVersion components.SelectField
	addr       components.TextField
	insecure   components.SelectField
	token      components.TextField
	timeout    components.TextField

	// où on va écrire
	configPath string

	done bool
}

type ResultNewContext struct {
	Name       string
	Addr       string
	Insecure   bool
	CACert     string
	ServerName string
	Token      string
	Timeout    string
	//MTls       MTls   `yaml:"mTls"`
	ApiVersion string
}

func initialModel() (model, error) {
	m := model{step: 0}

	// chemin où écrire le YAML
	p, err := config.Path()
	if err != nil {
		return m, err
	}
	m.configPath = p

	m.name = components.NewTextField("Step 1/7 — Name",
		"",
		"A friendly name for the new context", false, "")
	m.name.Focus()

	m.addr = components.NewTextField("Step 2/7 — Server address", "",
		"Address (host:port), ex: chirpstark.example.com:443", false, "")

	m.apiVersion = components.NewSelectField(
		"Step 3/7 — API version",
		"Use ←/→ or ↑/↓ (or Tab) to select • Enter to validate",
		0,
		[]string{"v4", "v3", "auto"})

	m.insecure = components.NewSelectField(
		"Step 4/7 — Insecure",
		"Use ←/→ or ↑/↓ (or Tab) to select • Enter to validate",
		0,
		[]string{"false", "true"})

	m.token = components.NewTextField("Step 5/7 — Token", "",
		"The token you have generated or provided by your provider", true, "")

	m.timeout = components.NewTextField("Step 6/7 — Timeout", "",
		"Timeout (ex: 10s)", false, "10s")

	return m, nil
}

// ---- Update / View
func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			if err := m.validateCurrent(); err != nil {
				m.err = err
				return m, nil
			}
			// progression
			if m.step < 6 {
				m.step++
				switch m.step {
				case 1:
					m.addr.Focus()
					m.name.Blur()
				case 3:
					// TLS step, pas d’input
				case 4:
					m.token.Focus()
				case 5:
					m.timeout.Focus()
				}
				return m, nil
			}
			m.done = true
			return m, tea.Quit
		}
	}

	// déléguer aux inputs selon l’étape
	switch m.step {
	case 0:
		var cmd tea.Cmd
		m.name, cmd = m.name.Update(msg)
		return m, cmd
	case 1:
		var cmd tea.Cmd
		m.addr, cmd = m.addr.Update(msg)
		return m, cmd
	case 2:
		var cmd tea.Cmd
		m.apiVersion, cmd = m.apiVersion.Update(msg)
		return m, cmd
	case 3:
		var cmd tea.Cmd
		m.insecure, cmd = m.insecure.Update(msg)
		return m, cmd
	case 4:
		var cmd tea.Cmd
		m.token, cmd = m.token.Update(msg)
		return m, cmd
	case 5:
		var cmd tea.Cmd
		m.timeout, cmd = m.timeout.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m model) View() string {
	if m.done {
		return styles.Confirm.Render(" ") + "\n"
	}

	switch m.step {
	case 0:
		return m.name.View()

	case 1:
		return m.addr.View()

	case 2:
		return m.apiVersion.View()

	case 3:
		return m.insecure.View()

	case 4:
		return m.token.View()

	case 5:
		return m.timeout.View()

	case 6:
		s := styles.Title.Render("Step 7/7 — Confirmation") + "\n\n"
		s += fmt.Sprintf("Path     : %s\n", m.configPath)
		s += fmt.Sprintf("Context  : %s\n", m.name.Value())
		s += fmt.Sprintf("Addr     : %s\n", m.addr.Value())
		s += fmt.Sprintf("API      : %s\n", m.apiVersion.Value())
		s += fmt.Sprintf("TLS      : %v (insecure=%v)\n", m.insecure.Value(), m.insecure.Value())
		if m.token.Value() != "" {
			s += "Token    : ********\n"
		} else {
			s += "Token    : (empty)\n"
		}
		s += fmt.Sprintf("Timeout  : %s\n\n", m.timeout.Value())
		s += styles.Confirm.Render("Enter to confirm • ESC to cancel") + "\n"
		return s
	}

	if m.err != nil {
		return styles.Error.Render("\n" + m.err.Error() + "\n")
	}
	return ""
}

// ---- validation + sauvegarde
func (m *model) validateCurrent() error {
	switch m.step {
	case 0:
		if v := m.name.Value(); v == "" {
			return errors.New("context name is required")
		}
	case 1:
		if v := m.addr.Value(); v == "" {
			return errors.New("gRPC address is required (host:port)")
		}
	case 5:
		if _, err := time.ParseDuration(m.timeout.Value()); err != nil {
			return fmt.Errorf("invalid timeout: %v", err)
		}
	}
	return nil
}

func RunNewContextWizard() (*ResultNewContext, error) {
	m, err := initialModel()
	if err != nil {
		return nil, err
	}
	final, err := tea.NewProgram(m).Run()
	if err != nil {
		return nil, err
	}
	md := final.(model)
	if !md.done {
		return nil, errors.New("canceled")
	}
	return &ResultNewContext{
		Name:       md.name.Value(),
		Addr:       md.addr.Value(),
		Insecure:   m.insecure.Value() == "true",
		CACert:     "",
		ServerName: md.addr.Value(),
		Token:      md.token.Value(),
		Timeout:    md.timeout.Value(),
		ApiVersion: md.apiVersion.Value(),
	}, nil
}
