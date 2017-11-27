package controllers

import (
	"encoding/json"
	"net/http"
)

// ReplyWithJSON replies to the client with a JSON representation of i
func ReplyWithJSON(w http.ResponseWriter, i interface{}) {
	body, _ := json.Marshal(i)
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
