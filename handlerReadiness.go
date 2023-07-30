package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	responseWithJSON(w, http.StatusOK, map[string]string{"status": "OK"})
}

func handlerErr(w http.ResponseWriter, r *http.Request) {
	responseWithError(w, http.StatusInternalServerError, "Internal server error")
}
