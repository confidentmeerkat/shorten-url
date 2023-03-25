package api

import (
	"encoding/json"
	"log"
	"net/http"
	"urlshort/database/postgres"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	db, err := postgres.New()
	if err != nil {
		w.WriteHeader(503)
		log.Fatal(err)
	}

	url, err := db.GetAll()
	if err != nil {
		w.WriteHeader(503)
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(url)

}
