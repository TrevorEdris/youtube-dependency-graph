package youtube

import (
	"context"
	"flag"
	"fmt"
	"html"
	"strings"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var (
	maxResults = flag.Int64("max-results", 25, "Max Youtube Results")
)

type Client interface {
	GetVideoByTitle(title string) (Video, error)
	GetVideoByURL(ytURL Url) (Video, error)
}

type ytClient struct {
	apiKey  string
	service *youtube.Service
}

func NewClient(apiKey string) (Client, error) {
	if apiKey == "" {
		return &ytClient{}, fmt.Errorf("provided api key is empty")
	}

	service, err := youtube.NewService(context.Background(), option.WithAPIKey(apiKey))
	if err != nil {
		return &ytClient{}, fmt.Errorf("unable to create youtube service: %s", err)
	}

	return &ytClient{
		apiKey:  apiKey,
		service: service,
	}, nil
}

func (c *ytClient) GetVideoByTitle(title string) (Video, error) {
	query := flag.String("query", title, "Search Query")

	// Perform a search, retrieving the 'id' and 'snippet' of the results,
	// limiting the number of results to maxResults
	searchListCall := c.service.Search.List([]string{"id", "snippet"}).Q(*query).MaxResults(*maxResults)
	response, err := searchListCall.Do()
	if err != nil {
		return &video{}, fmt.Errorf("unable to do call: %s", err)
	}

	// The Title's in the results are HTML-encoded, therefore to get an
	// accurate title match, we must also HTML-encode the given title
	titleToMatch := html.EscapeString(title)

	// Track the ID's of the video(s) that match, allowing us to query for
	// the full info of the video later on
	matchingVideoID := ""
	foundMatch := false
	for _, item := range response.Items {

		// We only care about Videos
		switch item.Id.Kind {
		case "youtube#video":

			// If the title of the item contains the title we're searching for,
			// then we treat that as a match.
			// TODO: Perform some input validation on the input value for title
			if strings.Contains(item.Snippet.Title, titleToMatch) {
				fmt.Printf("Found matching video! Title '%s'\n", item.Snippet.Title)
				matchingVideoID = item.Id.VideoId

				// Once we find a match, we don't really care about the rest
				// of the results
				foundMatch = true
			} else {
				fmt.Printf("Query found '%s' but '%s' was not a substring\n", item.Snippet.Title, titleToMatch)
			}
		}
		if foundMatch {
			break
		}
	}

	// TODO: Query for the full snippet of the video ID
	fmt.Printf("Video ID: %s\n", matchingVideoID)
	return c.getVideoByID(matchingVideoID)
}

func (c *ytClient) GetVideoByURL(url Url) (Video, error) {
	return c.getVideoByID(url.GetID())
}

func (c *ytClient) getVideoByID(id string) (Video, error) {
	// Query for all the relevant information
	videoListCall := c.service.Videos.List([]string{"id", "snippet", "contentDetails", "player"})

	// Get a video by the video ID
	videoListCall.Id(id)
	response, err := videoListCall.Do()
	if err != nil {
		return &video{}, fmt.Errorf("unable to perform video list by id: %s", err)
	}

	// We expect this ID to be unique, meaning only 0 or 1 result should be returned
	if len(response.Items) < 1 {
		return &video{}, fmt.Errorf("no videos found with id=%s", id)
	} else if len(response.Items) > 1 {
		return &video{}, fmt.Errorf("too many videos found (%d) with id=%s", len(response.Items), id)
	}

	fmt.Printf("Thumbnail: %s\n", response.Items[0].Snippet.Thumbnails.Maxres.Url)

	return newVideo(response.Items[0]), nil
}
