package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"os"

	ph "skillbox/attestat/handler/http"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	pHandler := ph.NewPostHandler()
	r.Route("/", func(rt chi.Router) {
		rt.Mount("/attestat", postRouter(pHandler))
	})

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	errorLog.Fatal(http.ListenAndServe(":8000", r))

}

// A completely separate router for posts routes
func postRouter(pHandler *ph.Post) http.Handler {
	r := chi.NewRouter()
	r.Get("/{id}", pHandler.GetFull)
	r.Post("/create", pHandler.Create)
	r.Delete("/{id}", pHandler.Delete)
	r.Put("/population/{id}", pHandler.SetPopulation)
	r.Put("/update", pHandler.Update)
	r.Get("/region/{region}", pHandler.GetFromRegion)
	r.Get("/district/{district}", pHandler.GetFromDistrict)
	r.Get("/population/range", pHandler.GetFromPopulation)
	r.Get("/foundation/range", pHandler.GetFromFoundation)

	return r
}