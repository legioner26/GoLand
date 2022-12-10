package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"skillbox/30-31/driver"
	repository "skillbox/30-31/repository"
	post "skillbox/30-31/repository/post"
)

// NewPostHandler ...
func NewPostHandler(db *driver.DB) *Post {
	return &Post{
		repo: post.NewSQLPostRepo(db.SQL),
	}
}

// Post ...
type Post struct {
	repo repository.PostRepo
}

// Fetch all post data
func (p *Post) Fetch(w http.ResponseWriter, r *http.Request) {

	payload, _ := p.repo.Fetch(r.Context(), string(r.URL.Path[1:]))
	fmt.Println("dfdsfdsf", payload)
	respondwithJSON(w, http.StatusOK, payload)
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
