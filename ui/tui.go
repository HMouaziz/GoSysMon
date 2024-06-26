package ui

import (
	"github.com/charmbracelet/bubbletea"
)

func start() {
	p := tea.NewProgram(
		newSimplePage("GoSysMon"),
	)
	if err, _ := p.Run(); err != nil {

	}
}
