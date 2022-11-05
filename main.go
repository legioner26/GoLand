package main

import (
	"fmt"
	"github.com/goombaio/namegenerator"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", browser)
	http.ListenAndServe("localhost:8000", nil)
}
func browser(w http.ResponseWriter, r *http.Request) {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)
	name := nameGenerator.Generate()
	fmt.Fprintf(w, "Привет %s", name)
}
