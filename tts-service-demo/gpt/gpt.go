package gpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Message map[string]string

var token string
var req *http.Request

func InitGPT() error {
	// Get API token
	token := os.Getenv("COGLOG_GPT_TOKEN")
	if token == "" {
		return fmt.Errorf("InitGPT: no OpenAI api token set in $COGLOG_GPT_TOKEN")
	}

	// Build Prompt
	data := map[string]any{
		"model": "gpt-3.5-turbo",
		"messages": []Message{
			Message{"role": "system", "content": "You are a salemen working for Taro Noodlehouse Corp"},
			Message{"role": "user", "content": "Come up with a brief pitch for Taro Noodlehouse Corps newest product 'CogLog'. It transforms brainwave data into TTS audio. Please respond with no more than 2 sentences"},
		},
	}

	// Marshal
	prompt, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("InitGPT: %v", err)
	}

	// Build request
	req, err = http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(prompt))
	if err != nil {
		return fmt.Errorf("InitGPT: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	return nil
}

func GetText() (string, error) {
	// Make request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("GPT: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return "", fmt.Errorf("GPT: %v", resp.Status)
	}

	// Return response
	var res map[string]any
	json.NewDecoder(resp.Body).Decode(&res)

	choices := (res["choices"].([]any))[0].(map[string]any)
	message := choices["message"].(map[string]any)
	content := message["content"].(string)
	content = strings.Replace(content, "\"", "", -1)

	return content, nil
}
