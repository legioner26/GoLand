package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"os"
	"skillbox/30-31/driver"
	ph "skillbox/30-31/handler/http"
)

func main() {
	os.Setenv("DB_NAME", "testgo")
	os.Setenv("DB_PASS", "")
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "3306")

	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	connection, err := driver.ConnectSQL(dbHost, dbPort, "root", dbPass, dbName)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	pHandler := ph.NewPostHandler(connection)
	r.Route("/", func(rt chi.Router) {
		rt.Mount("/30-31", postRouter(pHandler))
	})

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	errorLog.Fatal(http.ListenAndServe(":8000", r))

}

// A completely separate router for posts routes
func postRouter(pHandler *ph.Post) http.Handler {
	r := chi.NewRouter()
	r.Post("/select", pHandler.Fetch)
	r.Post("/create", pHandler.Create)
	r.Delete("/delete/{id:[0-9]+}", pHandler.Delete)
	r.Put("/update", pHandler.Update)
	/*r.Get("/{id:[0-9]+}", pHandler.GetByID)
	r.Post("/", pHandler.Create)
	r.Put("/{id:[0-9]+}", pHandler.Update)
	r.Delete("/{id:[0-9]+}", pHandler.Delete)*/

	return r
}
