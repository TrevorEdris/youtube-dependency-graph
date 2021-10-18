package main

import (
	"fmt"
	"os"

	"github.com/TrevorEdris/youtube-dependency-graph/app"
)

func main() {
	cfg, err := app.ParseConfig()
	if err != nil {
		fmt.Printf("ERROR: Unable to parse config: %s\n", err)
		os.Exit(1)
	}

	ydg, err := app.New(cfg)
	if err != nil {
		fmt.Printf("ERROR: Unable to create app: %s\n", err)
		os.Exit(1)
	}

	err = ydg.Run()
	if err != nil {
		fmt.Printf("ERROR: Application failed: %s\n", err)
	}
}
