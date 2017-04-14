package facebook

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// TODO(dtoledo23) : Implement Facebook crawler logic.
const (
	graphAPIHost    = "graph.facebook.com"
	graphAPIVersion = "v2.8"
)

var (
	extractedPageFields = []string{
		"name",
		"about",
		"fan_count",
		"category",
		"category_list",
		"location",
	}
)

// Crawler implements crawler logic to fetch data from Facebook Graph API.
type Crawler struct {
	accessToken string
}

// NewCrawler initializes a crawler and set it up.
func NewCrawler() *Crawler {
	return &Crawler{
		accessToken: os.Getenv("FACEBOOK_PAGE_ACCESS_TOKEN"),
	}
}

// Start crawling data wiht previous setup.
func (c *Crawler) Start() error {
	pageID := "683165801724841"

	// Setup request call.
	query := url.Values{}
	query.Set("fields", strings.Join(extractedPageFields, ","))
	query.Set("access_token", c.accessToken)

	requestPage := &http.Request{
		Method: "GET",
		URL: &url.URL{
			Scheme:   "https",
			Host:     graphAPIHost,
			Path:     graphAPIVersion + "/" + pageID,
			RawQuery: query.Encode(),
		},
	}

	// Make API call.
	response, err := http.Get(requestPage.URL.String())
	if err != nil {
		return fmt.Errorf("Facebook API call error: %s", err)
	}

	defer response.Body.Close()

	rawResponseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("Error reading content of Facebook response: %s", err)
	}

	var page facebookNode
	if err := json.Unmarshal(rawResponseBody, &page); err != nil {
		return fmt.Errorf("Error docoding page response: %s", err)
	}

	if page.isEmpty() {
		return fmt.Errorf("Facebook error response: %s", string(rawResponseBody[:]))
	}

	log.Println("A huevo!", page)
	// TODO(dtoled23): Do something with the page object.
	return nil
}
