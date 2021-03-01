package helper

import (
	"encoding/json"
	"log"
)

// MessageParser ...
func MessageParser(message string, dto interface{}) {
	byteMessage := []byte(message)

	err := json.Unmarshal(byteMessage, dto)

	if err != nil {
		log.Fatal(err)
	}
}
