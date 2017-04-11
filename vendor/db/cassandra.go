package db

var (
	db interface{}
)

// GetConnection returns the driver connected to Cassandra.
func GetConnection() interface{} {

	if db != nil {
		db = connect()
	}

	return db
}

// Connect forces and refreshes the driver connection.
func Connect() {
	db = connect()
}

func connect() interface{} {

	// TODO (dtoledo23): Implement Cassandra connection logic.
	return nil
}
