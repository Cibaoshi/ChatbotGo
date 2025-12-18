package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/chat", chatHandler)

	log.Println("Server listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
