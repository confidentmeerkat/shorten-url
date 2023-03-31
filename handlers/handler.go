// Package handlers serves HTTP requests.
package handlers

import (
	"net/http"
	"urlshort/handlers/api"
	"urlshort/handlers/web"
)

// Handle handles HTTP routes.
func Handle() {
	http.HandleFunc("/", web.Middleware)
	http.HandleFunc("/short", web.ShortHandler)
	http.HandleFunc("/api", web.APIGuide)

	http.HandleFunc("/assets/", web.ServeResource)

	http.HandleFunc("/api/create", api.CreateShort)
	http.HandleFunc("/api/get", api.GetShort)
	http.HandleFunc("/api/all", api.GetAll)
}
