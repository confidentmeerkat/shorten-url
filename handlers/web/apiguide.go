package web

import (
	"html/template"
	"log"
	"net/http"
)

func APIGuide(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method is not allowed"))

		return
	}

	t, err := template.ParseFiles("web/apiguide.html")
	if err != nil {
		log.Println("can't parse file: apiguide.html")

		return
	}

	t.Execute(w, nil)
}
