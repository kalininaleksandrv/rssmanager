package main

import (
	"database/sql"
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

	defer tx.Rollback() //how this is not works if tx.Commit() is called?

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

func (dbCfg *dbConfig) handlerFetchAllFeeds (w http.ResponseWriter, r *http.Request) {

	delay := -10 * time.Minute

	feeds, err := dbCfg.DB.GetFeedsForFetchUpdate(r.Context(), sql.NullTime{Time: time.Now().Add(delay), Valid: true})
	if err != nil {
		respondWithJson(w, http.StatusInternalServerError, map[string]string{"error": "No feeds to fetch"})
		return
	}
	sliceOfFeeds := []string{}
	for _, feed := range feeds {

		rssFeed, err := urlToFeed(feed.Url)
		if err != nil {
			log.Println("Failed to fetch feed: ", feed.Url)
		} else {
			_, err := dbCfg.DB.UpdateFeedLastFetch(r.Context(), database.UpdateFeedLastFetchParams{
				ID:            feed.ID,
				LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
			})
			if err != nil {
				log.Println("Failed to update feed: ", feed.ID)
			}
			print(rssFeed.Channel.Title)
			sliceOfFeeds = append(sliceOfFeeds, feed.Name)
		}

	}
	respondWithJson(w, http.StatusOK, map[string][]string{"Feeds being updated": sliceOfFeeds})
}
