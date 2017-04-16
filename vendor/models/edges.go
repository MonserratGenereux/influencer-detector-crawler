package models

import "cassandra"

// Edge defines a model to represent an edge in the social graph.
type Edge struct {
	// Source Node id
	Source string
	// Destination Node id
	Destination string
}

// Save the Node to the given database.
func (edge *Edge) Save() error {
	session := cassandra.GetSession()

	// id, platform, name, description, categories, fan_count, city, country, zip
	q := session.Query(cassandra.InsertEdgeQuery,
		edge.Source,
		edge.Destination,
	)

	return q.Exec()
}
