package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/cyxabima/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't Create user: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {

	respondWithJSON(w, 200, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetPostForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")
	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		respondWithError(w, 400, "Invalid limit")
		return
	}
	intOffset, err := strconv.Atoi(offset)

	if err != nil {
		respondWithError(w, 400, "Invalid offset")
		return
	}

	posts, err := apiCfg.DB.GetPostForUser(r.Context(),
		database.GetPostForUserParams{
			UserID: user.ID,
			Limit:  int32(intLimit),
			Offset: int32(intOffset),
		})

	if err != nil {
		respondWithError(w, 400, err.Error())
	}

	respondWithJSON(w, 200, databasePostsToPosts(posts))

}
