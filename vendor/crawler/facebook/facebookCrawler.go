package facebook

import (
	"fmt"
	"models"
	"os"
	"strconv"
)

const (
	maxDepth = 2
)

// Crawler implements crawler logic to fetch data from Facebook Graph API.
type Crawler struct {
	accessToken  string
	sourceID     string
	currentDepth int
}

// NewCrawler initializes a crawler and set it up.
func NewCrawler(sourceID string, depth int) *Crawler {
	return &Crawler{
		accessToken:  os.Getenv("FACEBOOK_PAGE_ACCESS_TOKEN"),
		sourceID:     sourceID,
		currentDepth: depth,
	}
}

// Start crawling data wiht previous setup.
// It fetches the page of the given pageID and crawls all connected nodes pages
// This is pages liked by the source page.
// TODO(dtoledo23): Implement recursive crawling with depth.
func (c *Crawler) Start() error {
	fmt.Printf("Initializing crawler for page '%s' in depth %d\n", c.sourceID, c.currentDepth)
	// Fetch starting point
	fbNode, err := c.fetchPage(c.sourceID)
	if err != nil {
		return err
	}

	fbNode.ToCrawlerNode(c.currentDepth).Save()

	neighborshannel := make(chan facebookNode)

	// Concurrent go routine to process fetched nodes
	go c.processNeighborsNodes(neighborshannel)
	err = c.fetchAdjacentNodes(c.sourceID, neighborshannel)

	if err != nil {
		return err
	}

	fmt.Printf("Finished crawler for page '%s' in depth %d\n", c.sourceID, c.currentDepth)
	return nil
}

func (c *Crawler) processNeighborsNodes(neighborshannel <-chan facebookNode) {
	for neighbor := range neighborshannel {
		// Transform to our defined Node model and save.

		sourceID, _ := strconv.ParseInt(c.sourceID, 10, 64)
		destID, _ := strconv.ParseInt(neighbor.ID, 10, 64)

		// Create edge and store.
		edge := models.Edge{
			Source:      sourceID,
			Destination: destID,
		}
		edge.Save()

		if c.currentDepth > 1 {
			err := NewCrawler(neighbor.ID, c.currentDepth-1).Start()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			neighbor.ToCrawlerNode(0).Save()
		}
	}
}
