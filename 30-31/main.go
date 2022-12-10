package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"skillbox/30-31/driver"
	ph "skillbox/30-31/handler/http"
)

func NewMultipleHostReverseProxy(targets []*url.URL) *httputil.ReverseProxy {
	director := func(req *http.Request) {
		target := targets[rand.Int()%len(targets)]
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = target.Path
	}
	return &httputil.ReverseProxy{Director: director}
}
func main() {
	dbName := os.Getenv("testgo")
	dbPass := os.Getenv("")
	dbHost := os.Getenv("127.0.0.1")
	dbPort := os.Getenv("3306")

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
		rt.Mount("/posts", postRouter(pHandler))
	})

	fmt.Println("Server listen at :8000")

	http.ListenAndServe(":8000", r)
}

// A completely separate router for posts routes
func postRouter(pHandler *ph.Post) http.Handler {
	r := chi.NewRouter()
	r.Post("/select", pHandler.Fetch)
	/*r.Get("/{id:[0-9]+}", pHandler.GetByID)
	r.Post("/", pHandler.Create)
	r.Put("/{id:[0-9]+}", pHandler.Update)
	r.Delete("/{id:[0-9]+}", pHandler.Delete)*/

	return r
}
