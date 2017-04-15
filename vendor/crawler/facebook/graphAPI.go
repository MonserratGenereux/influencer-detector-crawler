package facebook

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	graphAPIHost    = "graph.facebook.com"
	graphAPIVersion = "v2.8"
	depth           = 2
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

func (c *Crawler) fetchPage(pageID string) (*facebookNode, error) {
	rawResponseBody, err := callGraphAPI(pageID, c.accessToken)
	if err != nil {
		return nil, fmt.Errorf("Error fetching Facebook page: %s", err)
	}

	var page facebookNode
	if err := json.Unmarshal(rawResponseBody, &page); err != nil {
		return nil, fmt.Errorf("Error docoding page response: %s", err)
	}

	return &page, nil
}

func (c *Crawler) fetchEdges(pageID string) {

}

func (c *Crawler) fetchEdge(pageID string) error {
	rawResponseBody, err := callGraphAPI(pageID+"/likes", c.accessToken)
	if err != nil {
		return fmt.Errorf("Error fetching Facebook likes edges: %s", err)
	}

	var edges facebookEdges
	if err := json.Unmarshal(rawResponseBody, &edges); err != nil {
		return fmt.Errorf("Error docoding edges response: %s", err)
	}

	for i, edge := range edges.Edges {
		fmt.Println(i, edge)
	}

	fmt.Println(edges.Paging)
	return nil
}

func callGraphAPI(path, accessToken string) ([]byte, error) {
	query := url.Values{}
	query.Set("fields", strings.Join(extractedPageFields, ","))
	query.Set("access_token", accessToken)

	// Setup request
	request := &http.Request{
		Method: "GET",
		URL: &url.URL{
			Scheme:   "https",
			Host:     graphAPIHost,
			Path:     graphAPIVersion + "/" + path,
			RawQuery: query.Encode(),
		},
	}

	// Make API call.
	response, err := http.Get(request.URL.String())
	if err != nil {
		return nil, fmt.Errorf("Facebook API call error: %s", err)
	}

	// Proccess response
	defer response.Body.Close()
	rawResponseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading content of Facebook response: %s", err)
	}

	// Check if API returned an internal error.
	var apiError graphAPIError
	if err := json.Unmarshal(rawResponseBody, &apiError); err != nil {
		return nil, fmt.Errorf("Error docoding page response as error: %s", err)
	}

	if !apiError.isEmpty() {
		return nil, apiError
	}

	// All good.
	return rawResponseBody, nil
}
