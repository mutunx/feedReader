package matchers

import (
	"encoding/xml"
	"errors"
	"feedReader/search"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

// 定义rss结构体
type (
	// item defines the fields associated with the item tag
	// in the rss document.
	item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        string   `xml:"guid"`
		GeoRssPoint string   `xml:"georss:point"`
	}

	// image defines the fields associated with the image tag
	// in the rss document.
	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}

	// channel defines the fields associated with the channel tag
	// in the rss document.
	channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubDate        string   `xml:"pubDate"`
		LastBuildDate  string   `xml:"lastBuildDate"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          image    `xml:"image"`
		Item           []item   `xml:"item"`
	}

	// rssDocument defines the fields associated with the rss document.
	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
)

type rssMatcher struct{}

func (r rssMatcher) Search(feed *search.Feed, searchItem string) ([]*search.Result, error) {
	rss, err := retrieve(feed)
	if err != nil {
		return nil, err
	}

	var results []*search.Result

	for _, item := range rss.Channel.Item {
		// 搜索
		matched, err := regexp.MatchString(searchItem, item.Title)
		if err != nil {
			return nil, err
		}

		if matched {
			results = append(results, &search.Result{
				Title:       item.Title,
				Description: item.Description,
				Link:        item.Link,
			})
		}

	}

	return results, err

}

func init() {
	var r rssMatcher
	search.Register("rss", r)
}

// 请求链接过去xml文本
func retrieve(feed *search.Feed) (*rssDocument, error) {
	// 防止地址为空
	if feed.Link == "" {
		return nil, errors.New("link is empty")
	}

	resp, err := http.Get(feed.Link)
	if err != nil {
		log.Fatalf("http get error %s", err.Error())
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("link %s code %d", feed.Link, resp.StatusCode)
	}

	// 解析到结构体中
	var rss rssDocument
	xml.NewDecoder(resp.Body).Decode(&rss)

	return &rss, err
}
