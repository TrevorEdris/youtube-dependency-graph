package scrape

import (
	"fmt"
)

type Scraper interface {
	Scrape(ytURL string) error
}

func New() Scraper {
	return &scraper{}
}

type scraper struct{}

func (s *scraper) Scrape(ytURL string) error {
	fmt.Printf("Scraping %s\n", ytURL)
	return nil
}
