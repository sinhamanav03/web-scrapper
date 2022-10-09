package utils

import (
	"encoding/json"
	"net/http"
)

func DecodeRequest(r *http.Request, val interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(val); err != nil {
		return err
	}
	return nil
}

func JsonResponse(w http.ResponseWriter, statusCode int, resp interface{}) {
	response, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}
