package helper

import (
	"encoding/json"
	"net/http"
)

func ResponeJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
