package main

import (
	// "fmt"
	"log"

	list "github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	list list.Model
}

type article struct {
	url      string
	headline string
}

// This is required because bubbles/list expects items implementing list.Item.
type item string

func (i item) Title() string       { return string(i) }
func (i item) Description() string { return "" }
func (i item) FilterValue() string { return string(i) }

//

func initialModel(month string, year int) model {
	articles, err := scrape(month, year)
	if err != nil {
		log.Fatal(err)
	}

	items := make([]list.Item, 0, len(articles))
	for _, a := range articles {
		items = append(items, item(a.headline))
	}
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Current Affairs"

	return model{list: l}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width, msg.Height)
	}

	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.list.View()
}
