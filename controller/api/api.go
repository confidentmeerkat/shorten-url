package api

import (
	"encoding/json"
	"log"
	"net/http"
	"urlshort/database/postgres"
	"urlshort/types"
)

func GetShort(w http.ResponseWriter, r *http.Request) {
	db, err := postgres.New()
	if err != nil {
		customErr := types.Error{Err: "service unavailable"}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customErr)

		log.Println(err)

		return
	}

	origin := r.URL.Query().Get("origin")
	short := r.URL.Query().Get("short")

	if origin != "" {
		shortUrl, err := db.GetShort(origin)
		if err != nil {
			customErr := types.Error{Err: "not found"}

			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(customErr)

			log.Println(err)

			return
		}

		res := types.Url{Origin: origin, Short: shortUrl}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	} else if short != "" {
		originURL, err := db.GetOrigin(short)
		if err != nil {
			customErr := types.Error{Err: "not found"}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(customErr)

			return
		}

		res := types.Url{Origin: originURL, Short: short}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
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
