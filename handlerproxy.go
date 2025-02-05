package main

import (
	"net/http"
	"strconv"

	"github.com/kalininaleksandrv/rssmanager/internal/database"
)


type proxyHandler func(http.ResponseWriter, *http.Request, database.User)

func (dbCfg *dbConfig) handlerUserProxy (handler proxyHandler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		
	id, err:= extractUserIDFromURL(r)

	if err != nil {
		respondWithJson(w, http.StatusBadRequest, map[string]string{"error": "Invalid request: unable to parse User params"})
		return
	}

	fetchedUser, err := dbCfg.DB.GetUserById(r.Context(), int32(id))
	    if err != nil {
			respondWithJson(w, http.StatusInternalServerError, map[string]string{"error": "Can't found user with id " + strconv.Itoa(id)})
		    return
	    }
	    handler(w, r, fetchedUser)
    }
}