package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type payloadUser struct {
	Name string `json:"name"`
}

type payloadFeed struct {
	Name string `json:"name"`
	Url string `json:"url"`
	UserID int32 `json:"user_id"`
}

func parseUserJsonRequest(r *http.Request) (payloadUser, error) {

	payloadObj := payloadUser{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payloadObj)
	if err != nil {
		log.Printf("Error decoding request, %v", err)
		return payloadObj, err
	}
	return payloadObj, nil
}

func parseFeedJsonRequest(r *http.Request) (payloadFeed, error) {

	payloadObj := payloadFeed{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payloadObj)
	if err != nil {
		log.Printf("Error decoding request, %v", err)
		return payloadObj, err
	}
	return payloadObj, nil
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling response, %v", payload)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	addHeaders(w)
	w.WriteHeader(code)
	w.Write(response)
}

func addHeaders(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	return w
}

