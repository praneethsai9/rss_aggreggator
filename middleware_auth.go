package main

import (
	"fmt"
	"net/http"

	"github.com/praneethsai9/rss_aggregator/internal/auth"
	"github.com/praneethsai9/rss_aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Could't create user:%v", err))
			return
		}
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 404, fmt.Sprintf("Couldn't get user:%v", err))
		}
		respondWithJSON(w, 200, databaseUserToUser(user))
		handler(w, r, user)
	}

}
