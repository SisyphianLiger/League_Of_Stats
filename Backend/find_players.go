package main

import (
	"net/http"
)

func (cfg *apiConfig) FindAllPlayers(w http.ResponseWriter, r *http.Request) {

	cfg.dbq.FindAllPlayerCards(r.Context())
}
