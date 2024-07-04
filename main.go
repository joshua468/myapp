package main

import (
	"log"
	"net/http"

	api "github.com/joshua468/myapp/api"
)

func main() {
	http.HandleFunc("/api/hello", api.Handler)

	port := ":8081" // Define the port
	log.Println("Server is running on port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}
