package web

import (
	"net/http"
	"strings"
	"urlshort/database/postgres"
)

// Middleware checks if requested URL is a short link or not,
// if it is, it redirects to the original URL,
// if it's not, it serves the index web page.
func Middleware(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	path, _ = strings.CutPrefix(path, "/")

	db, err := postgres.New()
	if err != nil {
		indexHandler(w, r)

		return
	}
	defer db.Close()

	origin, err := db.GetOrigin(path)
	if err != nil {
		indexHandler(w, r)

		return
	}

	http.Redirect(w, r, origin, http.StatusFound)
}
