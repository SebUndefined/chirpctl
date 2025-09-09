package config

import (
	"errors"
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/sebundefined/chirpctl/internal/config"
)

// ---- Modèle Bubble Tea
type model struct {
	step       int
	focusIndex int
	err        error

	// inputs
	name    textinput.Model
	addr    textinput.Model
	token   textinput.Model
	timeout textinput.Model

	// selections
	insecure   bool
	apiChoices []string
	apiIndex   int

	// où on va écrire
	configPath string

	done bool
}

type Result struct {
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

	// inputs
	m.name = textinput.New()
	m.name.Placeholder = "Context name (ex: prod)"
	m.name.CharLimit = 64
	m.name.Focus()

	m.addr = textinput.New()
	m.addr.Placeholder = "gRPC address (host:port), ex: chipstark.example.com:443"

	m.token = textinput.New()
	m.token.Placeholder = "API token (Bearer)"
	m.token.EchoMode = textinput.EchoPassword

	m.timeout = textinput.New()
	m.timeout.Placeholder = "Timeout (ex: 10s)"
	m.timeout.SetValue("10s")

	// API choices
	m.apiChoices = []string{"v4", "v3", "auto"} // adapte l’ordre si tu veux
	m.apiIndex = 0                              // défaut: v4

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

		// Step 3 (TLS): bascule avec s/i
		case "s":
			if m.step == 3 {
				m.insecure = false
			}
		case "i":
			if m.step == 3 {
				m.insecure = true
			}

		// Step 2 (API): navigation dans les choix
		case "left", "right", "up", "down", "tab":
			if m.step == 2 {
				if msg.String() == "left" || msg.String() == "up" {
					m.apiIndex = (m.apiIndex - 1 + len(m.apiChoices)) % len(m.apiChoices)
				} else {
					m.apiIndex = (m.apiIndex + 1) % len(m.apiChoices)
				}
				return m, nil
			}

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

var (
	titleStyle   = lipgloss.NewStyle().Bold(true)
	errorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff5f5f"))
	helpStyle    = lipgloss.NewStyle().Faint(true)
	confirmStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#5fff87"))
)

func (m model) View() string {
	if m.done {
		return confirmStyle.Render(" ") + "\n"
	}
	s := ""

	switch m.step {
	case 0:
		s += titleStyle.Render("Step 1/7 — Context name") + "\n\n"
		s += m.name.View() + "\n"
		s += helpStyle.Render("\nEnter to validate • ESC to quit") + "\n"

	case 1:
		s += titleStyle.Render("Step 2/7 — gRPC address") + "\n\n"
		s += m.addr.View() + "\n"
		s += helpStyle.Render("\nEx: chipstark.example.com:443 • Enter to validate") + "\n"

	case 2:
		s += titleStyle.Render("Step 3/7 — API version") + "\n\n"
		// simple rendu “boutons” avec highlight
		for i, v := range m.apiChoices {
			if i == m.apiIndex {
				s += lipgloss.NewStyle().Bold(true).Underline(true).Render("[ " + v + " ]")
			} else {
				s += "[ " + v + " ]"
			}
			if i < len(m.apiChoices)-1 {
				s += "  "
			}
		}
		s += "\n"
		s += helpStyle.Render("\nUse ←/→ or ↑/↓ (or Tab) to select • Enter to validate") + "\n"

	case 3:
		s += titleStyle.Render("Step 4/7 — TLS") + "\n\n"
		if m.insecure {
			s += "Mode : [i] insecure (TLS disabled)\n"
		} else {
			s += "Mode : [s] secure (TLS enabled)\n"
		}
		s += helpStyle.Render("\nPress 's' for secure or 'i' for insecure, then Enter") + "\n"

	case 4:
		s += titleStyle.Render("Step 5/7 — API token (Bearer)") + "\n\n"
		s += m.token.View() + "\n"
		s += helpStyle.Render("\nToken can be left empty for tests • Enter to validate") + "\n"

	case 5:
		s += titleStyle.Render("Step 6/7 — Timeout") + "\n\n"
		s += m.timeout.View() + "\n"
		s += helpStyle.Render("\nFormat: 5s, 10s, 1m… • Enter to validate") + "\n"

	case 6:
		s += titleStyle.Render("Step 7/7 — Confirmation") + "\n\n"
		s += fmt.Sprintf("Path     : %s\n", m.configPath)
		s += fmt.Sprintf("Context  : %s\n", m.name.Value())
		s += fmt.Sprintf("Addr     : %s\n", m.addr.Value())
		s += fmt.Sprintf("API      : %s\n", m.apiChoices[m.apiIndex])
		s += fmt.Sprintf("TLS      : %v (insecure=%v)\n", !m.insecure, m.insecure)
		if m.token.Value() != "" {
			s += "Token    : ********\n"
		} else {
			s += "Token    : (empty)\n"
		}
		s += fmt.Sprintf("Timeout  : %s\n\n", m.timeout.Value())
		s += confirmStyle.Render("Enter to confirm • ESC to cancel") + "\n"
	}

	if m.err != nil {
		s += "\n" + errorStyle.Render("Error: "+m.err.Error()) + "\n"
	}
	return s
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
	case 2:
		// rien de spécial: la sélection est toujours valide
	case 5:
		if _, err := time.ParseDuration(m.timeout.Value()); err != nil {
			return fmt.Errorf("invalid timeout: %v", err)
		}
	}
	return nil
}

func Run() (*Result, error) {
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
		return nil, errors.New("annulé")
	}
	return &Result{
		Name:       md.name.Value(),
		Addr:       md.addr.Value(),
		Insecure:   md.insecure,
		CACert:     "",
		ServerName: md.addr.Value(),
		Token:      md.token.Value(),
		Timeout:    md.timeout.Value(),
		ApiVersion: md.apiChoices[md.apiIndex], // <-- pris du sélecteur
	}, nil
}
