package main

import (
	"log"
	"net/http"

	"asciiart/serv"
)

func main() {
	// This will serve as the homepage for the application
	http.HandleFunc("/", serv.Index)

	// This endpoint will likely handle the ASCII art generation and serve the related output
	http.HandleFunc("/ascii-art", serv.AsciiWeb)

	// Serve static files from the "static" directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Log a message to the console to inform that the server is running and accessible at localhost:8080
	log.Println("Server running at http://localhost:8080")

	// Start the HTTP server on port 8080 and listen for incoming requests
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
