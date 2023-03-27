package api

import (
	"encoding/json"
	"log"
	"net/http"
	"urlshort/database/postgres"
	"urlshort/types"
)

func GetShort(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")

	db, err := postgres.New()
	if err != nil {
		customErr := types.Error{Err: "service unavailable"}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customErr)

		log.Println(err)

		return
	}

	short, err := db.GetShort(url)
	if err != nil {
		customErr := types.Error{Err: "not found"}

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customErr)

		log.Println(err)

		return
	}

	res := types.Url{Origin: url, Short: short}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

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

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(url)

}
