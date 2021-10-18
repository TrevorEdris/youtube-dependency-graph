package scrape

import (
    "fmt"
)

type Scraper interface {
    Scrape(ytURL string)
}

func New() Scraper {
    return &scraper{}
}

type scraper struct {}

func (s *scraper) Scrape(ytURL string) {
    fmt.Printf("Scraping %s\n", ytURL)
}
