package controller

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"urlshort/database/postgres"

	"github.com/lib/pq"
)

// var SHORTLINK map[string]string

func Short(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	cToken, err := r.Cookie("token")
	if err != nil {
		w.WriteHeader(http.StatusFound)
	}

	// if r.Method == "POST" {
	// r.ParseForm()
	// token := r.FormValue("token")
	fmt.Println("token: ", token)
	fmt.Println("cookie: ", cToken.Value)
	if cToken.Value == token {
		url := template.HTMLEscapeString(r.FormValue("url"))
		// sl := shortLink(8)
		fmt.Println(url)

		// SHORTLINK := make(map[string]string)
		// SHORTLINK[url] = sl

		db, err := postgres.New()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("connected")

		short, err := db.CreateShort(url, 4)
		if err != nil {
			var e *pq.Error

			if errors.As(err, &e) {
				switch e.Code {
				case "23505":
					fmt.Println("WOW duplicate!")
				}
			}

			log.Fatal(err)
		}

		// template.HTMLEscape(w, []byte(r.FormValue("url")))
		// template.HTMLEscape(w, []byte(url))
		//fmt.Fprintln(w, SHORTLINK[url])
		// fmt.Fprintf(w, url)
		// fmt.Println("test")
		var cookie http.Cookie

		cookie.Name = "shortLink"
		cookie.Value = short
		// w.Header().Add("Set-Cookie", cookie.String())
		fmt.Println("right token")
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		fmt.Println("wrong token")
		http.Redirect(w, r, "/", http.StatusFound)
	}

}
