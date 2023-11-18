package main

import (
	"fmt"
	"log"

    "github.com/dustin-ward/gpt-tts-demo/gpt"
    "github.com/dustin-ward/gpt-tts-demo/tts"
)

func main() {
	if err := gpt.InitGPT(); err != nil {
		log.Fatal(err)
	}
	if err := tts.InitTTS(); err != nil {
		log.Fatal(err)
	}

	content, err := gpt.GetText()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(content)

    if err = tts.TTS(content); err != nil {
        log.Fatal(err)
    }
}
