package controller

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"time"
)

type link struct {
	Token     string
	ShortLink string
	Status    string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t, err := template.ParseFiles("view/index.html")
		if err != nil {
			fmt.Printf("can't parse file")

			return
		}

		token := csrfToken()

		http.SetCookie(w, &http.Cookie{
			Name:  "csrfToken",
			Value: token,
		})

		short := ""
		shortLink, _ := r.Cookie("shortLink")

		if shortLink != nil {
			short = shortLink.Value

			http.SetCookie(w, &http.Cookie{
				Name:   "shortLink",
				MaxAge: -1,
			})

		}

		status := ""
		statusCookie, _ := r.Cookie("status")

		if statusCookie != nil {
			status = statusCookie.Value

			http.SetCookie(w, &http.Cookie{
				Name:   "status",
				MaxAge: -1,
			})

		}

		l := link{Token: token, ShortLink: short, Status: status}
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
