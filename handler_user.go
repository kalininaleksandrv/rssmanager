package main

import (
	"net/http"
	"strconv"
	"strings"
	"time"

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
	respondWithJson(w, http.StatusCreated, createdUser)
}

func (dbCfg *dbConfig) handlerGetUser (w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		respondWithJson(w, http.StatusBadRequest, map[string]string{"error": "Invalid request URL"})
		return
	}
	
	id, err := strconv.Atoi(parts[2]) // Convert ID to int
	if err != nil {
		respondWithJson(w, http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
		return
	}
	fetchedUser, err := dbCfg.DB.GetUserById(r.Context(), int32(id))
	if err != nil {
		respondWithJson(w, http.StatusInternalServerError, map[string]string{"error": "Can't found user with id " + parts[2]})
		return
	}
	respondWithJson(w, http.StatusOK, fetchedUser)
}

func (dbCfg *dbConfig) handlerUpdateUser (w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		respondWithJson(w, http.StatusBadRequest, map[string]string{"error": "Invalid request URL"})
		return
	}
	
	id, err := strconv.Atoi(parts[2]) // Convert ID to int
	if err != nil {
		respondWithJson(w, http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
		return
	}
	user, err := parseJsonRequest(r)
	if err != nil {
		respondWithJson(w, http.StatusBadRequest, map[string]string{"error": "Invalid request: unable to parse User params"})
		return
	}
	updatedUser, err := dbCfg.DB.UpdateUser(r.Context(), database.UpdateUserParams{
		ID:        int32(id),
		Name:      user.Name,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		respondWithJson(w, http.StatusInternalServerError, map[string]string{"error": "Failed to update user with id " + parts[2]})
		return
	}
	respondWithJson(w, http.StatusOK, updatedUser)
}