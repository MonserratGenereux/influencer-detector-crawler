package facebook

import (
	"log"
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
func (c *Crawler) Start() error {

	if c.currentDepth >= maxDepth {
		return nil
	}

	fbNode, err := c.fetchPage(c.sourceID)
	if err != nil {
		return err
	}

	fbNode.ToCrawlerNode().Save()

	nodes := make(chan *facebookNode, 100)
	go procesNodes(nodes)
	err = c.fetchEdges(c.sourceID, nodes)
	close(nodes)

	if err != nil {
		return err
	}

	return nil
}

func procesNodes(nodesChannel chan *facebookNode) {
	for node := range nodesChannel {
		log.Println("Going to proccess", node)
	}
}
