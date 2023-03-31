// Package web serves web page requests.
package web

import (
	"bufio"
	"net/http"
	"os"
	"strings"
)

// ServeResource handles resources like CSS, JS, and images.
func ServeResource(w http.ResponseWriter, req *http.Request) {
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
		contentType = "image/svg+xml"
	} else if strings.HasSuffix(path, ".js") {
		contentType = "text/javascript"
	} else {
		contentType = "text/plain"
	}

	f, err := os.Open(path)
	if err != nil {
		w.WriteHeader(404)

		return
	}
	defer f.Close()

	w.Header().Add("Content-Type", contentType)

	br := bufio.NewReader(f)
	br.WriteTo(w)
}
