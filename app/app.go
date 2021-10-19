package app

import (
	"fmt"

	"github.com/inconshreveable/log15"

	"github.com/TrevorEdris/youtube-dependency-graph/pkg/graph"
	"github.com/TrevorEdris/youtube-dependency-graph/pkg/youtube"
)

type App interface {
	GraphFromURL(url string) error
	GraphFromTitle(title string) error
	GraphFromID(id string) error
}

type app struct {
	cfg    Config
	client youtube.Client
	log    log15.Logger
}

func New(cfg Config, log log15.Logger) (App, error) {
	client, err := youtube.NewClient(cfg.Youtube.APIKey, log)
	if err != nil {
		return &app{}, err
	}

	return &app{
		cfg:    cfg,
		client: client,
		log:    log,
	}, nil
}

func (a *app) GraphFromURL(url string) error {
	video, err := a.client.GetVideoByURL(url)
	if err != nil {
		return err
	}

	return a.graphFromVideo(video)
}

func (a *app) GraphFromTitle(title string) error {
	video, err := a.client.GetVideoByTitle(title)
	if err != nil {
		return err
	}

	return a.graphFromVideo(video)
}

func (a *app) GraphFromID(id string) error {
	video, err := a.client.GetVideoByID(id)
	if err != nil {
		return err
	}

	return a.graphFromVideo(video)
}

func (a *app) graphFromVideo(video youtube.Video) error {
	a.log.Info("Generating graph for Video", "title", video.GetTitle(), "channel", video.GetChannelTitle())

	// Create a new graph, letting it create a unique ID for the graph
	g := graph.NewGraph("", "Youtube Video Dependencies", "ydg")

	refs, err := a.getReferences(g, video, 0)
	if err != nil {
		return err
	}
	a.log.Debug("Recursion completed")

	for _, vid := range refs {
		if vid == nil {
			continue
		}
		a.log.Debug("Video Reference", "title", vid.GetTitle(), "channel", vid.GetChannelTitle())
	}

	fmt.Println(g.ToCustomJSON())

	return nil
}

func (a *app) getReferences(g graph.Graph, video youtube.Video, currentDepth int) ([]youtube.Video, error) {
	a.log.Debug(fmt.Sprintf("======================[ %s, %s, %d ]======================", video.GetID(), video.GetChannelID(), currentDepth))
	a.log.Debug("Base video", "title", video.GetTitle(), "channel", video.GetChannelTitle())

	parentNode, err := graph.NewNode(video.GetID(), video.GetTitle())
	if err != nil {
		return []youtube.Video{}, nil
	}
	g.AddNode(parentNode)

	// TODO: Does this need to be the very top of the function?
	// Stop recursion once we reach the maximum recursion depth
	if currentDepth > a.cfg.Graph.MaxDepth {
		return []youtube.Video{}, nil
	}

	// Get all the URLs from the description of video
	urls := video.GetUrlsFromDescription()
	allRefs := make([]youtube.Video, len(urls)+1)
	allRefs = append(allRefs, video)

	// For each of the URLs in the description
	for _, url := range urls {

		// Get the Video that the URL is referencing
		referencedVideo, err := a.client.GetVideoByURL(url)
		if err != nil {
			a.log.Warn("Unable to get video", "input", url, "error", err)
			continue
		}

		childNode, err := graph.NewNode(referencedVideo.GetID(), referencedVideo.GetTitle())
		if err != nil {
			a.log.Warn("Unable to create new node from referenced video", "input", referencedVideo.GetChannelID(), "error", err)
			continue
		}

		g.AddEdge(parentNode, childNode, "references_via_description")
		a.log.Debug("Video reference", "title", referencedVideo.GetTitle(), "url", url, "channel", video.GetChannelTitle())

		// Get all the URLs in the referenced video's description
		childVids, err := a.getReferences(g, referencedVideo, currentDepth+1)
		if err != nil {
			a.log.Error("Unable to getReferences", "referencedVideoID", referencedVideo.GetID(), "currentDepth", currentDepth, "error", err)
			continue
		}

		// Append all the referenced video's URLs to the overall list of URLs
		allRefs = append(allRefs, childVids...)
	}

	return allRefs, nil
}
