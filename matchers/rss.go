package matchers

import (
	"encoding/xml"
	"feedReader/search"
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

func (r rssMatcher) Search(link string, searchItem string) ([]*search.Result, error) {
	// 请求地址
	resp, err := http.Get(link)
	if err != nil {
		log.Fatalf("http get error %s", err.Error())
	}

	if resp.StatusCode != 200 {
		log.Printf("link %s code %d", link, resp.StatusCode)
		return nil, nil
	}

	// 解析到结构体中
	var rss rssDocument
	xml.NewDecoder(resp.Body).Decode(&rss)

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
