package twitpic

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

const (
	// first matching group is img url
	imgURLPattern = `<meta name="twitter:image" value="(.+)"`
	// second matching group is file extension
	fileExtPattern = `\/\d+\.([A-z]+)\?`
)

// Info is a wrapper around a twitpic img url and its extension
type Info struct {
	URL string
	Ext string
}

// GetInfo exists because twitpic urls return a html page that references to the actual image on another server
// this function is just to abstract away the html parsing etc
func GetInfo(url string) (Info, error) {
	res, err := http.Get(url)
	if err != nil {
		return Info{}, fmt.Errorf("unable to http get url %s: %v", url, err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return Info{}, fmt.Errorf("twitpic responded with non 200 status: %v", res.Status)
	}
	html, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Info{}, fmt.Errorf("unable to read response body: %v", err)
	}
	img, err := Parse(string(html))
	if err != nil {
		return Info{}, fmt.Errorf("unable to parse twitpic html: %v", err)
	}
	return img, nil
}

// Parse will do some regex magic on twitpic html and return a useable Img wrapper
func Parse(html string) (Info, error) {
	r, err := regexp.Compile(imgURLPattern)
	if err != nil {
		return Info{}, fmt.Errorf("unable to compile regex pattern %s: %v", imgURLPattern, err)
	}
	m := r.FindStringSubmatch(html)
	url := m[1]

	r, err = regexp.Compile(fileExtPattern)
	if err != nil {
		return Info{}, fmt.Errorf("unable to compile regex pattern %s: %v", fileExtPattern, err)
	}
	m = r.FindStringSubmatch(url)
	ext := m[1]
	return Info{URL: url, Ext: ext}, nil
}
