package backup

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/GoosvandenBekerom/tweaner/twitter"
	"github.com/GoosvandenBekerom/tweaner/util/twitpic"
)

// Save creates a backup of the given tweets in the folder specified by rootPath.
// It organises them by saving them in a folder structure like <rootPath>/<year>/<month>/<tweet.timestamp>.json
func Save(rootPath string, tweet twitter.Tweet) error {
	// Normalize rootPath by removing trailing path separators
	if strings.HasSuffix(rootPath, string(os.PathSeparator)) {
		rootPath = strings.TrimRight(rootPath, string(os.PathSeparator))
	}
	err := downloadMedia(rootPath, tweet)
	if err != nil {
		return fmt.Errorf("unable to download media of tweet with id %d: %v", tweet.Id, err)
	}
	path, err := generateFilePath(rootPath, "", "json", tweet)
	if err != nil {
		return fmt.Errorf("unable to generate file path: %v", err)
	}
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("unable to create file with path %s: %v", path, err)
	}
	defer f.Close()
	encoder := json.NewEncoder(f)
	if err := encoder.Encode(tweet); err != nil {
		return fmt.Errorf("unable to encode tweet with id %d into file %s: %v", tweet.Id, path, err)
	}
	return nil
}

func downloadMedia(rootPath string, tweet twitter.Tweet) error {
	for i, urls := range tweet.Entities.Urls {
		if !strings.HasPrefix(urls.Display_url, "twitpic.com") {
			// it used to be very common to tweet pictures using twitpic.com
			// (at least for me, might remove this or make it more generic later)
			continue
		}
		info, err := twitpic.GetInfo(urls.Expanded_url)
		if err != nil {
			return fmt.Errorf("unable to download twitpic image: %v", err)
		}
		path, err := generateFilePath(rootPath, fmt.Sprintf("_twitpic_%d", i), info.Ext, tweet)
		if err != nil {
			return fmt.Errorf("unable to generate file path: %v", err)
		}
		return downloadFile(info.URL, path)
	}
	for i, media := range tweet.Entities.Media {
		switch media.Type {
		case "photo":
			ext := filepath.Ext(media.Media_url)
			path, err := generateFilePath(rootPath, fmt.Sprintf("_%d", i), ext, tweet)
			if err != nil {
				return fmt.Errorf("unable to generate file path: %v", err)
			}
			if err = downloadFile(media.Media_url_https, path); err != nil {
				return fmt.Errorf("unable to download image: %v", err)
			}
		default:
			// ignore media type
			continue
		}
	}
	return nil
}

func downloadFile(url, path string) error {
	out, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("unable to create file: %v", err)
	}
	defer out.Close()

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("unable to http get url %s: %v", url, err)
	}
	defer res.Body.Close()

	_, err = io.Copy(out, res.Body)
	if err != nil {
		return fmt.Errorf("unable to write response body into file %s: %v", url, err)
	}
	return nil
}

func generateFilePath(root, suffix, ext string, tweet twitter.Tweet) (string, error) {
	t, err := tweet.CreatedAtTime()
	if err != nil {
		return "", fmt.Errorf("unable to parse timestamp of tweet with id %d: %v", tweet.Id, err)
	}
	path := fmt.Sprintf("%s%c%d%c%s%c",
		root,
		os.PathSeparator,
		t.Year(),
		os.PathSeparator,
		t.Month(),
		os.PathSeparator)
	if err = os.MkdirAll(path, 0666); err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%d%s.%s", path, tweet.Id, suffix, ext), nil
}
