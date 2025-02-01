package main

import (
	"net/http"

	"github.com/kalininaleksandrv/rssmanager/internal/database"
)

func (dbCfg *dbConfig) handlerCreateUser (w http.ResponseWriter, r *http.Request) {

	user, err := parseJsonRequest(r)
	if err != nil {
		respondWithJson(w, http.StatusBadRequest, map[string]string{"error": "Invalid request: unable to parse User params"})
		return
	}
	createdUser, err := dbCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		Name:      user.Name,
	})
	if err != nil {
		respondWithJson(w, http.StatusInternalServerError, map[string]string{"error": "Failed to create user with name " + user.Name})
		return
	}
	respondWithJson(w, http.StatusOK, createdUser)
}

func (dbCfg *dbConfig) handlerGetUser (w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, http.StatusOK, map[string]string{"message": "Hello from RssManager backend!"})
}