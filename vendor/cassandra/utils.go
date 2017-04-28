package cassandra

import "github.com/gocql/gocql"

// GetNodeDepth fetched the depth of an already writen node
func GetNodeDepth(id int64) int {
	depthQuery := "SELECT depth FROM nodes WHERE id = ? LIMIT 1"
	row := session.Query(depthQuery, id).Consistency(gocql.One)

	var depth int
	err := row.Scan(&depth)

	if err != nil {
		return 0
	}

	return depth
}
