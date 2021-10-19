package main

import (
	"fmt"
	"os"

	"github.com/inconshreveable/log15"

	"github.com/TrevorEdris/youtube-dependency-graph/app"
)

var (
	log = log15.New("module", "ydg")
)

func initLogging(logLevel string) {
	log.SetHandler(log15.MultiHandler(
		log15.StreamHandler(os.Stderr, log15.LogfmtFormat()),
		log15.LvlFilterHandler(
			log15.LvlInfo,
			log15.Must.FileHandler("errors.json", log15.JsonFormat()),
		),
	))
}

func main() {
	cfg, err := app.ParseConfig()
	if err != nil {
		fmt.Printf("ERROR: Unable to parse config: %s\n", err)
		os.Exit(1)
	}

	ydg, err := app.New(cfg, log)
	if err != nil {
		fmt.Printf("ERROR: Unable to create app: %s\n", err)
		os.Exit(1)
	}

	err = ydg.Run()
	if err != nil {
		fmt.Printf("ERROR: Application failed: %s\n", err)
	}
}
