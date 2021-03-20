package search

import (
	"encoding/json"
	"os"
)

type Feed struct {
	Site string `json:"site"`
	Link string `json:"link"`
	Type string `json:"type"`
}

const path = "data/source.json"

func GetFeeds() ([]*Feed, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	return feeds, err

}
