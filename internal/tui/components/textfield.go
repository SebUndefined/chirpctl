package components

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/sebundefined/chirpctl/internal/tui/styles"
)

type TextField struct {
	Title string
	Help  string
	input textinput.Model
	err   error
}

func NewTextField(title, placeholder, help string, masked bool, defaultValue string) TextField {
	ti := textinput.New()
	ti.Placeholder = placeholder
	if defaultValue != "" {
		ti.SetValue(defaultValue)
		ti.CursorEnd()
	}
	if masked {
		ti.EchoMode = textinput.EchoPassword
	}

	return TextField{
		Title: title,
		Help:  help,
		input: ti,
	}
}

func (t TextField) Init() tea.Cmd {
	return textinput.Blink
}

func (t TextField) Update(msg tea.Msg) (TextField, tea.Cmd) {
	var cmd tea.Cmd
	t.input, cmd = t.input.Update(msg)
	return t, cmd
}

func (t TextField) View() string {
	toRender := styles.Title.Render(t.Title) + "\n"
	toRender += t.input.View()
	if t.Help != "" {
		toRender += styles.Help.Render("\n"+t.Help) + "\n"
	}
	if t.err != nil {
		toRender += styles.Error.Render("Error:"+t.err.Error()) + "\n"
	}
	return toRender
}

func (t TextField) Value() string {
	return t.input.Value()
}

func (t *TextField) Focus() {
	t.input.Focus()
}

func (t *TextField) Blur() {
	t.input.Blur()
}
