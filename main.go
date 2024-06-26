package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mbndr/figlet4go"
	"os"
	"strings"
)

var keyMap = struct {
	NextTab   key.Binding
	PrevTab   key.Binding
	MoveLeft  key.Binding
	MoveRight key.Binding
	MoveDown  key.Binding
	MoveUp    key.Binding
	Select    key.Binding
	Exit      key.Binding
}{
	NextTab: key.NewBinding(
		key.WithKeys("right"),
		key.WithHelp("right", "next tab"),
	),
	PrevTab: key.NewBinding(
		key.WithKeys("left"),
		key.WithHelp("left", "previous tab"),
	),
	MoveLeft: key.NewBinding(
		key.WithKeys("left"),
		key.WithHelp("left", "move left"),
	),
	MoveRight: key.NewBinding(
		key.WithKeys("right"),
		key.WithHelp("right", "move right"),
	),
	MoveDown: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("down", "move down"),
	),
	MoveUp: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("up", "move up"),
	),
	Select: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
	),
	Exit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q/esc/ctrl+c", "exit"),
	),
}

type Model struct {
	options        []string
	selectedIndex  int
	tabs           []string
	activeTabIndex int
	inTabs         bool
	Width          int
	Height         int
}

func initialModel() Model {
	return Model{
		options:        []string{"Advanced Mode", "Settings", "Exit"},
		selectedIndex:  0,
		tabs:           []string{"CPU", "Memory", "Disk", "Network", "Processes"},
		activeTabIndex: 0,
		inTabs:         true,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func renderTitle() string {
	ascii := figlet4go.NewAsciiRender()
	options := figlet4go.NewRenderOptions()
	options.FontColor = []figlet4go.Color{
		figlet4go.ColorCyan,
	}
	options.FontName = "slant"
	err := ascii.LoadFont("assets/fonts/slant.flf")
	if err != nil {
		return "Error loading font: " + err.Error()
	}
	renderStr, _ := ascii.RenderOpts("GoSysMon", options)
	return renderStr
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case m.inTabs:
			switch {
			case key.Matches(msg, keyMap.NextTab):
				m.activeTabIndex = min(m.activeTabIndex+1, len(m.tabs)-1)
			case key.Matches(msg, keyMap.PrevTab):
				m.activeTabIndex = max(m.activeTabIndex-1, 0)
			case key.Matches(msg, keyMap.MoveDown):
				m.inTabs = !m.inTabs
			}
		case !m.inTabs:
			switch {
			case key.Matches(msg, keyMap.MoveLeft):
				if m.selectedIndex > 0 {
					m.selectedIndex--
				}
			case key.Matches(msg, keyMap.MoveRight):
				if m.selectedIndex < len(m.options)-1 {
					m.selectedIndex++
				}
			case key.Matches(msg, keyMap.MoveUp):
				m.inTabs = true
			case key.Matches(msg, keyMap.Select):
				if m.options[m.selectedIndex] == "Exit" {
					return m, tea.Quit
				}
			}
		case key.Matches(msg, keyMap.Exit):
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
	var b strings.Builder

	b.WriteString(renderTitle())
	b.WriteString("\n\n")

	for i, tab := range m.tabs {
		if m.inTabs && i == m.activeTabIndex {
			b.WriteString(fmt.Sprintf("[ %s ]", tab))
		} else {
			b.WriteString(fmt.Sprintf(" %s ", tab))
		}
	}
	b.WriteString("\n\n")
	b.WriteString(fmt.Sprintf("Content for %s tab \n\n", m.tabs[m.activeTabIndex]))

	b.WriteString("Options: ")
	for i, option := range m.options {
		if !m.inTabs && i == m.selectedIndex {
			b.WriteString(fmt.Sprintf("[ %s ]", option))
		} else {
			b.WriteString(fmt.Sprintf(" %s ", option))
		}
	}
	return b.String()
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
