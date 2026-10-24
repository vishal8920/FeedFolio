package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/FilledEther/rss_reader/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (apiConfig *apiConfig) handlerCreateFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing json %s", err))
	}

	feedfollows, err := apiConfig.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create feeds to follow: %s", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowtoFeedFollow(feedfollows))
}


func (apiConfig *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing json %s", err))
	}

	feedfollows, err := apiConfig.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get feed follows: %s", err))
		return
	}

	respondWithJSON(w, 200, databaseFeedFollowstoFeedFollows(feedfollows))
}


func (apiConfig *apiConfig) handlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	//parameter is required by to delete the feed followed
	feed_param := chi.URLParam(r, "{feed_follow_id}")
	feed_id, err := uuid.Parse(feed_param)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't parse feed follows %v", err))
		return
	}

	err = apiConfig.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feed_id,
		UserID: user.ID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't delete feed follows %v", err))
		return
	}
	respondWithJSON(w, 200, struct{}{})
}
