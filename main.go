package main

import (
	"log"
	"net/http"

	"github.com/joshua468/myapp/api/handler"
)

func main() {
	http.HandleFunc("/api/hello", handler.Handler)
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
