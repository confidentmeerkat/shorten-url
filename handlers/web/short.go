package web

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"urlshort/database/postgres"
)

func Short(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		token := r.FormValue("token")

		cToken, err := r.Cookie("csrfToken")
		if err != nil {
			http.SetCookie(w, &http.Cookie{
				Name:  "status",
				Value: "no CSRF token",
			})

			http.Redirect(w, r, "/", http.StatusFound)

			return
		}

		if cToken.Value == token {
			u := template.HTMLEscapeString(r.FormValue("url"))
			_, err := url.ParseRequestURI(u)
			if err != nil {
				log.Println(err)

				http.SetCookie(w, &http.Cookie{
					Name:  "status",
					Value: "no valid URL",
				})

				http.Redirect(w, r, "/", http.StatusFound)

				return
			}

			db, err := postgres.New()
			if err != nil {
				log.Println(err)

				http.SetCookie(w, &http.Cookie{
					Name:  "status",
					Value: "service unavailable",
				})

				http.Redirect(w, r, "/", http.StatusFound)

				return
			}

			short, err := db.GetShort(u)
			if err != nil {
				short, err = db.CreateShort(u, 4)
				if err != nil {
					log.Println(err)

					http.SetCookie(w, &http.Cookie{
						Name:  "status",
						Value: "creating short link failed",
					})

					http.Redirect(w, r, "/", http.StatusFound)

					return
				}
			}

			short = os.Getenv("SHORTENER_DOMAIN") + "/" + short

			http.SetCookie(w, &http.Cookie{
				Name:  "shortLink",
				Value: short,
			})
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			log.Println("wrong token")

			http.SetCookie(w, &http.Cookie{
				Name:  "status",
				Value: "CSRF token mismatch",
			})

			http.Redirect(w, r, "/", http.StatusFound)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method is not allowed"))
	}

}
