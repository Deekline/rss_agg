package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/deekline/rss_agg/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: params.FeedID,
	})
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, fmt.Sprintf("Could not create follow feed: %s", err))
	}

	responseWithJSON(w, http.StatusCreated, databaseFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollow, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, fmt.Sprintf("Could not get follow feed: %s", err))
	}

	responseWithJSON(w, http.StatusCreated, databaseFeedFollowsToFeedFollows(feedFollow))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedUUID := chi.URLParam(r, "feedFollowID")
	feedId, err := uuid.Parse(feedUUID)
	if err!= nil {
        responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Could not parse feed id: %s", err))
        return
    }
	
	err = apiCfg.DB.DeleteFeedFollows(r.Context(), database.DeleteFeedFollowsParams{
		ID: feedId,
		UserID: user.ID,
	})
	if err!= nil {
		responseWithError(w, http.StatusInternalServerError, fmt.Sprintf("Could not delete follow feed: %s", err))
		return
	}

	responseWithJSON(w, http.StatusOK, struct{}{})
}


