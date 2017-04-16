package facebook

import (
	"encoding/json"
	"fmt"
)

const (
	graphAPIHost    = "graph.facebook.com"
	graphAPIVersion = "v2.8"
)

var (
	// The fields we are going to extract from each fetched page.
	extractedPageFields = []string{
		"name",
		"about",
		"fan_count",
		"category",
		"category_list",
		"location",
	}
)

// Fetches a single page given a pageID. It is used to fetch the source point
// in the crawling mechanism.
func (c *Crawler) fetchPage(pageID string) (*facebookNode, error) {
	rawResponseBody, err := defaultCallGraphAPI(pageID, c.accessToken)
	if err != nil {
		return nil, fmt.Errorf("Error fetching Facebook page: %s", err)
	}

	var page facebookNode
	if err := json.Unmarshal(rawResponseBody, &page); err != nil {
		return nil, fmt.Errorf("Error docoding page response: %s", err)
	}

	return &page, nil
}
