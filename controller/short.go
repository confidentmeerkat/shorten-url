package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"urlshort/database/postgres"
)

func Short(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		token := r.FormValue("token")

		cToken, err := r.Cookie("csrfToken")
		if err != nil {
			http.Redirect(w, r, "/", http.StatusFound)

			return
		}

		fmt.Println("token: ", token)
		fmt.Println("cookie: ", cToken.Value)
		if cToken.Value == token {
			u := template.HTMLEscapeString(r.FormValue("url"))
			_, err := url.ParseRequestURI(u)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(u)

			db, err := postgres.New()
			if err != nil {
				w.WriteHeader(503)

				log.Fatal(err)
			}
			fmt.Println("connected")

			short, err := db.GetShort(u)
			if err != nil {
				short, err = db.CreateShort(u, 4)
				if err != nil {
					w.WriteHeader(503)

					log.Fatal(err)
				}
			}

			var cookie http.Cookie

			cookie.Name = "shortLink"
			cookie.Value = short

			fmt.Println("right token")
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			fmt.Println("wrong token")
			http.Redirect(w, r, "/", http.StatusFound)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
