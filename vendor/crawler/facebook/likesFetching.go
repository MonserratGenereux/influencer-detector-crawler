package facebook

import (
	"encoding/json"
	"fmt"
	"log"
	"models"
	"strconv"
)

// Fetches all edges from the source pageID and sends them through the provided
// channel. The caller is responsible to receive all objects channeled and processed them.
// It uses the pagination mechanism from Facebook Graph API to iterate over all chunks
// Read more: https://developers.facebook.com/docs/graph-api/using-graph-api#paging
func (c *Crawler) fetchAdjacentNodes(pageID string) error {

	// Initial pagination call url.
	url := defaultGraphAPIURL(pageID+"/likes", c.accessToken).String()
	var err error

	// Iterate until there is no next page returned by fetchEdgeChunk.
	for len(url) != 0 {
		url, err = c.fetchAdjacentNodesChunk(url)
		if err != nil {
			return err
		}
	}

	return nil
}

// Fetches data chunk using the provided url and channels the results through
// neighborshannel. It returns a pagination pointer reference to the next page call.
func (c *Crawler) fetchAdjacentNodesChunk(url string) (string, error) {
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

	// Process neighbors connected by edges.
	c.processNeighborsNodes(neighbors.Edges)

	// Return reference to next page.
	return neighbors.Paging.Next, nil
}

// processNeighborsNodes is in charge of creating the edge object in the db.
// It also spawns a new crawler to continue fetching recursively.
func (c *Crawler) processNeighborsNodes(neighbors []facebookNode) {
	for _, neighbor := range neighbors {
		// Transform to our defined Node model and save.

		sourceID, _ := strconv.ParseInt(c.sourceID, 10, 64)
		destID, _ := strconv.ParseInt(neighbor.ID, 10, 64)

		// Create edge and store.
		edge := models.Edge{
			Source:      sourceID,
			Destination: destID,
		}

		err := edge.Save()
		if err != nil {
			log.Println(err)
		}

		crawler := NewCrawler(neighbor.ID, c.currentDepth-1)
		go crawler.Start()
	}
}
