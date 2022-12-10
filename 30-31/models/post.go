package models

// Post type details
type RequestCreate struct {
	Name    string   `json:"name"`
	Age     int64    `json:"age"`
	Friends []string `json:"friends"`
}
