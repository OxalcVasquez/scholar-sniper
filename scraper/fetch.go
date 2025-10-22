package scraper

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type ScrapData struct {
	Title        string
	Deadline     string
	Requirements string
}

func ScrapeScholarships(url string) (*ScrapData, error) {
	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		return nil, err
	}

	title := strings.TrimSpace(doc.Find("h1").First().Text())
	deadline := strings.TrimSpace(doc.Find(".deadline").First().Text())
	requirements := strings.TrimSpace(doc.Find(".requirements").First().Text())

	return &ScrapData{
		Title:        title,
		Deadline:     deadline,
		Requirements: requirements,
	}, nil
}
