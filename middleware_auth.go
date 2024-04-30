package main

import (
	"net/http"

	"github.com/deekline/rss_agg/internal/auth"
	"github.com/deekline/rss_agg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(h authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			responseWithError(w, http.StatusBadRequest, "Could not get API key")
			return
		}
		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			responseWithError(w, http.StatusBadRequest, "Could not get user")
			return
		}
		h(w, r, user)
	}
}
