package facebook

import (
	"log"
	"os"
)

// TODO(dtoledo23) : Implement Facebook crawler logic.
// Crawler implements crawler logic to fetch data from Facebook Graph API.
type Crawler struct {
	accessToken string
	sourceID    string
}

// NewCrawler initializes a crawler and set it up.
func NewCrawler(sourceID string) *Crawler {
	return &Crawler{
		accessToken: os.Getenv("FACEBOOK_PAGE_ACCESS_TOKEN"),
		sourceID:    sourceID,
	}
}

// Start crawling data wiht previous setup.
func (c *Crawler) Start() error {

	fbNode, err := c.fetchPage(c.sourceID)
	if err != nil {
		return err
	}

	log.Println("Fetched page: ", fbNode)
	fbNode.ToCrawlerNode().Save()

	// c.fetchEdge(c.sourceID)
	return nil
}
