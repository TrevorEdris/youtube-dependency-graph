package main

import (
	"fmt"

	"github.com/TrevorEdris/youtube-dependency-graph/pkg/scrape"
)

func main() {
    fmt.Println("Hello, world!")

    // TODO: Parse CLI args here using urfave/cli/v2
    s := scrape.New()
    s.Scrape("https://www.youtube.com/watch?v=iDIcydiQOhc")
}
