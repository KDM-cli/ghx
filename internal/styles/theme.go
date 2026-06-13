package styles

import "github.com/charmbracelet/lipgloss"

var (
	App = lipgloss.NewStyle().
		Padding(1, 2)

	Panel = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("63")).
		Padding(1, 2)

	Header = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("213")).
		MarginBottom(1)

	Selected = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("86"))

	Cursor = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("86"))

	Muted = lipgloss.NewStyle().
		Foreground(lipgloss.Color("246"))

	Subtle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("244"))

	Info = lipgloss.NewStyle().
		Foreground(lipgloss.Color("81"))

	Success = lipgloss.NewStyle().
		Foreground(lipgloss.Color("82"))

	Error = lipgloss.NewStyle().
		Foreground(lipgloss.Color("196"))

	Badge = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("230")).
		Background(lipgloss.Color("63")).
		Padding(0, 1)

	Bar = lipgloss.NewStyle().
		MarginTop(1).
		PaddingTop(1).
		BorderTop(true).
		BorderForeground(lipgloss.Color("238"))
)
