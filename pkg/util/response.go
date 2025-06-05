package util

import (
	"encoding/json"
	"net/http"
)

func Response(w http.ResponseWriter, responses interface{}, statusCode int) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(responses)
	if err != nil {
		panic(err)
	}
}
