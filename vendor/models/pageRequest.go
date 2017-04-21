package models

type PageRequest struct {
	ID    string `json:"id"`
	Depth int    `json:"depth"`
}
