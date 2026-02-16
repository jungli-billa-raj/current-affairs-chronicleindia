package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// https://www.chronicleindia.in/current-affairs/monthly/february-2025

func scrapeArticlePage(month string, year int) ([]article, error) {
	baseURL := "https://www.chronicleindia.in/current-affairs/monthly/"
	baseURL += fmt.Sprintf("%s-%d", month, year)

	// Request the HTML Page
	res, err := http.Get(baseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Status Code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	articles := []article{}

	doc.Find("h2.heading > a").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		href, _ := s.Attr("href")

		articles = append(articles, article{
			url:      "https://www.chronicleindia.in" + href,
			headline: text,
		})
	})
	// Rule to remember
	// Classes over structure. Always.

	return articles, nil
}

// func scrapeArticleDetails(url string)
