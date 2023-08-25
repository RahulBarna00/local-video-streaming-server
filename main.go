package main

import (
	"net/http"
)

func main() {
	// Serve index.html at the root
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Handle video streaming
	http.HandleFunc("/videos/", func(w http.ResponseWriter, r *http.Request) {
		videoPath := "videos/" + r.URL.Path[len("/videos/"):]
		http.ServeFile(w, r, videoPath)
	})

	ipAddress := "localhost" // Replace with your desired IP address
	port := ":8080"
	address := ipAddress + port

	http.ListenAndServe(address, nil)
}
