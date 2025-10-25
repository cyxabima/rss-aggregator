package main

import "net/http"

func handlerReadinessHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{}) // anonymous empty struct
}
