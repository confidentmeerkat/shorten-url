package api

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"urlshort/configs"
	"urlshort/database/postgres"
	"urlshort/types"
)

type create struct {
	Origin string
}

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

	_, err = url.ParseRequestURI(link.Origin)
	if err != nil {
		handleError(w, err, "not a valid URL")

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
