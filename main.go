package main

import (
	"log"
	"net/http"

	handler "github.com/joshua468/myapp/api"
)

func main() {
	http.HandleFunc("/api/hello", handler.Handler)

	port := ":8081"
	log.Println("Server is running on port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}
