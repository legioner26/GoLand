package models

// Post type details
type RequestCreate struct {
	Name    string   `json:"name"`
	Age     int64    `json:"age"`
	Friends []string `json:"friends"`
}
type RequestMakeFriend struct {
	SourceID string `json:"source_id"`
	TargetID string `json:"target_id"`
}

type RequestDeleteUser struct {
	TargetID string `json:"target_id"`
}

type RequestAge struct {
	NewAge string `json:"new age"`
}
