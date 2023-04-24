// Package web serves web page requests.
package web

import (
	"html/template"
	"log"
	"net/http"
	"urlshort/configs"
	"urlshort/database/postgres"
	"urlshort/pkg"
)

// ShortHandler handles POST requests from the index web page,
// and creates a short link.
func ShortHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method is not allowed"))

		return
	}

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

	if cToken.Value != token {
		log.Println("wrong token")

		http.SetCookie(w, &http.Cookie{
			Name:  "status",
			Value: "CSRF token mismatch",
		})

		http.Redirect(w, r, "/", http.StatusFound)

		return
	}

	u := template.HTMLEscapeString(r.FormValue("url"))

	if !pkg.IsValidURL(u) {
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
	defer db.Close()

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

	short = configs.Domain + "/" + short

	http.SetCookie(w, &http.Cookie{
		Name:  "shortLink",
		Value: short,
	})

	http.Redirect(w, r, "/", http.StatusFound)
}
