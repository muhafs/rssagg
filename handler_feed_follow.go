package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/muhamadhaspin/rssagg/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't createe feed follow: %v", err))
		return
	}

	respondWithJSON(w, 201, FeedFollowFromDB(feedFollow))
}

func (apiCfg *apiConfig) handlerListFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiCfg.DB.ListFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get feed follows: %v", err))
		return
	}

	respondWithJSON(w, 200, FeedFollowsFromDB(feedFollows))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing feed follow ID: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't delete feed follow: %v", err))
		return
	}

	respondWithJSON(w, 200, struct{}{})
}
