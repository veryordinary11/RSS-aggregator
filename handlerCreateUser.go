package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/veryordinary11/RSS-aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		Name: params.Name,
		ID:   uuid.New(),
	})
}
