package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {

	data := url.Values{
		"name": {"Igor"},
	}

	resp, err := http.PostForm("localhost:8000/create", data)

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res["form"])
}
