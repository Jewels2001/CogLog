package main

import (
	"io"
	"log"
	"net/http"

	"github.com/Jewels2001/CogLog/tts-service/tts"
	"github.com/Jewels2001/CogLog/tts-service/bcidecode"
)

const PORT = ":8080"

func main() {
	if err := tts.InitTTS(); err != nil {
		log.Fatal(err)
	}
	if err := bcidecode.InitBci(); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/bci-decode/audio", audio)
	log.Println("Serving on port " + PORT + "...")
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func audio(w http.ResponseWriter, r *http.Request) {
	log.Println("Request:", r.URL)

	// Enable CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		// For preflight requests, respond with a 200 OK status
		return
	}

	// Parse Request
	idx, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("main:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("idx:", idx)

	// Get Text from BCI data
	text, err := bcidecode.BciDecode(idx)
	if err != nil {
		log.Println("main:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
    log.Println("text:", text)

	// Convert to speech
	data, err := tts.GetAudioData(text, "en-CA", "Male", "en-CA-LiamNeural")
	if err != nil {
		log.Println("main:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
    log.Println("audio:", len(*data))

	// Return Data
	w.Header().Set("Content-Type", "audio/mp3")
	w.Write(*data)
}
