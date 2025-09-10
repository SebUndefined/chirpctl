package components

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sebundefined/chirpctl/internal/tui/styles"
)

type SelectField struct {
	Title   string
	Help    string
	Choices []string
	Index   int
}

func NewSelectField(title, help string, index int, choices []string) SelectField {
	return SelectField{
		Title:   title,
		Help:    help,
		Choices: choices,
		Index:   index,
	}
}

func (s SelectField) Init() tea.Cmd {
	return nil
}

func (s SelectField) Update(msg tea.Msg) (SelectField, tea.Cmd) {
	switch k := msg.(type) {
	case tea.KeyMsg:
		switch k.String() {
		case "left", "up":
			s.Index = (s.Index - 1 + len(s.Choices)) % len(s.Choices)
		case "right", "down", "tab":
			s.Index = (s.Index + 1) % len(s.Choices)
		}
	}
	return s, nil
}

func (s SelectField) View() string {
	toRender := styles.Title.Render(s.Title + "\n")
	if s.Help != "" {
		toRender += styles.Title.Render(s.Help) + "\n"
	}
	for index, choice := range s.Choices {
		if s.Index == index {
			toRender += fmt.Sprintf("[%s]", styles.Selected.Render(choice))
		} else {
			toRender += fmt.Sprintf("[%s]", choice)
		}
		if index < len(s.Choices)-1 {
			toRender += "  "
		}
	}
	return toRender
}

func (s SelectField) Value() string {
	if len(s.Choices) == 0 {
		return ""
	}
	return s.Choices[s.Index]
}
