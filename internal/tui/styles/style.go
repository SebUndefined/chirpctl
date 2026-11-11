package styles

import "github.com/charmbracelet/lipgloss"

var (
	Title    = lipgloss.NewStyle().Bold(true)
	Error    = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff5f5f"))
	Help     = lipgloss.NewStyle().Faint(true)
	Confirm  = lipgloss.NewStyle().Foreground(lipgloss.Color("#5fff87"))
	Selected = lipgloss.NewStyle().Bold(true).Underline(true)
)
