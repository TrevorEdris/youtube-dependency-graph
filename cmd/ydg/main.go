package main

import (
	"fmt"
	"os"

	"github.com/TrevorEdris/youtube-dependency-graph/pkg/youtube"
)

func main() {
	fmt.Println("Hello, world!")

	cfg, err := parseConfig()
	if err != nil {
		fmt.Printf("ERROR: Unable to parse config: %s\n", err)
		os.Exit(1)
	}

	// TODO: Parse CLI args here using urfave/cli/v2
	//ytURL := "https://www.youtube.com/watch?v=iDIcydiQOhc"
	//ytURL := "https://www.youtube.com/watch?v=acnvRrpvwlk"
	ytURL := "https://www.youtube.com/watch?v=-IfmgyXs7z8asdfasdf"
	//title := "New Results in Quantum Tunneling vs. The Speed of Light"

	client, err := youtube.NewClient(cfg.APIKey)
	if err != nil {
		fmt.Printf("ERROR: Unable to create client: %s\n", err)
		os.Exit(1)
	}

	//video, err := client.GetVideoByTitle(title)
	video, err := client.GetVideoByURL(ytURL)
	if err != nil {
		fmt.Printf("ERROR: Unable to create client: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Video: %s\n", video.GetTitle())
	video.GetUrlsFromDescription()
}
