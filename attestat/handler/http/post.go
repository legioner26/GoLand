package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	cities "skillbox/attestat/models"
	repository "skillbox/attestat/repository"
	"strconv"
	"unicode"
)

// NewPostHandler ...
func NewPostHandler() *Post {
	return &Post{}
}

// Post ...
type Post struct {
	repo repository.PostRepo
}

func (h *Post) Updates(w http.ResponseWriter, r *http.Request) {}

func (h *Post) getFull(w http.ResponseWriter, r *http.Request) {
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

// create processes a request to create a city.
// If successful, it writes a JSON response containing information
// about the creation of the city with id, status 201
func (h *Post) create(w http.ResponseWriter, r *http.Request) {
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

// delete processes a request to delete a city by id.
// If successful, it writes a response in JSON format
// containing information about the deletion of the city, status of 200
func (h *Post) delete(w http.ResponseWriter, r *http.Request) {
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

// setPopulation processes a request to update the population of the city by id.
// If successful, it writes a response in JSON format
// containing information about the city's population update, status 200
func (h *Post) setPopulation(w http.ResponseWriter, r *http.Request) {
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

// containing a list of cities, status 200
func (h *Post) getFromRegion(w http.ResponseWriter, r *http.Request) {
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

// getFromDistrict processes a request to get a list of cities by district.
// If successful, it writes a response in JSON format
// containing a list of cities, status 200
func (h *Post) getFromDistrict(w http.ResponseWriter, r *http.Request) {
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

// getFromPopulation processes a request to get a list of cities by population.
// If successful, it writes a response in JSON format
// containing a list of cities, status 200
func (h *Post) getFromPopulation(w http.ResponseWriter, r *http.Request) {
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

// getFromFoundation processes a request to get a list of cities by foundation.
// If successful, it writes a response in JSON format
// containing a list of cities, status 200
func (h *Post) getFromFoundation(w http.ResponseWriter, r *http.Request) {
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

// toUpperFirst changes the first letter to uppercase
// returns the string
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
