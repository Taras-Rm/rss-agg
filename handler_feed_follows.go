package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Taras-Rm/rss-agg/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feedId"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Can not get parameters: %v", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedId,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Can not create feed follow: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Can not create feed follow: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowsToFeedFollows(feedFollows))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIdStr := chi.URLParam(r, "feedFollowId")
	if feedFollowIdStr == "" {
		respondWithError(w, 400, "Empty feed follow id")
		return
	}

	feedFollowId, err := uuid.Parse(feedFollowIdStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Can not parse feed follow id: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowId,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Can not delete feed follow: %v", err))
		return
	}

	respondWithJSON(w, 201, struct{}{})
}
