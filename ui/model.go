package ui

import (
	"fmt"

	"github.com/brittonhayes/roll/parse"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/go-playground/validator/v10"
	"github.com/knipferrc/teacup/statusbar"
)

var validate = validator.New()

type Model struct {
	// Icon of the application
	icon string

	// Title displayed in the UI
	title string

	// Quantity of dice of selected dice in hand
	quantity int `validate:"positive"`

	// Cursor is the currently selected dice index
	cursor int `validate:"gte=0"`

	// Dice is the list of available die to choose from
	dice []string

	// Hand of dice +/- modifers to be rolled
	hand string

	// Roll result
	history string
	roll    string

	// Keys for the terminal interface
	keys keymap

	// Valid is the state of validity of the hand
	// for rolling
	valid bool

	// Err message from validation
	err error

	// Status is the current application status
	status statusbar.Bubble
	height int

	// Help message content
	help help.Model
}

func New() Model {
	return Model{
		icon:     "🎲",
		title:    "🎲 Roll CLI",
		quantity: 1,
		hand:     "1d4",
		cursor:   0,
		history:  "",
		dice:     []string{"d4", "d6", "d8", "d10", "d12", "d20", "d100"},
		keys:     DefaultKeyMap,
		help:     help.New(),
		status: statusbar.New(
			statusbar.ColorConfig{
				Foreground: lipgloss.AdaptiveColor{Light: "#202020", Dark: "#4a4e69"},
				Background: lipgloss.AdaptiveColor{Light: "#F25D94", Dark: "#7b2cbf"},
			},
			statusbar.ColorConfig{
				Foreground: lipgloss.AdaptiveColor{Light: "#202020", Dark: "#4a4e69"},
				Background: lipgloss.AdaptiveColor{Light: "#3c3836", Dark: "#303030"},
			},
			statusbar.ColorConfig{
				Foreground: lipgloss.AdaptiveColor{Light: "#202020", Dark: "#ebebeb"},
				Background: lipgloss.AdaptiveColor{Light: "#A550DF", Dark: "#3c096c"},
			},
			statusbar.ColorConfig{
				Foreground: lipgloss.AdaptiveColor{Light: "#202020", Dark: "#ebebeb"},
				Background: lipgloss.AdaptiveColor{Light: "#6124DF", Dark: "#5a189a"},
			},
		),
	}
}

func (m Model) IsValid() (bool, error) {
	if err := validate.Struct(m); err != nil {
		return false, err
	}

	_, err := parse.Match(m.hand)
	if err != nil || (m.hand == "") {
		return false, err
	}

	return true, nil
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Selection() string {
	return fmt.Sprintf("%d%s", m.quantity, m.dice[m.cursor])
}

func (m Model) IncreaseQuantity() (tea.Model, tea.Cmd) {
	if m.quantity > 0 {
		m.quantity++
	}

	return m.SetStatus()
}

func (m Model) DecreaseQuantity() (tea.Model, tea.Cmd) {
	if m.quantity > 1 {
		m.quantity--
	}

	return m.SetStatus()
}

func (m Model) NextDie() (tea.Model, tea.Cmd) {
	if m.cursor < len(m.dice) {
		m.cursor++
	}

	if m.cursor >= len(m.dice) {
		m.cursor = 0
	}

	return m.SetStatus()
}

func (m Model) PrevDie() (tea.Model, tea.Cmd) {
	if m.cursor >= 0 {
		m.cursor--
	}

	if m.cursor < 0 {
		m.cursor = len(m.dice) - 1
	}

	return m.SetStatus()
}

func (m Model) SetStatus() (tea.Model, tea.Cmd) {
	m.hand = m.Selection()
	if m.help.ShowAll {
		m.status.SetContent(m.icon, fmt.Sprintf("History: %s", m.history), fmt.Sprintf("#: %dx", m.quantity), fmt.Sprintf("Selection: %s", m.dice[m.cursor]))
		return m, nil
	}

	m.status.SetContent(m.icon, fmt.Sprintf("History: %s", m.history), fmt.Sprintf("%dx", m.quantity), fmt.Sprintf("%s", m.dice[m.cursor]))
	return m, nil
}

func (m Model) Roll() (tea.Model, tea.Cmd) {
	if m.valid == false {
		return m, nil
	}

	p, err := parse.NewParser(m.hand)
	if err != nil {
		m.roll = err.Error()
		return m, nil
	}

	m.history = m.roll
	m.roll = fmt.Sprintf("Rolled '%d' with %s", p.Roll(), p.String())
	return m.SetStatus()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.help.Width = msg.Width
		m.status.SetSize(msg.Width)
		m.status.SetContent(m.icon, "", fmt.Sprint(m.quantity), m.dice[m.cursor])

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, DefaultKeyMap.Left):
			m.valid, m.err = m.IsValid()
			m, cmd := m.PrevDie()
			return m, cmd

		case key.Matches(msg, DefaultKeyMap.Right):
			m.valid, m.err = m.IsValid()
			m, cmd := m.NextDie()
			return m, cmd

		case key.Matches(msg, DefaultKeyMap.Up):
			m.valid, m.err = m.IsValid()
			m, cmd := m.IncreaseQuantity()
			return m, cmd

		case key.Matches(msg, DefaultKeyMap.Down):
			m.valid, m.err = m.IsValid()
			m, cmd := m.DecreaseQuantity()
			return m, cmd

		case key.Matches(msg, DefaultKeyMap.Clear):
			m.roll = ""
			return m, nil

		case key.Matches(msg, DefaultKeyMap.Roll):
			m.valid, m.err = m.IsValid()
			if !m.valid {
				m.roll = fmt.Sprintf("%s - %q must be in the format 1d6 or 1d6+/-2", StyleErr(m.err.Error()), m.hand)
				return m, nil
			}

			return m.Roll()

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
			return m.SetStatus()

		case key.Matches(msg, DefaultKeyMap.Quit):
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m Model) View() string {

	title := StyleHeader(m.title)
	result := StyleResult(m.roll)
	help := m.help.View(m.keys)

	content := lipgloss.JoinVertical(lipgloss.Left, title, result, help)

	return lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.NewStyle().Height(m.height-statusbar.Height).Render(content),
		m.status.View(),
	)
}
