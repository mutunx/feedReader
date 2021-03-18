package search

import (
	"encoding/json"
	"log"
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
		log.Fatalf("file open error %s", err.Error())
	}

	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	return feeds, err

}
