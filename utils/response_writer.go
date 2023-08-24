package utils

import (
	"Task-Management/constants"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// JsonWriter : Return Response for string format
func JsonWriterString(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(msg))
}

// JsonWriter1 : Return Response for []byte format
func JsonWriterBytes(w http.ResponseWriter, status int, msg []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(msg)
}

func WriteResponseStruct(ctx context.Context, w http.ResponseWriter, status int, response interface{}) {
	Outjson, err := json.Marshal(response)
	if err != nil {
		JsonWriterString(w, http.StatusBadRequest, fmt.Sprintf(`{"status":"failed","message":"%s"}`, constants.MarshalError))
		return
	}
	JsonWriterBytes(w, status, Outjson)
}
