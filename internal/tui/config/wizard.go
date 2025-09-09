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
	profile textinput.Model
	addr    textinput.Model
	token   textinput.Model
	timeout textinput.Model

	insecure bool

	// où on va écrire
	configPath string

	done bool
}

type Result struct {
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
	m.profile = textinput.New()
	m.profile.Placeholder = "Profile name (ex: prod)"
	m.profile.CharLimit = 64
	m.profile.Focus()

	m.addr = textinput.New()
	m.addr.Placeholder = "gRPC address (host:port), ex: chipstark.example.com:443"

	m.token = textinput.New()
	m.token.Placeholder = "Token API (bearer)"
	m.token.EchoMode = textinput.EchoPassword

	m.timeout = textinput.New()
	m.timeout.Placeholder = "Timeout (ex: 10s)"
	m.timeout.SetValue("10s")

	return m, nil
}

// ---- Update / View
func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		// Quit partout
		case "ctrl+c", "esc":
			return m, tea.Quit

		// Basculer secure/insecure à l’étape 2
		case "s":
			if m.step == 2 {
				m.insecure = false
			}
		case "i":
			if m.step == 2 {
				m.insecure = true
			}

		case "enter":
			// validation & passage à l’étape suivante
			if err := m.validateCurrent(); err != nil {
				m.err = err
				return m, nil
			}
			if m.step < 5 {
				m.step++
				// focus champ suivant
				switch m.step {
				case 1:
					m.addr.Focus()
					m.profile.Blur()
				case 3:
					m.token.Focus()
				case 4:
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
		m.profile, cmd = m.profile.Update(msg)
		return m, cmd
	case 1:
		var cmd tea.Cmd
		m.addr, cmd = m.addr.Update(msg)
		return m, cmd
	case 3:
		var cmd tea.Cmd
		m.token, cmd = m.token.Update(msg)
		return m, cmd
	case 4:
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
		return confirmStyle.Render("✔ Config enregistrée !") + "\n"
	}
	s := ""

	switch m.step {
	case 0:
		s += titleStyle.Render("Step 1/6 — Profile name") + "\n\n"
		s += m.profile.View() + "\n"
		s += helpStyle.Render("\nEnter for validation • ESC for quit") + "\n"

	case 1:
		s += titleStyle.Render("Step 2/6 — gRPC address") + "\n\n"
		s += m.addr.View() + "\n"
		s += helpStyle.Render("\nEx: chipstark.example.com:443 • enter for validation") + "\n"

	case 2:
		s += titleStyle.Render("Step 3/6 — TLS") + "\n\n"
		if m.insecure {
			s += "Mode : [i] insecure (TLS disabled)\n"
		} else {
			s += "Mode : [s] secure (TLS enabled)\n"
		}
		s += helpStyle.Render("\nAppuie sur 's' pour secure ou 'i' pour insecure, puis Entrée") + "\n"

	case 3:
		s += titleStyle.Render("Step 4/6 — Token API (Bearer)") + "\n\n"
		s += m.token.View() + "\n"
		s += helpStyle.Render("\nThe token can be left empty for test.") + "\n"

	case 4:
		s += titleStyle.Render("Step 5/6 — Timeout") + "\n\n"
		s += m.timeout.View() + "\n"
		s += helpStyle.Render("\nFormat: 5s, 10s, 1m… • Enter for validation") + "\n"

	case 5:
		s += titleStyle.Render("Step 6/6 — Confirmation") + "\n\n"
		s += fmt.Sprintf("Path: %s\n", m.configPath)
		s += fmt.Sprintf("Profil: %s\n", m.profile.Value())
		s += fmt.Sprintf("Addr  : %s\n", m.addr.Value())
		s += fmt.Sprintf("TLS   : %v (insecure=%v)\n", !m.insecure, m.insecure)
		if m.token.Value() != "" {
			s += "Token : ********\n"
		} else {
			s += "Token : (vide)\n"
		}
		s += fmt.Sprintf("Timeout: %s\n\n", m.timeout.Value())
		s += confirmStyle.Render("Enter for confirm - ESC for cancel") + "\n"
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
		if v := m.profile.Value(); v == "" {
			return errors.New("le nom de profil est requis")
		}
	case 1:
		if v := m.addr.Value(); v == "" {
			return errors.New("l'adresse gRPC est requise (host:port)")
		}
	case 4:
		if _, err := time.ParseDuration(m.timeout.Value()); err != nil {
			return fmt.Errorf("timeout invalide: %v", err)
		}
	}
	return nil
}

// ---- API externes
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
		Addr:       md.addr.Value(),
		Insecure:   md.insecure,
		CACert:     "",
		ServerName: md.addr.Value(),
		Token:      md.token.Value(),
		Timeout:    md.timeout.Value(),
		ApiVersion: "v4",
	}, nil
}
