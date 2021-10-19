package main

import (
	"os"

	"github.com/inconshreveable/log15"
	"github.com/urfave/cli/v2"

	"github.com/TrevorEdris/youtube-dependency-graph/app"
)

const (
	appName        = "ydg"
	defaultVersion = "v0.0.0"

	flagURL   = "url"
	flagTitle = "title"
	flagID    = "id"
)

var (
	cfg app.Config

	appVersion = getVersion()
	log        = log15.New("module", appName)
)

func getVersion() string {
	v := os.Getenv("VERSION")
	if v == "" {
		log.Warn("No value for VERSION found, using default", "defaultVersion", defaultVersion)
		return defaultVersion
	}
	return v
}

func initLogging(logLevel, logFmt string) {
	lvl, err := log15.LvlFromString(logLevel)
	if err != nil {
		log.Warn("Invalid LOG_LEVEL specified, defaulting to info", "error", err)
		lvl = log15.LvlInfo
	}

	var loggerFormat log15.Format
	switch logFmt {
	case "logfmt":
		loggerFormat = log15.LogfmtFormat()
	case "jsonfmt":
		loggerFormat = log15.JsonFormat()
	case "terminalfmt":
		loggerFormat = log15.TerminalFormat()
	default:
		log.Warn("Invalid LOG_FMT specified, defaulting to terminalfmt", "LOG_FMT", logFmt)
		loggerFormat = log15.TerminalFormat()
	}

	log.SetHandler(
		log15.LvlFilterHandler(
			lvl,
			log15.StreamHandler(os.Stdout, loggerFormat),
		),
	)
}

func initSettings(c *cli.Context) error {
	var err error
	cfg, err = app.ParseConfig()
	if err != nil {
		log.Error("Unable to parse config", "error", err)
		return err
	}

	initLogging(cfg.Log.LogLevel, cfg.Log.LogFmt)
	return nil
}

func cliCreateGraphFromURL(c *cli.Context) error {
	ydg, err := app.New(cfg, log)
	if err != nil {
		log.Error("Unable to create ydg app", "error", err)
		return err
	}

	err = ydg.GraphFromURL(c.String(flagURL))
	if err != nil {
		log.Error("ydg execution failed", "error", err)
		return err
	}
	return nil
}

func cliCreateGraphFromTitle(c *cli.Context) error {
	ydg, err := app.New(cfg, log)
	if err != nil {
		log.Error("Unable to create ydg app", "error", err)
		return err
	}

	err = ydg.GraphFromTitle(c.String(flagTitle))
	if err != nil {
		log.Error("ydg execution failed", "error", err)
		return err
	}
	return nil
}

func cliCreateGraphFromID(c *cli.Context) error {
	ydg, err := app.New(cfg, log)
	if err != nil {
		log.Error("Unable to create ydg app", "error", err)
		return err
	}

	err = ydg.GraphFromID(c.String(flagID))
	if err != nil {
		log.Error("ydg execution failed", "error", err)
		return err
	}
	return nil
}

func newApplication() *cli.App {
	application := &cli.App{}
	application.Name = appName
	application.Version = appVersion
	application.Before = initSettings
	application.Commands = []*cli.Command{
		{
			Name:   "from-url",
			Usage:  "Create a dependency graph from a URL",
			Action: cliCreateGraphFromURL,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     flagURL,
					Usage:    "The URL of the youtube video to begin the graph with",
					Value:    "",
					Required: true,
				},
			},
		},
		{
			Name:   "from-title",
			Usage:  "Create a dependency graph from a video title",
			Action: cliCreateGraphFromTitle,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     flagTitle,
					Usage:    "The title of the youtube video to begin the graph with",
					Value:    "",
					Required: true,
				},
			},
		},
		{
			Name:   "from-id",
			Usage:  "Create a dependency graph from a video title",
			Action: cliCreateGraphFromID,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     flagID,
					Usage:    "The id of the youtube video to begin the graph with",
					Value:    "",
					Required: true,
				},
			},
		},
	}
	return application
}

func main() {
	ydg := newApplication()
	ydg.Run(os.Args)
}
