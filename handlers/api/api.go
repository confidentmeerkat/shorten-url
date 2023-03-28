// Package api serves API requests.
package api

import (
	"encoding/json"
	"log"
	"net/http"
	"urlshort/configs"
	"urlshort/database/postgres"
	"urlshort/pkg"
	"urlshort/types"
)

type create struct {
	Origin string
}

// CreateShort gets a URL in JSON on the POST method and creates a short link.
func CreateShort(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method is not allowed"))

		return
	}

	var link create

	err := json.NewDecoder(r.Body).Decode(&link)
	if err != nil {
		handleError(w, err, "can't parse json")

		return
	}

	if !pkg.IsValidURL(link.Origin) {
		handleError(w, err, "no valid URL")

		return
	}

	db, err := postgres.New()
	if err != nil {
		handleError(w, err, "service unavailable")

		return
	}

	res := types.URL{}

	res.Short, err = db.GetShort(link.Origin)
	if err != nil {
		res.Short, err = db.CreateShort(link.Origin, 4)
		if err != nil {
			handleError(w, err, "creating short link failed")

			return
		}
	}

	res.Origin = link.Origin
	res.Short = configs.Domain + "/" + res.Short

	w.Header().Set("Content-Type", "applicaton/json")
	json.NewEncoder(w).Encode(res)
}

// GetShort returns a short link or original URL in JSON format to the user.
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

// GetAll returns all URLs and their short link in JSON format to the user.
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

// handleError creates a custom error in JSON format.
func handleError(w http.ResponseWriter, err error, customError string) {
	customErr := types.Error{Err: customError}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customErr)

	log.Println(err)
}
