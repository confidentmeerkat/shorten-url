package handlers

import (
	"net/http"
	"urlshort/handlers/api"
	"urlshort/handlers/web"
)

func Handle() {
	http.HandleFunc("/", web.Handler)
	http.HandleFunc("/short", web.Short)

	http.HandleFunc("/assets/", web.ServeResource)

	http.HandleFunc("/api/get", api.GetShort)
	http.HandleFunc("/api/all", api.GetAll)
}
