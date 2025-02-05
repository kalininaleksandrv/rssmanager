package main

import (
	"log"
	"net/http"

	"github.com/kalininaleksandrv/rssmanager/internal/database"
)

func (dbCfg *dbConfig) handlerCreateFeed (w http.ResponseWriter, r *http.Request, user database.User) {

	feed, err := parseFeedJsonRequest(r)
	if err != nil {
		respondWithJson(w, http.StatusBadRequest, map[string]string{"error": "Invalid request: unable to parse Feed params"})
		return
	}
	createdFeed, err := dbCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		Name:      feed.Name,
        Url:	   feed.Url,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithJson(w, http.StatusInternalServerError, map[string]string{"error": "Failed to create feed with name " + feed.Name})
		log.Println(err)
		return
	}
	respondWithJson(w, http.StatusCreated, createdFeed)
}
