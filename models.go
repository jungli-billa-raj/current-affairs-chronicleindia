package main

import (
	"fmt"
	"log"

	list "github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	list list.Model
}

type articles struct {
	article []article
	cursor  int
}

type article struct {
	url      string
	headline string
}

func initialModel(month string, year int) articles {
	articles, err := scrape(month, year)
	if err != nil {
		log.Fatal(err)
	}
	listModel := model{}
	returnModel := append(listModel.list.Items(), articles)
	return returnModel
}

func (m articles) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m articles) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// is it a key press???
	case tea.KeyMsg:
		// cool, what was the actual key pressed?
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.article)-1 {
				m.cursor++
			}
		}
	}
	return m, nil
}

func (m articles) View() string {
	s := "Today's current Affairs\n"
	footer := "\nj/k UP/DOWN  ctrl+c QUIT"
	for i, v := range m.article {
		cursor := "  "
		if m.cursor == i {
			cursor = "->"
		}
		index := i + 1
		headline := v.headline
		s += fmt.Sprintf("%s %d.%s\n", cursor, index, headline)
	}
	s += footer
	return s
}
