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
		handleError(w, err, "service unavailable")

		return
	}

	origin := r.URL.Query().Get("origin")
	short := r.URL.Query().Get("short")

	if origin != "" {
		shortUrl, err := db.GetShort(origin)
		if err != nil {
			handleError(w, err, "not found")

			return
		}

		res := types.URL{Origin: origin, Short: shortUrl}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	} else if short != "" {
		originURL, err := db.GetOrigin(short)
		if err != nil {
			handleError(w, err, "not found")

			return
		}

		res := types.URL{Origin: originURL, Short: short}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	} else {
		err := types.Error{Err: "no acceptable query"}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
	}
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	db, err := postgres.New()
	if err != nil {
		handleError(w, err, "service unavailable")

		return
	}

	url, err := db.GetAll()
	if err != nil {
		handleError(w, err, "not found")

		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(url)

}

func handleError(w http.ResponseWriter, err error, customError string) {
	customErr := types.Error{Err: customError}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customErr)

	log.Println(err)
}
