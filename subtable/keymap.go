package subtable

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
)

type KeyMap interface {
	ShortHelp() []key.Binding
	FullHelp() [][]key.Binding
	asInternalTableMap() table.KeyMap
}

type KM struct {
	LineUp       key.Binding
	LineDown     key.Binding
	PageUp       key.Binding
	PageDown     key.Binding
	HalfPageUp   key.Binding
	HalfPageDown key.Binding
	GotoTop      key.Binding
	GotoBottom   key.Binding
}

// DefaultKeyMap returns a default set of keybindings.
func DefaultKeyMap() KM {
	return KM{
		LineUp: key.NewBinding(
			key.WithKeys("up"),
			key.WithHelp("↑", "up"),
		),
		LineDown: key.NewBinding(
			key.WithKeys("down"),
			key.WithHelp("↓", "down"),
		),
		PageUp: key.NewBinding(
			key.WithKeys("pgup"),
			key.WithHelp("pgup", "page up"),
		),
		PageDown: key.NewBinding(
			key.WithKeys("pgdown"),
			key.WithHelp("pgdn", "page down"),
		),
		HalfPageUp: key.NewBinding(
			key.WithKeys("u"),
			key.WithHelp("u", "½ page up"),
		),
		HalfPageDown: key.NewBinding(
			key.WithKeys("d"),
			key.WithHelp("d", "½ page down"),
		),
		GotoTop: key.NewBinding(
			key.WithKeys("t"),
			key.WithHelp("t", "go to top"),
		),
		GotoBottom: key.NewBinding(
			key.WithKeys("b"),
			key.WithHelp("b", "go to bottom"),
		),
	}
}

func (k KM) ShortHelp() []key.Binding {
	return []key.Binding{
		k.LineUp,
		k.LineDown,
		k.PageUp,
		k.PageDown,
		k.GotoTop,
		k.GotoBottom,
	}
}

func (k KM) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			k.LineUp,
			k.LineDown,
		},
		{
			k.PageUp,
			k.PageDown,
		},
		{
			k.HalfPageUp,
			k.HalfPageDown,
		},
		{
			k.GotoTop,
			k.GotoBottom,
		},
	}
}

func (k KM) asInternalTableMap() table.KeyMap {
	return table.KeyMap{
		LineUp:       k.LineUp,
		LineDown:     k.LineDown,
		PageUp:       k.PageUp,
		PageDown:     k.PageDown,
		HalfPageUp:   k.HalfPageUp,
		HalfPageDown: k.HalfPageDown,
		GotoTop:      k.GotoTop,
		GotoBottom:   k.GotoBottom,
	}
}
