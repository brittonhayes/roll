package ui

import (
	"github.com/charmbracelet/lipgloss"
)

func StyleHeader(s string) string {
	return lipgloss.NewStyle().Faint(true).MarginBottom(2).Render(s)
}

func StyleResult(s string) string {
	return lipgloss.NewStyle().Bold(true).MarginBottom(2).Render(s)
}

func StyleStatus(s string) string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("#f15bb5")).Render(s)
}

func StyleErr(s string) string {
	return lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{
		Light: "#202020",
		Dark:  "#f72585",
	}).Render(s)
}

func StyleKey(s string) string {
	return lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{
		Light: "#202020",
		Dark:  "#4cc9f0",
	}).Italic(true).Render(s)
}
