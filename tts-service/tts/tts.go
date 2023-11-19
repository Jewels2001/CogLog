package tts

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

var token string

func InitTTS() error {
	// Get API token
	token = os.Getenv("COGLOG_SPEECH_KEY")
	if token == "" {
		return fmt.Errorf("InitTTS: no Azure speech key set in $COGLOG_SPEECH_KEY")
	}

	return nil
}

func GetAudioData(content, lang, gender, voice string) (*[]byte, error) {
	// Build Body
	data := fmt.Sprintf("<speak version='1.0' xml:lang='%s'><voice xml:lang='%s' xml:gender='%s' name='%s'>%s</voice></speak>",
		lang, lang, gender, voice, content)

	// Build request
	req, err := http.NewRequest("POST", "https://centralus.tts.speech.microsoft.com/cognitiveservices/v1", bytes.NewBuffer([]byte(data)))
	if err != nil {
		return nil, fmt.Errorf("TTS(getAudioData): %v", err)
	}
	req.Header.Add("Ocp-Apim-Subscription-Key", token)
	req.Header.Add("Content-Type", "application/ssml+xml")
	req.Header.Add("X-Microsoft-OutputFormat", "audio-16khz-128kbitrate-mono-mp3")

	// Make request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("TTS(getAudioData): %v", err)
	}
	defer resp.Body.Close()

	audio, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("TTS(getAudioData): %v", err)
	}

	return &audio, nil
}
