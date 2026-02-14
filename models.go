package main

import (
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
		}},
		cursor: 0,
	}
}

func (m articles) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}
