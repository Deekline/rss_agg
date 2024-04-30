package main

import (
	"net/http"
)

func handler_readiness(w http.ResponseWriter, r *http.Request) {
	responseWithJSON(w, http.StatusOK, struct{}{})
}
