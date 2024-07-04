package main

import (
	"log"
	"net/http"

   "github.com/joshua468/myapp/api"
)

func main() {
    http.HandleFunc("/api/hello", handler.Handler)

    log.Println("Server is running on port 8081")
    if err := http.ListenAndServe(":8081", nil); err != nil {
        log.Fatalf("could not start server: %s\n", err)
    }
}
