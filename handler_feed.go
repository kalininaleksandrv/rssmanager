package main

import (
	"log"
	"net/http"
	"time"

	"github.com/kalininaleksandrv/rssmanager/internal/database"
)

func (dbCfg *dbConfig) handlerCreateFeed (w http.ResponseWriter, r *http.Request, user database.User) {

	feed, err := parseFeedJsonRequest(r)
	if err != nil {
		respondWithJson(w, http.StatusBadRequest, map[string]string{"error": "Invalid request: unable to parse Feed params"})
		return
	}

	tx, err := dbCfg.DBSQL.Begin()

	if err != nil {
        respondWithJson(w, http.StatusInternalServerError, map[string]string{"error": "Failed to begin transaction"})
        return
    }

	defer tx.Rollback()

	qtx := dbCfg.DB.WithTx(tx)

	createdFeed, err := qtx.CreateFeed(r.Context(), database.CreateFeedParams{
		Name:      feed.Name,
        Url:	   feed.Url,
		UserID:    user.ID,
	})

	err2 := qtx.UpdateUserCounter(r.Context(), database.UpdateUserCounterParams{
		ID:        user.ID,
		UpdatedAt: time.Now(),
		Counter:   user.Counter + 1,
	})

	if err != nil || err2 != nil {
		respondWithJson(w, http.StatusInternalServerError, map[string]string{"error": "Failed to create feed with name " + feed.Name})
		log.Println(err)
		return
	}

	if err := tx.Commit(); err != nil {
        respondWithJson(w, http.StatusInternalServerError, map[string]string{"error": "Failed to commit transaction"})
        log.Println(err)
        return
    }

	respondWithJson(w, http.StatusCreated, createdFeed)
}
