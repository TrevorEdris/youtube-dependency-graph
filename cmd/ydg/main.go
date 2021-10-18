package main

import (
	"fmt"
	"os"

	"github.com/TrevorEdris/youtube-dependency-graph/pkg/scrape"
)

func main() {
	fmt.Println("Hello, world!")

	// TODO: Parse CLI args here using urfave/cli/v2
	s := scrape.New()
	err := s.Scrape("https://www.youtube.com/watch?v=iDIcydiQOhc")
	if err != nil {
		fmt.Printf("ERROR: Unable to scrape: %s\n", err)
		os.Exit(1)
	}
}
