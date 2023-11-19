package bcidecode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var addr string

func InitBci() error {
	addr = os.Getenv("COMPUTE_ADDR")
	if addr == "" {
		return fmt.Errorf("InitBci: no compute addr set in $COMPUTE_ADDR")
	}
	return nil
}

func BciDecode(idx []byte) (string, error) {
	// Build request
	req, err := http.NewRequest("POST", addr, bytes.NewBuffer(idx))
	if err != nil {
		return "", fmt.Errorf("BciDecode: %v", err)
	}

	// Make request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("BciDecode: %v", err)
	}
	defer resp.Body.Close()

	txt := struct {
		Text string `json:"text"`
	}{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&txt); err != nil {
		return "", fmt.Errorf("BciDecode: %v", err)
	}

	return txt.Text, nil
}
