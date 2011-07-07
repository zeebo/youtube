package youtube

import (
	"xml"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"http"
)

type entry struct {
	Title      string
	Rating     rating
	Statistics statistics
	Group      group
}

type rating struct {
	Average string "attr"
}

type statistics struct {
	ViewCount string "attr"
}

type group struct {
	Duration duration
}

type duration struct {
	Seconds string "attr"
}

type VideoInfo struct {
	Title    string
	Rating   float64
	Duration int
	Views    int
}

func Load(r io.Reader) (*VideoInfo, os.Error) {
	result := new(entry)
	if err := xml.Unmarshal(r, &result); err != nil {
		return nil, err
	}
	rating, err := strconv.Atof64(result.Rating.Average)
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

func LoadPath(url string) (*VideoInfo, os.Error) {
	path := fmt.Sprint("http://gdata.youtube.com/feeds/api/videos/", url)

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
	url, err := http.ParseURL(incoming)
	if err != nil {
		return false, ""
	}

	//check the host
	if url.Host != "youtube.com" && url.Host != "www.youtube.com" {
		return false, ""
	}
	
	//check the path to be a watch
	if url.Path != "/watch" {
		return false, ""
	}
	
	//Grab the v parameter from the query string
	v := url.Query().Get("v")
	
	//If we have a v paramater, return true
	return v != "", v
}
