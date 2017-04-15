package facebook

import (
	"encoding/json"
	"fmt"
)

// Fetches all edges from the source pageID and sends them through the provided
// channel. The caller is responsible to receive all objects channeled and processed them.
// It uses the pagination mechanism from Facebook Graph API to iterate over all chunks
// Read more: https://developers.facebook.com/docs/graph-api/using-graph-api#paging
func (c *Crawler) fetchAdjacentNodes(pageID string, neighborshannel chan<- facebookNode) error {

	// Initial pagination call url.
	url := defaultGraphAPIURL(pageID+"/likes", c.accessToken).String()
	var err error

	// Iterate until there is no next page returned by fetchEdgeChunk.
	for len(url) != 0 {
		url, err = c.fetchAdjacentNodesChunk(url, neighborshannel)
		if err != nil {
			return err
		}
	}

	return nil
}

// Fetches data chunk using the provided url and channels the results through
// neighborshannel. It returns a pagination pointer reference to the next page call.
func (c *Crawler) fetchAdjacentNodesChunk(url string, neighborshannel chan<- facebookNode) (string, error) {
	// Make likes edges request.
	rawResponseBody, err := callGraphAPI(url)
	if err != nil {
		return "", fmt.Errorf("Error fetching Facebook likes edges: %s", err)
	}

	// Parse json
	var neighbors facebookEdges
	if err := json.Unmarshal(rawResponseBody, &neighbors); err != nil {
		return "", fmt.Errorf("Error decoding neighbors response: %s", err)
	}

	// Send results to channel receiver
	// Edge processing logic is going to happend elsewhere.
	for _, neighbor := range neighbors.Edges {
		neighborshannel <- neighbor
	}

	// Return reference to next page.
	return neighbors.Paging.Next, nil
}
