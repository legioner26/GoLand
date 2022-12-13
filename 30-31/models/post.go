package models

// Post type details
type RequestSelect struct {
	Name   string `json:"name"`
	Age    int64  `json:"age"`
	Friend string `json:"friend"`
}
type RequestMakeFriend struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Friend string `json:"friend"`
}

type RequestAge struct {
	NewAge string `json:"new_age"`
}
