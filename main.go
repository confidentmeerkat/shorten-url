package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"strings"
	"urlshort/handlers"
	"urlshort/handlers/api"
)

func main() {
	http.HandleFunc("/", handlers.Handler)
	http.HandleFunc("/short", handlers.Short)

	http.HandleFunc("/public/", serveResource)

	http.HandleFunc("/api/get", api.GetShort)
	http.HandleFunc("/api/all", api.GetAll)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func serveResource(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	path, _ = strings.CutPrefix(path, "/")

	var contentType string
	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	} else if strings.HasSuffix(path, ".jpg") {
		contentType = "image/jpg"
	} else if strings.HasSuffix(path, ".svg") {
		contentType = "image/svg"
	} else if strings.HasSuffix(path, ".js") {
		contentType = "text/javascript"
	} else {
		contentType = "text/plain"
	}

	f, err := os.Open(path)

	if err == nil {
		defer f.Close()
		w.Header().Add("Content-Type", contentType)

		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}
