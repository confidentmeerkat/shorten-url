package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"strings"
	"urlshort/controller"
)

func main() {
	http.HandleFunc("/", controller.Handler)
	http.HandleFunc("/short", controller.Short)

	// http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/public/", serveResource)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func serveResource(w http.ResponseWriter, req *http.Request) {
	// fmt.Println(req.URL.Path)
	// fmt.Println(req.URL.RawPath)

	// path := "public" + req.URL.Path
	path := req.URL.Path
	path, _ = strings.CutPrefix(path, "/")

	var contentType string
	if strings.HasSuffix(path, ".css") {
		// path = "public/css/"
		contentType = "text/css"
		// fmt.Println(contentType)
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
		// fmt.Println(contentType)
	} else if strings.HasSuffix(path, ".jpg") {
		contentType = "image/jpg"
		// fmt.Println(contentType)
	} else if strings.HasSuffix(path, ".js") {
		contentType = "text/javascript"
		// fmt.Println(contentType)
	} else {
		contentType = "text/plain"
		// fmt.Println(contentType)
	}
	// fmt.Println(path)
	// o, _ := os.Getwd()
	// fmt.Println(o, path)
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
