package main

import (
	"net/http"

	"github.com/joshua468/myapp/api"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	api.Handler(w, r)
}

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}
