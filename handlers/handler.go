package handlers

import (
	"net/http"
	"urlshort/handlers/api"
	"urlshort/handlers/web"
)

func Handle() {
	http.HandleFunc("/", web.Middleware)
	http.HandleFunc("/short", web.Short)

	http.HandleFunc("/assets/", web.ServeResource)

	http.HandleFunc("/api/create", api.CreateShort)
	http.HandleFunc("/api/get", api.GetShort)
	http.HandleFunc("/api/all", api.GetAll)
}
