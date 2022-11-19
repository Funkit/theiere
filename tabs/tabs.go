package tabs

import (
	"github.com/Funkit/crispy-engine/fancytext"
	"github.com/Funkit/crispy-engine/subview"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

var (
	//when the tab is active, the bottom border is empty
	activeTabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      " ",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┘",
		BottomRight: "└",
	}

	//when the tab is active, the bottom border is set
	inactiveTabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┴",
		BottomRight: "┴",
	}
	helpStyle = list.DefaultStyles().HelpStyle
)

type Model struct {
	Tabs             []string
	TabContents      []subview.Model
	ActiveTab        int
	inactiveTabStyle lipgloss.Style
	activeTabStyle   lipgloss.Style
	KeyMap           KeyMap
	Help             help.Model
	maxWidth         int
}

// Tab each tab is defined by its title and its content
type Tab struct {
	Name    string
	Content subview.Model
}

func NewTab(name string, content ...subview.Model) (Tab, error) {
	if len(content) != 0 {
		return Tab{
			Name:    name,
			Content: content[0],
		}, nil
	}

	cnt, err := fancytext.New("Hello World")
	if err != nil {
		return Tab{}, err
	}

	return Tab{
		Name:    name,
		Content: &cnt,
	}, nil
}

type options struct {
	width     *int
	color     *lipgloss.AdaptiveColor
	fixedSize bool
}

type Option func(options *options) error

func WithWidth(width int) Option {
	return func(options *options) error {
		options.width = &width

		return nil
	}
}

func WithColor(color lipgloss.AdaptiveColor) Option {
	return func(options *options) error {
		options.color = &color

		return nil
	}
}

// New builds a new tab model. Will panic if tab.content is not set to an actual common.Model.
func New(availableTabs []Tab, opts ...Option) (*Model, error) {
	var options options
	for _, opt := range opts {
		err := opt(&options)
		if err != nil {
			return nil, err
		}
	}

	width := 30
	if options.width != nil {
		width = *options.width
	}

	color := lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	if options.color != nil {
		color = *options.color
	}

	var tabNames []string
	var tabElements []subview.Model
	for _, val := range availableTabs {
		tabNames = append(tabNames, val.Name)
		tabElements = append(tabElements, val.Content)
	}

	inactiveTabStyle := lipgloss.NewStyle().Border(inactiveTabBorder, true).BorderForeground(color)

	return &Model{
		Tabs:             tabNames,
		TabContents:      tabElements,
		ActiveTab:        0,
		inactiveTabStyle: inactiveTabStyle,
		activeTabStyle:   inactiveTabStyle.Copy().Border(activeTabBorder, true),
		Help:             help.New(),
		KeyMap:           DefaultKeyMap(),
		maxWidth:         width,
	}, nil
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (subview.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.KeyMap.PrevPage):
			m.ActiveTab = max(m.ActiveTab-1, 0)
			return m, nil
		case key.Matches(msg, m.KeyMap.NextPage):
			m.ActiveTab = min(m.ActiveTab+1, len(m.Tabs)-1)
			return m, nil
		case key.Matches(msg, m.KeyMap.Quit):
			return m, subview.GoUp
		}
		m.TabContents[m.ActiveTab], _ = m.TabContents[m.ActiveTab].Update(msg)
	}

	_, cmd := m.TabContents[m.ActiveTab].Update(msg)

	return m, cmd
}

func (m *Model) View() string {
	var renderedTabs []string

	for i, t := range m.Tabs {
		var style lipgloss.Style
		isActive := i == m.ActiveTab
		if isActive {
			style = m.activeTabStyle.Copy()
		} else {
			style = m.inactiveTabStyle.Copy()
		}
		border, _, _, _, _ := style.GetBorder()
		style = style.Border(border)
		renderedTabs = append(renderedTabs, style.Render(t))
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)

	tabGap := m.inactiveTabStyle.Copy().
		BorderTop(false).
		BorderLeft(false).
		BorderRight(false)

	tabGapStr := strings.Repeat(" ", max(0, m.maxWidth-lipgloss.Width(row)))

	row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, tabGap.Render(tabGapStr))

	row = lipgloss.JoinVertical(lipgloss.Center, row, m.TabContents[m.ActiveTab].View())
	return lipgloss.JoinVertical(lipgloss.Left, row, helpStyle.Render(m.Help.View(m.KeyMap)))
}

func (m *Model) SetWidth(width int) {
	m.maxWidth = width
	for i := range m.TabContents {
		m.TabContents[i].SetWidth(width - 2)
	}
}

func (m *Model) SetHeight(height int) {
	for i := range m.TabContents {
		m.TabContents[i].SetHeight(height - 2)
	}
}

func (m *Model) Reset() {
	m.ActiveTab = 0
	for i := 0; i < len(m.TabContents); i++ {
		m.TabContents[i].Reset()
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
