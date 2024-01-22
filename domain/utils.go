package domain

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func BodyParser(r *http.Request) []byte {
	body, _ := io.ReadAll(r.Body)
	return body
}

func BodyToJson(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-type", "application/json; charset=UTF8")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	CheckErr(err)
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
