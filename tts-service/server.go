// Example streaming server
package main

import (
	_ "embed"
	"log"
	"net/http"
)

//go:embed test.mp3
var test_mp3 []byte

const PORT = ":8080"

func main() {
	http.HandleFunc("/test.mp3", StreamTestMP3)
	log.Println("Serving on localhost" + PORT + "...")
	http.ListenAndServe(PORT, nil)
}

func StreamTestMP3(w http.ResponseWriter, r *http.Request) {
	log.Println("Request:", r.URL)

    // Enable CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		// For preflight requests, respond with a 200 OK status
		return
	}

    // Return Data
    w.Header().Set("Content-Type", "audio/mp3")
	w.Write(test_mp3)
}
