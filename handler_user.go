package main

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/kalininaleksandrv/rssmanager/internal/database"
)



func (dbCfg *dbConfig) handlerCreateUser (w http.ResponseWriter, r *http.Request) {

	user, err := parseUserJsonRequest(r)
	if err != nil {
		respondWithJson(w, http.StatusBadRequest, map[string]string{"error": "Invalid request: " + err.Error()})
		return
	}
	createdUser, err := dbCfg.DB.CreateUser(r.Context(), user.Name)
	if err != nil {
		respondWithJson(w, http.StatusInternalServerError, map[string]string{"error": "Failed to create user with name " + user.Name})
		return
	}
	respondWithJson(w, http.StatusCreated, createdUser)
}

func (cfg *dbConfig) handlerUsersGet(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJson(w, http.StatusOK, user)
}

func (dbCfg *dbConfig) handlerUpdateUser (w http.ResponseWriter, r *http.Request) {

	id, err:= extractUserIDFromURL(r)

	if err != nil {
		respondWithJson(w, http.StatusBadRequest, map[string]string{"error": "Invalid request: unable to parse User params"})
		return
	}

	user, err := parseUserJsonRequest(r)
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
		respondWithJson(w, http.StatusInternalServerError, map[string]string{"error": "Failed to update user with id " + strconv.Itoa(id)})
		return
	}
	respondWithJson(w, http.StatusOK, updatedUser)
}

func extractUserIDFromURL(r *http.Request) (int, error) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		return 0, errors.New("invalid request URL")
	}
	
	id, err := strconv.Atoi(parts[2]) // Convert ID to int
	if err != nil {
		return 0, errors.New("invalid user ID")
	}
	return id, nil
}