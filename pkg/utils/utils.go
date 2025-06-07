package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading HTTP request body: %v", err)
		return
	}
	err = json.Unmarshal(body, x)
	if err != nil {
		log.Printf("Error unmarshalling JSON: %v", err)
		return
	}
}