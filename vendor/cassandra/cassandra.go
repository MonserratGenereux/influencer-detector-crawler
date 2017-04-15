package cassandra

import (
	"strings"

	"os"

	"log"

	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

// GetSession returns the driver connected to Cassandra.
func GetSession() *gocql.Session {

	if session == nil {
		session = connect()
	}

	return session
}

// Connect forces and refreshes the driver connection.
func Connect() {
	session = connect()
}

func connect() *gocql.Session {
	rawIPs := os.Getenv("CASSANDRA_IP_ADDRESSES")
	clusterIPs := strings.Split(rawIPs, ",")
	cluster := gocql.NewCluster(clusterIPs...)

	keyspace := os.Getenv("CASSANDRA_KEYSPACE")
	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.Any
	session, err := cluster.CreateSession()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Succesfully connected to Keyspace '%s' on Cassandra cluster: %s",
		keyspace,
		rawIPs)

	return session
}
