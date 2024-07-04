package main

import (
	"net/http"

	"github.com/joshua468/myapp/api"
)

func main() {
	http.HandleFunc("/api/hello", api.Handler)
	http.ListenAndServe(":3000", nil)
}
