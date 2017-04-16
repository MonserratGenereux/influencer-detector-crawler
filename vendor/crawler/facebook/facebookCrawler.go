package facebook

import (
	"models"
	"os"
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
func NewCrawler(sourceID string) *Crawler {
	return &Crawler{
		accessToken:  os.Getenv("FACEBOOK_PAGE_ACCESS_TOKEN"),
		sourceID:     sourceID,
		currentDepth: 0,
	}
}

// Start crawling data wiht previous setup.
// It fetches the page of the given pageID and crawls all connected nodes pages
// This is pages liked by the source page.
// TODO(dtoledo23): Implement recursive crawling with depth.
func (c *Crawler) Start() error {

	if c.currentDepth >= maxDepth {
		return nil
	}

	// Fetch starting point
	fbNode, err := c.fetchPage(c.sourceID)
	if err != nil {
		return err
	}

	fbNode.ToCrawlerNode().Save()

	neighborshannel := make(chan facebookNode)

	// Concurrent go routine to process fetched nodes
	go c.processNeighborsNodes(neighborshannel)
	err = c.fetchAdjacentNodes(c.sourceID, neighborshannel)

	if err != nil {
		return err
	}

	return nil
}

func (c *Crawler) processNeighborsNodes(neighborshannel <-chan facebookNode) {
	for neighbor := range neighborshannel {
		// Transform to our defined Node model and save.
		neighbor.ToCrawlerNode().Save()

		// Create edge and store.
		edge := models.Edge{
			Source:      c.sourceID,
			Destination: neighbor.ID,
		}
		edge.Save()
	}
}
