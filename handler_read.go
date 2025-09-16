package main

import "net/http"

func handler_read(w http.ResponseWriter, r *http.Request) {

	respondWithJSON(w, 400, "smoething went wrong")
}
