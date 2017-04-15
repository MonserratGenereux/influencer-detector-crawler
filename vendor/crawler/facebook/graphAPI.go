package facebook

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

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

func (c *Crawler) fetchEdges(pageID string, nodes chan *facebookNode) error {
	url := defaultGraphAPIURL(pageID+"/likes", c.accessToken).String()
	var err error

	for len(url) != 0 {
		url, err = c.fetchEdgeChunk(url, nodes)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Crawler) fetchEdgeChunk(url string, nodes chan *facebookNode) (string, error) {
	// Make likes edges request.
	rawResponseBody, err := callGraphAPI(url)
	if err != nil {
		return "", fmt.Errorf("Error fetching Facebook likes edges: %s", err)
	}

	// Decode json
	var edges facebookEdges
	if err := json.Unmarshal(rawResponseBody, &edges); err != nil {
		return "", fmt.Errorf("Error docoding edges response: %s", err)
	}

	// Send results to channel receiver
	// Edge processing logic is going to happend elsewhere.
	for _, edge := range edges.Edges {
		log.Println("Sending node", edge)
		nodes <- &edge
	}

	// Finished pagination crawling.
	if len(edges.Paging.Next) == 0 {
		return "", nil
	}

	return edges.Paging.Next, nil
}

func defaultCallGraphAPI(path, accessToken string) ([]byte, error) {
	return callGraphAPI(defaultGraphAPIURL(path, accessToken).String())
}

func defaultGraphAPIURL(path, accessToken string) *url.URL {
	query := url.Values{}
	query.Set("fields", strings.Join(extractedPageFields, ","))
	query.Set("access_token", accessToken)

	return &url.URL{
		Scheme:   "https",
		Host:     graphAPIHost,
		Path:     graphAPIVersion + "/" + path,
		RawQuery: query.Encode(),
	}
}

func callGraphAPI(url string) ([]byte, error) {

	// Make API call.
	response, err := http.Get(url)
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
