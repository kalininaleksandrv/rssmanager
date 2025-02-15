package main

import (
	"database/sql"
	"log"
	"net/http"
	"sync"
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
	waitGroup := &sync.WaitGroup{}

	for _, feed := range feeds {

		waitGroup.Add(1)
		// new struct for the clean way to get the result of fetching the feed
		type feedResult struct {
			rssFeed RssFeed
			err     error
		}
		//we use channel to get the result of fetching the feed
		resultChan := make(chan feedResult)

		go func() {
			//here we fecching the feed
			rssFeed, err := urlToFeed(feed.Url, waitGroup)

			if err == nil {
				_, err := dbCfg.DB.UpdateFeedLastFetch(r.Context(), database.UpdateFeedLastFetchParams{
					ID:            feed.ID,
					LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
				})
				if err != nil {
					log.Println("Failed to update feed: ", feed.ID)
				}
				print(rssFeed.Channel.Title)
			}
			//here we send the result to the channel
			resultChan <- feedResult{rssFeed, err}
		}()

		//here we get the result from the channel
		result := <-resultChan
		rssFeed := result.rssFeed
		err := result.err

		if(err != nil) {
			log.Println("Failed to fetch feed: ", feed.Url)
		}
		sliceOfFeeds = append(sliceOfFeeds, rssFeed.Channel.Title)
	}
	waitGroup.Wait()
	respondWithJson(w, http.StatusOK, map[string][]string{"Feeds being updated": sliceOfFeeds})
}
