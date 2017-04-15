package cassandra

// InsertNodeQuery defines query structure for inserting a new node.
const InsertNodeQuery = `
	INSERT INTO influencer_detector.nodes
  	( id, platform, name, description, categories, fan_count, city, country, zip)
	VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ? );
	`

// InsertEdgeQuery defines query structure for inserting a new edge.
const InsertEdgeQuery = `
INSERT INTO influencer_detector.edges ( source, destination )
VALUES ( ?, ?);
`
