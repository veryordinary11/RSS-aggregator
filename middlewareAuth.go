package main

import (
	"net/http"

	"github.com/veryordinary11/RSS-aggregator/internal/auth"
	"github.com/veryordinary11/RSS-aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			responseWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			responseWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		handler(w, r, user)
	}
}
