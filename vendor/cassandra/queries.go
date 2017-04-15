package cassandra

// InsertEdgeQuery defines query structure for inserting a new edge.
var InsertEdgeQuery = `
	INSERT INTO influencer_detector.nodes
  	( id, platform, name, description, categories, fan_count, city, country, zip)
	VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ? );
	`
