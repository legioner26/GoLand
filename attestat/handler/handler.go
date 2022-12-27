package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"skillbox/attestat/models"
	"skillbox/attestat/repository"
	"strconv"
	"unicode"
)

// NewPostHandler ...
func NewPostHandler(db *repository.Service) *Post {
	return &Post{repo: db}
}

// Post ...
type Post struct {
	repo repository.PostRepo
}

func (h *Post) Update(w http.ResponseWriter, r *http.Request) {}

func (h *Post) GetFull(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {

		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	infoLog.Printf("GET: Full info %v", id)

	city, err := h.repo.GetFull(id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	message, err := json.Marshal(city)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	infoLog.Printf("Sending: %v", string(message))

	w.WriteHeader(http.StatusOK)
	w.Write(message)
}

func (h *Post) Create(w http.ResponseWriter, r *http.Request) {
	var city cities.CityRequest
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	infoLog.Printf("POST: New city %v", string(content))

	err = json.Unmarshal(content, &city)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	city.Name = toUpperFirst(city.Name)
	city.Region = toUpperFirst(city.Region)
	city.District = toUpperFirst(city.District)

	id, err := h.repo.Create(city)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	message := fmt.Sprintf("The city of %v was created with the id %v", city.Name, id)
	respondwithJSON(w, http.StatusCreated, message)
}

func (h *Post) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	infoLog.Printf("DELETE: Delete city %v", id)

	err = h.repo.Delete(id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	message := fmt.Sprintf("City with identifier %v deleted", id)
	respondwithJSON(w, http.StatusOK, message)
}

func (h *Post) SetPopulation(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	infoLog.Printf("PUT: Update population city %v", id)

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var population cities.SetPopulationRequest
	err = json.Unmarshal(content, &population)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.repo.SetPopulation(id, population.Population)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondwithJSON(w, http.StatusOK, "Change was successful")
}

func (h *Post) GetFromRegion(w http.ResponseWriter, r *http.Request) {
	region := chi.URLParam(r, "region")
	region = toUpperFirst(region)

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	infoLog.Printf("GET: Сities by region %v", region)

	cityNames, err := h.repo.GetFromRegion(region)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondwithJSON(w, http.StatusOK, cityNames)
}

func (h *Post) GetFromDistrict(w http.ResponseWriter, r *http.Request) {
	district := chi.URLParam(r, "district")
	district = toUpperFirst(district)

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	infoLog.Printf("GET: Сities by district %v", district)

	cityNames, err := h.repo.GetFromDistrict(district)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondwithJSON(w, http.StatusOK, cityNames)
}

func (h *Post) GetFromPopulation(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	infoLog.Printf("GET: Сities by population range %v", string(content))

	var populationRange cities.RangeRequest
	err = json.Unmarshal(content, &populationRange)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	cityNames, err := h.repo.GetFromPopulation(populationRange.Start, populationRange.End)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondwithJSON(w, http.StatusOK, cityNames)
}

func (h *Post) GetFromFoundation(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	infoLog.Printf("GET: Сities by foundation range %v", string(content))

	var foundationRange cities.RangeRequest
	err = json.Unmarshal(content, &foundationRange)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	cityNames, err := h.repo.GetFromFoundation(foundationRange.Start, foundationRange.End)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}
	respondwithJSON(w, http.StatusOK, cityNames)
}

func toUpperFirst(text string) string {
	textRune := []rune(text)
	for i := range textRune {
		textRune[i] = unicode.ToLower(textRune[i])
	}
	return string(unicode.ToUpper(textRune[0])) + string(textRune[1:])
}

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

// Fetch all post data
/*func (p *Post) Fetch(w http.ResponseWriter, r *http.Request) {

	patchdata := &cities.RequestSelect{}

	err := json.NewDecoder(r.Body).Decode(patchdata)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	infoLog.Printf("Получаем аргумент %s", patchdata)
	if err == nil {

		//payload, _ := p.repo.Fetch(r.Context(), string(patchdata.Name))

		//	respondwithJSON(w, http.StatusOK, payload)
	}

}
func (p *Post) Create(w http.ResponseWriter, r *http.Request) {

	patchdata := &models.RequestSelect{}

	err := json.NewDecoder(r.Body).Decode(patchdata)
	if err == nil {

		payload, _ := p.repo.Create(r.Context(), patchdata)

		respondwithJSON(w, http.StatusOK, payload)
	}

}
func (p *Post) Update(w http.ResponseWriter, r *http.Request) {

	patchdata := &models.RequestMakeFriend{}

	err := json.NewDecoder(r.Body).Decode(patchdata)
	if err == nil {

		payload, _ := p.repo.Update(r.Context(), patchdata)

		respondwithJSON(w, http.StatusOK, payload)
	}

}
func (p *Post) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	_, err := p.repo.Delete(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Delete Successfully"})
}*/

// respondwithJSON write json response format
