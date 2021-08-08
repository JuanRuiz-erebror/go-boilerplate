package model

// Bottle example
type Bottle struct {
	ID      int     `json:"id" example:"1"`
	Name    string  `json:"name" example:"bottle_name"`
	Account Account `json:"account"`
}
