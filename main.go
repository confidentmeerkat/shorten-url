package main

import (
	"log"
	"net/http"
	"urlshort/handlers/api"
	"urlshort/handlers/web"
)

func main() {
	http.HandleFunc("/", web.Handler)
	http.HandleFunc("/short", web.Short)

	http.HandleFunc("/assets/", web.ServeResource)

	http.HandleFunc("/api/get", api.GetShort)
	http.HandleFunc("/api/all", api.GetAll)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
