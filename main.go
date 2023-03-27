package main

import (
	"log"
	"net/http"
	"urlshort/handlers"
)

func main() {
	handlers.Handle()

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
