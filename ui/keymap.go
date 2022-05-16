package ui

import (
	"github.com/charmbracelet/bubbles/key"
)

type keymap struct {
	Up    key.Binding
	Down  key.Binding
	Left  key.Binding
	Right key.Binding
	Roll  key.Binding
	Clear key.Binding
	Help  key.Binding
	Quit  key.Binding
}

var DefaultKeyMap = keymap{
	Up: key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp("↑/k", "increase quantity"),
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("↓/j", "decrease quantity"),
	),
	Left: key.NewBinding(
		key.WithKeys("h", "left"),
		key.WithHelp("🡠/h", "prev die"),
	),
	Right: key.NewBinding(
		key.WithKeys("l", "right"),
		key.WithHelp("🡢/l", "next die"),
	),
	Roll: key.NewBinding(
		key.WithKeys("e", "enter"),
		key.WithHelp("e/enter", "roll dice"),
	),
	Clear: key.NewBinding(
		key.WithKeys("c", "ctrl+l"),
		key.WithHelp("c", "clear"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c", "esc"),
		key.WithHelp("q", "quit"),
	),
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the key.Map interface.
func (k keymap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Roll, k.Quit}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k keymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Help, k.Roll, k.Quit},
		{k.Up, k.Down, k.Left, k.Right},
	}
}
