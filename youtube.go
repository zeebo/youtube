package youtube

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type entry struct {
	Title      string
	Rating     rating
	Statistics statistics
	Group      group
}

type rating struct {
	Average string `xml:"attr"`
}

type statistics struct {
	ViewCount string `xml:"attr"`
}

type group struct {
	Duration duration
}

type duration struct {
	Seconds string `xml:"attr"`
}

type VideoInfo struct {
	Title    string
	Rating   float64
	Duration int
	Views    int
}

func Load(r io.Reader) (*VideoInfo, error) {
	result := new(entry)
	if err := xml.NewDecoder(r).Decode(&result); err != nil {
		return nil, err
	}
	rating, err := strconv.ParseFloat(result.Rating.Average, 64)
	if err != nil {
		return nil, err
	}
	duration, err := strconv.Atoi(result.Group.Duration.Seconds)
	if err != nil {
		return nil, err
	}
	views, err := strconv.Atoi(result.Statistics.ViewCount)
	if err != nil {
		return nil, err
	}

	return &VideoInfo{
		Title:    result.Title,
		Rating:   rating,
		Duration: duration,
		Views:    views,
	}, nil
}

func LoadPath(u string) (*VideoInfo, error) {
	path := fmt.Sprint("http://gdata.youtube.com/feeds/api/videos/", u)

	response, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return Load(response.Body)
}

func ValidUrl(incoming string) (resp bool, video string) {
	//Benefit of the doubt. Prefix with http:// if it's missing
	if !strings.HasPrefix(incoming, "http://") {
		incoming = "http://" + incoming
	}

	//Attempt to parse the url
	u, err := url.Parse(incoming)
	if err != nil {
		return false, ""
	}

	//check the host
	if u.Host != "youtube.com" && u.Host != "www.youtube.com" {
		return false, ""
	}

	//check the path to be a watch
	if u.Path != "/watch" {
		return false, ""
	}

	//Grab the v parameter from the query string
	v := u.Query().Get("v")

	//If we have a v paramater, return true
	return v != "", v
}
