package main

import (
	"fmt"
	"net/http"

	"github.com/Taras-Rm/rss-agg/internal/auth"
	"github.com/Taras-Rm/rss-agg/internal/database"
)

type authedHandler func(w http.ResponseWriter, r *http.Request, user database.User)

// middlewareAuth - auth middleware
func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Api key error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Can not get user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
