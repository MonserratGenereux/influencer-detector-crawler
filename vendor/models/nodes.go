package models

// Node defines a model to represent an node in the social graph.
type Node struct {
	// Page info
	Name        string
	Description string
	ID          string

	// Relevance
	Category     string
	CategoryList []string
	FanCount     int64

	// Location
	City    string
	Country string
	ZIP     int64

	// Developers controll
	Fetched bool
}
