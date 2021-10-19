package youtube

import (
	"fmt"
	"regexp"
)

const (
	// TODO: Help wanted: Improve this regex
	vidIDregex = `((watch\?v=\S{11}){1})|(youtu\.be\/-?\w{10}\w?)|(^\w{11}$)`
)

type Url interface {
	GetID() string
	GetOrigin() string
}

type url struct {
	origin string
	ytID   string
}

func NewURL(ytURL string) (Url, error) {
	// Given "https://www.youtube.com/watch?v=iDIcydiQOhc", extract iDIcydiQOhc
	// and query yt API for that ID
	re := regexp.MustCompile(vidIDregex)
	res := re.FindAllStringSubmatch(ytURL, -1)
	if len(res) != 1 {
		return &url{}, fmt.Errorf("ERROR: invalid input string; Unable to parse video ID out of URL; Expected format %s, got result %v", vidIDregex, res)
	}

	// Pull the last 11 chars out of the match
	id := res[0][0][len(res[0][0])-11:]
	//fmt.Printf("Parsed id %s out of url %s\n", id, ytURL)

	return &url{
		origin: ytURL,
		ytID:   id,
	}, nil
}

func (u *url) GetID() string {
	return u.ytID
}

func (u *url) GetOrigin() string {
	return u.origin
}
