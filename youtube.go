package youtube

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type entry struct {
	Title  string `xml:"title"`
	Rating struct {
		XMLName xml.Name `xml:"rating"`
		Average float64  `xml:"average,attr"`
	}
	Group struct {
		XMLName  xml.Name `xml:"group"`
		Duration struct {
			XMLName xml.Name `xml:"duration"`
			Seconds int      `xml:"seconds,attr"`
		}
	}
	Statisitcs struct {
		XMLName xml.Name `xml:"statistics"`
		Views   int      `xml:"viewCount,attr"`
	}
}

type VideoInfo struct {
	Title    string
	Rating   float64
	Duration int
	Views    int
}

func Load(r io.Reader) (*VideoInfo, error) {
	var result entry
	if err := xml.NewDecoder(r).Decode(&result); err != nil {
		return nil, err
	}
	return &VideoInfo{
		Title:    result.Title,
		Rating:   result.Rating.Average,
		Duration: result.Group.Duration.Seconds,
		Views:    result.Statisitcs.Views,
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

	//ignore case
	u.Host = strings.ToLower(u.Host)

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
	resp = (v != "")
	return resp, v
}
