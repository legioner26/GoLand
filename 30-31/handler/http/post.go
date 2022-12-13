package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"skillbox/30-31/driver"
	"skillbox/30-31/models"
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

	patchdata := &models.RequestCreate{}

	err := json.NewDecoder(r.Body).Decode(patchdata)
	if err == nil {
		infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
		infoLog.Printf("Получаем ответ %s", patchdata.Name)
		payload, _ := p.repo.Fetch(r.Context(), string(patchdata.Name))
		respondwithJSON(w, http.StatusOK, payload)
	}
	//name, _ := strconv.Atoi(chi.URLParam(r, "name"))
	//name, _ := r.Context().Value("name").(string)

}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
