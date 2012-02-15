package youtube

import (
	"bytes"
	"testing"
)

func TestValidUrl(t *testing.T) {
	cases := []struct {
		in   string
		resp bool
		out  string
	}{
		{"http://youtube.com/watch?v=BgmuWFcbQUU", true, "BgmuWFcbQUU"},
		{"http://youtube.com/watch?v=BgmuWFcbQUz", true, "BgmuWFcbQUz"},
		{"http://youtube.com/watch?v=", false, ""},
		{"http://www.youtube.com/watch?v=BgmuWFcbQUz", true, "BgmuWFcbQUz"},
		{"youtube.com/watch?v=BgmuWFcbQUz", true, "BgmuWFcbQUz"},
		{"Youtube.com/watch?v=BgmuWFcbQUz", true, "BgmuWFcbQUz"},
		{"metube.com/watch?v=BgmuWFcbQUz", false, ""},
	}

	for _, ca := range cases {
		if r, out := ValidUrl(ca.in); r != ca.resp || out != ca.out {
			t.Errorf("%q: Expected (%v, %q). Got (%v, %q)", ca.in, ca.resp, ca.out, r, out)
		}
	}
}

func TestLoadPath(t *testing.T) {
	cases := []struct {
		path  string
		title string
	}{
		{"BgmuWFcbQUU", "Dog Obedience FAIL"},
	}

	for _, ca := range cases {
		info, err := LoadPath(ca.path)
		if err != nil {
			t.Errorf("%q: Expected %v. Got %v", ca.path, nil, err)
			continue
		}

		if info.Title != ca.title {
			t.Errorf("%q: Expected %q. Got %q", ca.path, ca.title, info.Title)
		}
	}
}

const typical_response = `<?xml version="1.0" encoding="UTF-8"?>
<entry xmlns="http://www.w3.org/2005/Atom" xmlns:media="http://search.yahoo.com/mrss/" xmlns:gd="http://schemas.google.com/g/2005" xmlns:yt="http://gdata.youtube.com/schemas/2007">
    <id>http://gdata.youtube.com/feeds/api/videos/BgmuWFcbQUU</id>
    <published>2010-10-11T22:47:21.000Z</published>
    <updated>2012-02-15T20:36:57.000Z</updated>
    <category scheme="http://schemas.google.com/g/2005#kind" term="http://gdata.youtube.com/schemas/2007#video"/>
    <category scheme="http://gdata.youtube.com/schemas/2007/categories.cat" term="Entertainment" label="Entertainment"/>
    <category scheme="http://gdata.youtube.com/schemas/2007/keywords.cat" term="dog"/>
    <category scheme="http://gdata.youtube.com/schemas/2007/keywords.cat" term="beach"/>
    <category scheme="http://gdata.youtube.com/schemas/2007/keywords.cat" term="ocean"/>
    <category scheme="http://gdata.youtube.com/schemas/2007/keywords.cat" term="black"/>
    <category scheme="http://gdata.youtube.com/schemas/2007/keywords.cat" term="lady"/>
    <category scheme="http://gdata.youtube.com/schemas/2007/keywords.cat" term="voice"/>
    <category scheme="http://gdata.youtube.com/schemas/2007/keywords.cat" term="crap"/>
    <category scheme="http://gdata.youtube.com/schemas/2007/keywords.cat" term="ball"/>
    <category scheme="http://gdata.youtube.com/schemas/2007/keywords.cat" term="eww"/>
    <category scheme="http://gdata.youtube.com/schemas/2007/keywords.cat" term="nasty"/>
    <category scheme="http://gdata.youtube.com/schemas/2007/keywords.cat" term="defiant"/>
    <category scheme="http://gdata.youtube.com/schemas/2007/keywords.cat" term="obedient"/>
    <category scheme="http://gdata.youtube.com/schemas/2007/keywords.cat" term="lol"/>
    <category scheme="http://gdata.youtube.com/schemas/2007/keywords.cat" term="funny"/>
    <category scheme="http://gdata.youtube.com/schemas/2007/keywords.cat" term="lmao"/>
    <category scheme="http://gdata.youtube.com/schemas/2007/keywords.cat" term="rwj"/>
    <category scheme="http://gdata.youtube.com/schemas/2007/keywords.cat" term="=3"/>
    <category scheme="http://gdata.youtube.com/schemas/2007/keywords.cat" term="rofl"/>
    <category scheme="http://gdata.youtube.com/schemas/2007/keywords.cat" term="epic"/>
    <category scheme="http://gdata.youtube.com/schemas/2007/keywords.cat" term="boxer"/>
    <category scheme="http://gdata.youtube.com/schemas/2007/keywords.cat" term="don't"/>
    <category scheme="http://gdata.youtube.com/schemas/2007/keywords.cat" term="go"/>
    <category scheme="http://gdata.youtube.com/schemas/2007/keywords.cat" term="in"/>
    <category scheme="http://gdata.youtube.com/schemas/2007/keywords.cat" term="water"/>
    <category scheme="http://gdata.youtube.com/schemas/2007/keywords.cat" term="so"/>
    <category scheme="http://gdata.youtube.com/schemas/2007/keywords.cat" term="good"/>
    <category scheme="http://gdata.youtube.com/schemas/2007/keywords.cat" term="joby"/>
    <category scheme="http://gdata.youtube.com/schemas/2007/keywords.cat" term="jobi"/>
    <category scheme="http://gdata.youtube.com/schemas/2007/keywords.cat" term="shit"/>
    <category scheme="http://gdata.youtube.com/schemas/2007/keywords.cat" term="poo"/>
    <category scheme="http://gdata.youtube.com/schemas/2007/keywords.cat" term="poop"/>
    <category scheme="http://gdata.youtube.com/schemas/2007/keywords.cat" term="old"/>
    <category scheme="http://gdata.youtube.com/schemas/2007/keywords.cat" term="accident"/>
    <title type="text">Dog Obedience FAIL</title>
    <content type="text">This dog is the SHIT lol</content>
    <link rel="alternate" type="text/html" href="http://www.youtube.com/watch?v=BgmuWFcbQUU&amp;feature=youtube_gdata"/>
    <link rel="http://gdata.youtube.com/schemas/2007#video.responses" type="application/atom+xml" href="http://gdata.youtube.com/feeds/api/videos/BgmuWFcbQUU/responses"/>
    <link rel="http://gdata.youtube.com/schemas/2007#video.related" type="application/atom+xml" href="http://gdata.youtube.com/feeds/api/videos/BgmuWFcbQUU/related"/>
    <link rel="http://gdata.youtube.com/schemas/2007#mobile" type="text/html" href="http://m.youtube.com/details?v=BgmuWFcbQUU"/>
    <link rel="self" type="application/atom+xml" href="http://gdata.youtube.com/feeds/api/videos/BgmuWFcbQUU"/>
    <author>
        <name>critikill</name>
        <uri>http://gdata.youtube.com/feeds/api/users/critikill</uri>
    </author>
    <gd:comments>
        <gd:feedLink rel="http://gdata.youtube.com/schemas/2007#comments" href="http://gdata.youtube.com/feeds/api/videos/BgmuWFcbQUU/comments" countHint="282"/>
    </gd:comments>
    <media:group>
        <media:category label="Entertainment" scheme="http://gdata.youtube.com/schemas/2007/categories.cat">Entertainment</media:category>
        <media:content url="http://www.youtube.com/v/BgmuWFcbQUU?version=3&amp;f=videos&amp;app=youtube_gdata" type="application/x-shockwave-flash" medium="video" isDefault="true" expression="full" duration="93" yt:format="5"/>
        <media:content url="rtsp://v1.cache2.c.youtube.com/CiILENy73wIaGQlFQRtXWK4JBhMYDSANFEgGUgZ2aWRlb3MM/0/0/0/video.3gp" type="video/3gpp" medium="video" expression="full" duration="93" yt:format="1"/>
        <media:content url="rtsp://v6.cache4.c.youtube.com/CiILENy73wIaGQlFQRtXWK4JBhMYESARFEgGUgZ2aWRlb3MM/0/0/0/video.3gp" type="video/3gpp" medium="video" expression="full" duration="93" yt:format="6"/>
        <media:description type="plain">This dog is the SHIT lol</media:description>
        <media:keywords>dog, beach, ocean, black, lady, voice, crap, ball, eww, nasty, defiant, obedient, lol, funny, lmao, rwj, =3, rofl, epic, boxer, don't, go, in, water, so, good, joby, jobi, shit, poo, poop, old, accident</media:keywords>
        <media:player url="http://www.youtube.com/watch?v=BgmuWFcbQUU&amp;feature=youtube_gdata_player"/>
        <media:thumbnail url="http://i.ytimg.com/vi/BgmuWFcbQUU/0.jpg" height="360" width="480" time="00:00:46.500"/>
        <media:thumbnail url="http://i.ytimg.com/vi/BgmuWFcbQUU/1.jpg" height="90" width="120" time="00:00:23.250"/>
        <media:thumbnail url="http://i.ytimg.com/vi/BgmuWFcbQUU/2.jpg" height="90" width="120" time="00:00:46.500"/>
        <media:thumbnail url="http://i.ytimg.com/vi/BgmuWFcbQUU/3.jpg" height="90" width="120" time="00:01:09.750"/>
        <media:title type="plain">Dog Obedience FAIL</media:title>
        <yt:duration seconds="93"/>
    </media:group>
    <gd:rating average="4.8702292" max="5" min="1" numRaters="2096" rel="http://schemas.google.com/g/2005#overall"/>
    <yt:statistics favoriteCount="1680" viewCount="239116"/>
</entry>`

func TestLoad(t *testing.T) {
	buf := bytes.NewBufferString(typical_response)
	vi, err := Load(buf)
	if err != nil {
		t.Fatal(err)
	}
	ex := &VideoInfo{
		Title:    "Dog Obedience FAIL",
		Rating:   4.8702292,
		Duration: 93,
		Views:    239116,
	}

	if *vi != *ex {
		t.Errorf("Expected %v. Got %v", ex, vi)
	}
}
