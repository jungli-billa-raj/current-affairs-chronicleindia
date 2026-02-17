package main

import (
	// "fmt"
	"log"

	list "github.com/charmbracelet/bubbles/list"
	viewport "github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	list     list.Model
	viewport viewport.Model
	view     view
}

type article struct {
	url      string
	headline string
}

type view int

const (
	listView view = iota
	articleView
)

// This is required because bubbles/list expects items implementing list.Item.
type item string

func (i item) Title() string       { return string(i) }
func (i item) Description() string { return "" }
func (i item) FilterValue() string { return string(i) }

//

func initialModel(month string, year int) model {
	articles, err := scrapeArticlePage(month, year)
	if err != nil {
		log.Fatal(err)
	}

	items := make([]list.Item, 0, len(articles))
	for _, a := range articles {
		items = append(items, item(a.headline))
	}
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Current Affairs"

	vp := viewport.New(0, 0)
	vp.SetContent("Loading Article.....")

	return model{list: l, viewport: vp, view: listView}
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
		case "esc":
			if m.view == articleView {
				m.view = listView
			} // I had put return m, nil line iside the if block. This propagated the message to listView and closed the app. Lesson Learnt
			return m, nil
		case "enter":
			if m.view == listView {
				m.view = articleView
				m.viewport.SetContent("Article Text here:")
			}
			return m, nil
		}
	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width, msg.Height)
		m.viewport.Width = msg.Width // so I'm assuming that msg contains the current Width and Height of the terminal
		m.viewport.Height = msg.Height
	}

	switch m.view {
	case articleView:
		m.viewport, cmd = m.viewport.Update(msg)
	case listView:
		m.list, cmd = m.list.Update(msg)
	}

	return m, cmd
}

func (m model) View() string {
	if m.view == articleView {
		return m.viewport.View()
	} else {
		return m.list.View()
	}
}
