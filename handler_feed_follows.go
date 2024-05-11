package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/armalam/go-freecodecamp/internal/database"
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
		responseWithError(w, 400, fmt.Sprintf("Err parsing JSON %v", err))

		return

	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Cant create feed follow %v", err))

		return
	}

	responseWithJSON(w, 201, databaseFeedFollowToFeedFollow(feedFollow))

}
