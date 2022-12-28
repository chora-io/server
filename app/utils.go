package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// compactJSONString removes unnecessary whitespace from a json string
func compactJSONString(str string) string {
	var bz = []byte(str)
	buf := new(bytes.Buffer)
	if err := json.Compact(buf, bz); err != nil {
		fmt.Println(err)
	}
	return buf.String()
}

// respondJSON makes the response with payload as json format
func respondJSON(w http.ResponseWriter, status int, payload any) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error())) //nolint
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response) //nolint
}

// respondError makes the error response with payload as json format
func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}
