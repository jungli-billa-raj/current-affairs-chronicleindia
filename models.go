package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type articles struct {
	article []article
	cursor  int
}

type article struct {
	url      string
	headline string
}

func initialModel() articles {
	return articles{
		article: []article{{
			url:      "www.google.com",
			headline: "I'm feeling lucky",
		}, {
			url:      "www.youtube.com",
			headline: "Let's watch some videos",
		}, {
			url:      "mail.google.com",
			headline: "Check your emails regularly!!",
		}, {
			url:      "www.wikipedia.com",
			headline: "Let's learn something",
		}, {
			url:      "www.learnspanish.com",
			headline: "learn some spanish.",
		}},
		cursor: 0,
	}
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
