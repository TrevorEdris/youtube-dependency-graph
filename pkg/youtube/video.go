package youtube

import (
	"fmt"
	"regexp"

	"google.golang.org/api/youtube/v3"
)

const (
	urlRegex = `(https?:\/\/(?:www\.|(^www))[a-zA-Z0-9][a-zA-Z0-9-]+[a-zA-Z0-9]\.[^\s]{2,}|www\.[a-zA-Z0-9][a-zA-Z0-9-]+[a-zA-Z0-9]\.[^\s]{2,}|https?:\/\/(?:www\.|(^www))[a-zA-Z0-9]+\.[^\s]{2,}|www\.[a-zA-Z0-9]+\.[^\s]{2,})`
)

type Video interface {
	GetTitle() string
	GetDescription() string
	GetThumbnailURL() string
	GetUrlsFromDescription() []string
	GetYoutubeUrlsFromDescription() []string
}

/*
{
  "kind": "youtube#video",
  "etag": etag,
  "id": string,
  "snippet": {
    "publishedAt": datetime,
    "channelId": string,
    "title": string,
    "description": string,
    "thumbnails": {
      (key): {
        "url": string,
        "width": unsigned integer,
        "height": unsigned integer
      }
    },
    "channelTitle": string,
    "tags": [
      string
    ],
  },
  "contentDetails": {
    "duration": string,
  },
  "statistics": {
    "viewCount": unsigned long,
    "likeCount": unsigned long,
    "dislikeCount": unsigned long,
    "favoriteCount": unsigned long,
    "commentCount": unsigned long
  }
}
*/
type video struct {
	// Embed the youtube.Video type in our custom video type
	youtube.Video
}

func newVideo(vid *youtube.Video) Video {
	v := &video{}

	// Support all fields, because we don't know what future extensions will be made to
	// the video query
	v.AgeGating = vid.AgeGating
	v.ContentDetails = vid.ContentDetails
	v.Etag = vid.Etag
	v.FileDetails = vid.FileDetails
	v.Id = vid.Id
	v.Kind = vid.Kind
	v.LiveStreamingDetails = vid.LiveStreamingDetails
	v.Localizations = vid.Localizations
	v.MonetizationDetails = vid.MonetizationDetails
	v.Player = vid.Player
	v.ProcessingDetails = vid.ProcessingDetails
	v.ProjectDetails = vid.ProjectDetails
	v.RecordingDetails = vid.RecordingDetails
	v.Snippet = vid.Snippet
	v.Statistics = vid.Statistics
	v.Status = vid.Status
	v.Suggestions = vid.Suggestions
	v.TopicDetails = vid.TopicDetails
	return v
}

func (v *video) GetTitle() string {
	return v.Snippet.Title
}

func (v *video) GetDescription() string {
	return v.Snippet.Description
}

func (v *video) GetThumbnailURL() string {
	return v.Snippet.Thumbnails.Maxres.Url
}

func (v *video) GetUrlsFromDescription() []string {
	re := regexp.MustCompile(urlRegex)
	res := re.FindAllStringSubmatch(v.Snippet.Description, -1)
	urls := []string{}
	for _, matchGroup := range res {
		for _, match := range matchGroup {

			// Skip empty matches
			if match == "" {
				continue
			}

			// Only track unique matches
			if !contains(urls, match) {
				fmt.Printf("Found url: %s\n", match)
				urls = append(urls, match)
			}
		}
	}
	return urls
}

func (v *video) GetYoutubeUrlsFromDescription() []string {
	urls := v.GetUrlsFromDescription()
	// TODO: Filter out non-youtube URLs
	return urls
}

func contains(someList []string, someElement string) bool {
	contains := false
	for _, element := range someList {
		if element == someElement {
			contains = true
			break
		}
	}
	return contains
}
