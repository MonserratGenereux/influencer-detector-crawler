package facebook

import (
	"cassandra"
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
	// Fetch starting point
	fbNode, err := c.fetchPage(c.sourceID)
	if err != nil {
		return err
	}

	node := fbNode.ToCrawlerNode(c.currentDepth)
	cacheDepth := cassandra.GetNodeDepth(node.ID)

	if (cacheDepth >= c.currentDepth) && (c.currentDepth != 0) {
		return nil
	}

	node.Save()

	if c.currentDepth <= 0 {
		return nil
	}

	// Concurrent go routine to process fetched nodes
	err = c.fetchAdjacentNodes(c.sourceID)

	if err != nil {
		return err
	}

	return nil
}
