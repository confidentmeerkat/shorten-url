package controller

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

type link struct {
	Token     string
	ShortLink string
}

func indexHandler() (*template.Template, error) {
	t, err := template.ParseFiles("view/index.html")
	if err != nil {
		return nil, fmt.Errorf("parsing view file: %v", err)
	}

	return t, nil
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t, err := indexHandler()
		if err != nil {
			log.Fatal(err)
		}

		token := csrfToken()

		http.SetCookie(w, &http.Cookie{
			Name:  "token",
			Value: token,
		})

		sLink := ""
		shortLink, err := r.Cookie("shortLink")
		if err == nil {
			sLink = shortLink.Value
			shortLink.MaxAge = 0
			shortLink.Value = ""
			http.SetCookie(w, shortLink)
		}
		l := link{Token: token, ShortLink: sLink}
		t.Execute(w, l)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func csrfToken() string {
	h := md5.New()
	crutime := time.Now().Unix()

	io.WriteString(h, strconv.FormatInt(crutime, 10))
	io.WriteString(h, "MySup9erSecureSalt*/45+`~jhsFh")

	return fmt.Sprintf("%x", h.Sum(nil))
}
