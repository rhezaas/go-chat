package helper

import (
	"encoding/json"
	"log"
)

// ToString ...
func ToString(data interface{}) string {
	byt, err := json.Marshal(data)

	if err != nil {
		log.Fatal(err)
	}

	return string(byt)
}
