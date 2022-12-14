package main

import (
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// NewMultipleHostReverseProxy creates a reverse proxy that will randomly
// select a host from the passed `targets`
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
	proxy := NewMultipleHostReverseProxy([]*url.URL{
		{
			Scheme: "http",
			Host:   "localhost:8002",
		},
		{
			Scheme: "http",
			Host:   "localhost:8003",
		},
	})
	log.Fatal(http.ListenAndServe(":8001", proxy))
}
