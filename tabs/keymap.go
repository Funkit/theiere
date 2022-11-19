package tabs

import "github.com/charmbracelet/bubbles/key"

// KeyMap defines keybindings. It satisfies to the help.KeyMap interface
type KeyMap struct {
	// Keybindings used when browsing tabs.
	NextPage key.Binding
	PrevPage key.Binding

	// The quit keybinding. This won't be caught when filtering.
	Quit key.Binding
}

// DefaultKeyMap returns a default set of keybindings.
func DefaultKeyMap() KeyMap {
	return KeyMap{
		PrevPage: key.NewBinding(
			key.WithKeys("shift+tab"),
			key.WithHelp("shift+tab", "prev tab"),
		),
		NextPage: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "next tab"),
		),
		// Quitting.
		Quit: key.NewBinding(
			key.WithKeys("q", "esc"),
			key.WithHelp("q", "quit"),
		),
	}
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.NextPage,
		k.PrevPage,
		k.Quit,
	}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{{
		k.NextPage,
		k.PrevPage,
		k.Quit,
	}}
}
