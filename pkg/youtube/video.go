package youtube

import (
	"regexp"
	"strings"

	"google.golang.org/api/youtube/v3"
)

const (
	//urlRegex = `(https?:\/\/(?:www\.|(^www))[a-zA-Z0-9][a-zA-Z0-9-]+[a-zA-Z0-9]\.[^\s]{2,}|www\.[a-zA-Z0-9][a-zA-Z0-9-]+[a-zA-Z0-9]\.[^\s]{2,}|https?:\/\/(?:www\.|(^www))[a-zA-Z0-9]+\.[^\s]{2,}|www\.[a-zA-Z0-9]+\.[^\s]{2,})`
	//urlRegex = `(http(s)?:\/\/)?((w){3}.)?youtu(be|.be)?(\.com)?\/.+`
	urlRegex = `(?:https?:\/\/)?(?:www\.)?(?:youtu\.be\/|youtube\.com\/(?:embed\/|v\/|watch\?v=|watch\?.+&v=))((\w|-){11})?`
)

type Video interface {
	GetID() string
	GetTitle() string
	GetDescription() string
	GetThumbnailURL() string
	GetUrlsFromDescription() []string
	GetChannelID() string
	GetChannelTitle() string
}

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

func (v *video) GetID() string {
	return strings.TrimSpace(v.Id)
}

func (v *video) GetTitle() string {
	return strings.TrimSpace(v.Snippet.Title)
}

func (v *video) GetDescription() string {
	return strings.TrimSpace(v.Snippet.Description)
}

func (v *video) GetThumbnailURL() string {
	return strings.TrimSpace(v.Snippet.Thumbnails.Maxres.Url)
}

func (v *video) GetUrlsFromDescription() []string {
	re := regexp.MustCompile(urlRegex)
	res := re.FindAllStringSubmatch(v.Snippet.Description, -1)
	urls := []string{}
	for _, matchGroup := range res {
		if len(matchGroup) == 0 {
			continue
		}

		// The first match in the regex matches the most subgroups
		match := strings.TrimSpace(matchGroup[0])

		// Skip empty matches
		if match == "" {
			continue
		}

		// Only track unique matches
		if !contains(urls, match) {
			//fmt.Printf("Found url: %s\n", match)
			urls = append(urls, match)
		}
	}
	return urls
}

func (v *video) GetEmbedHTML() string {
	return v.Player.EmbedHtml
}

// https://en.wikipedia.org/wiki/ISO_8601#Durations
func (v *video) GetDuration() string {
	return v.ContentDetails.Duration
}

func (v *video) GetChannelTitle() string {
	return v.Snippet.ChannelTitle
}

func (v *video) GetChannelID() string {
	return v.Snippet.ChannelId
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
