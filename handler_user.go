package main

import (
	"net/http"
)

func (dbCfg *dbConfig) handlerCreateUser (w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, http.StatusOK, map[string]string{"message": "Hello from RssManager backend!"})
}

func (dbCfg *dbConfig) handlerGetUser (w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, http.StatusOK, map[string]string{"message": "Hello from RssManager backend!"})
}