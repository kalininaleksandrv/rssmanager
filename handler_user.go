package main

import (
	"net/http"
	"strconv"
	"strings"
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

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 || len(parts) > 4 {
		http.Error(w, "Invalid request URL", http.StatusBadRequest)
		return
	}
	
	id, err := strconv.Atoi(parts[2]) // Convert ID to int
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	fetchedUser, err := dbCfg.DB.GetUserById(r.Context(), int32(id))
	if err != nil {
		respondWithJson(w, http.StatusInternalServerError, map[string]string{"error": "Can't found user with id " + parts[2]})
		return
	}
	respondWithJson(w, http.StatusOK, fetchedUser)
}