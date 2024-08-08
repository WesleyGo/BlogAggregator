package main

import (
	"encoding/json"
	"net/http"
	"time"

	"strings"

	"github.com/WESLEYGO/BlogAggregator/internal/database"
	httpHelper "github.com/WESLEYGO/BlogAggregator/internal/http"
	"github.com/google/uuid"
)

type apiConfig struct {
	DB *database.Queries
}

func (cfg *apiConfig) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		httpHelper.RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		Name:      params.Name,
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		httpHelper.RespondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}

	httpHelper.RespondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}

func (cfg *apiConfig) handlerUsersGet(w http.ResponseWriter, r *http.Request) {
	apikey := r.Header.Get("Authorization")

	apikey = strings.TrimPrefix(apikey, "APIKey ")

	user, err := cfg.DB.GetUserByAPIKey(r.Context(), apikey)

	if err != nil {
		httpHelper.RespondWithError(w, http.StatusInternalServerError, "Couldn't get user")
		return
	}

	httpHelper.RespondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}
