package backup

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/GoosvandenBekerom/tweaner/twitter"
)

// Save creates a backup of the given tweets in the folder specified by rootPath.
// It organises them by saving them in a folder structure like <rootPath>/<year>/<month>/<tweet.timestamp>.json
func Save(rootPath string, tweet twitter.Tweet) error {
	// Normalize rootPath by removing trailing path separators
	if strings.HasSuffix(rootPath, string(os.PathSeparator)) {
		rootPath = strings.TrimRight(rootPath, string(os.PathSeparator))
	}
	timestamp, err := tweet.CreatedAtTime()
	if err != nil {
		return fmt.Errorf("unable to parse timestamp of tweet with id %d: %v", tweet.Id, err)
	}
	path, err := generateFilePath(rootPath, timestamp)
	if err != nil {
		return fmt.Errorf("unable to create folder: %v", err)
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

func generateFilePath(root string, t time.Time) (string, error) {
	path := fmt.Sprintf("%s%c%d%c%s%c",
		root,
		os.PathSeparator,
		t.Year(),
		os.PathSeparator,
		t.Month(),
		os.PathSeparator)

	err := os.MkdirAll(path, 0666)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%d.json", path, t.Unix()), nil
}
