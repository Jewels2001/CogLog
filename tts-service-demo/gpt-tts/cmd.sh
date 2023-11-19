curl --location --request POST "https://centralus.tts.speech.microsoft.com/cognitiveservices/v1" \
--header "Ocp-Apim-Subscription-Key: ${COGLOG_SPEECH_KEY}" \
--header 'Content-Type: application/ssml+xml' \
--header 'X-Microsoft-OutputFormat: audio-16khz-128kbitrate-mono-mp3' \
--header 'User-Agent: curl' \
--data-raw '<speak version='\''1.0'\'' xml:lang='\''en-US'\''>
    <voice xml:lang='\''en-CA'\'' xml:gender='\''Male'\'' name='\''en-CA-LiamNeural'\''>
        Hi
    </voice>
</speak>'
