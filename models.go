package main

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
