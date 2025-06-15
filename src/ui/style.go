package ui

import "github.com/charmbracelet/lipgloss"

// Neon Color Scheme (https://colorhunt.co/palette/793fdf7091f597fff4fffd8c & https://colorhunt.co/palette/0b1d51725cad8ccdebffe3a9)
var (
	StyleTitle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("#725CAD")).
			Padding(0, 2).
			MarginTop(1).
			MarginBottom(1)

	StyleSelected = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#7091F5")).
			Bold(true).
			PaddingLeft(2).
			Underline(true)

	StyleCursor = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#97FFF4")).
			Bold(true)

	StyleError = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF0000")).
			Bold(true).
			MarginTop(1)

	StyleInput = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#FFFD8C")).
			Padding(0, 1)

	StyleCommand = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			MarginBottom(1).
			Width(80).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#6272a4"))

	StyleBox = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#97FFF4")).
			Padding(1, 2).
			Margin(1, 0)
)
