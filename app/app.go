package app

import (
	"fmt"

	"github.com/TrevorEdris/youtube-dependency-graph/pkg/youtube"
)

type App interface {
	Run() error
}

type app struct {
	cfg Config
}

func New(cfg Config) (App, error) {
	return &app{
		cfg: cfg,
	}, nil
}

func (a *app) Run() error {
	// TODO: Parse CLI args here using urfave/cli/v2

	client, err := youtube.NewClient(a.cfg.Youtube.APIKey)
	if err != nil {
		return err
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
		return err
	}

	video, err := client.GetVideoByURL(ytUrl)
	if err != nil {
		return err
	}

	fmt.Printf("Video: %s\n", video.GetTitle())
	video.GetUrlsFromDescription()

	return nil
}
