package main

import (
	"net/http"
)

func handler_error(w http.ResponseWriter, r *http.Request) {
	responseWithError(w, http.StatusBadRequest, "Something went wrong")
}
