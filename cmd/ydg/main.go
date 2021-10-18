package main

import (
	"fmt"
	"os"

	"github.com/TrevorEdris/youtube-dependency-graph/pkg/youtube"
)

func main() {
	cfg, err := parseConfig()
	if err != nil {
		fmt.Printf("ERROR: Unable to parse config: %s\n", err)
		os.Exit(1)
	}

	// TODO: Parse CLI args here using urfave/cli/v2

	client, err := youtube.NewClient(cfg.APIKey)
	if err != nil {
		fmt.Printf("ERROR: Unable to create client: %s\n", err)
		os.Exit(1)
	}

	/*
		title := "New Results in Quantum Tunneling vs. The Speed of Light"
		video, err := client.GetVideoByTitle(title)
		if err != nil {
			fmt.Printf("ERROR: Unable to get video by title: %s\n", err)
			os.Exit(1)
		}
	*/

	//url := "https://www.youtube.com/watch?v=iDIcydiQOhc"
	//url := "https://www.youtube.com/watch?v=acnvRrpvwlk"
	url := "https://www.youtube.com/watch?v=-IfmgyXs7z8asdfasdf"

	ytUrl, err := youtube.NewURL(url)
	if err != nil {
		fmt.Printf("ERROR: Invalid URL provided: %s\n", err)
	}

	video, err := client.GetVideoByURL(ytUrl)
	if err != nil {
		fmt.Printf("ERROR: Unable to get video by url: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Video: %s\n", video.GetTitle())
	video.GetUrlsFromDescription()
}
