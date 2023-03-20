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
	// w.Header().Add("Content-Type", "text/html")
	count := 0
	if r.Method == "GET" {
		// if tok, err := r.Cookie("token"); err == nil {
		// 	// tok.Value = ""
		// 	// http.SetCookie(w, tok)
		// 	tok.MaxAge = -1
		// 	tok.Name = "token"
		// 	tok.Value = ""
		// 	http.SetCookie(w, tok)
		// 	fmt.Println(count+1, "co")
		// }
		count++
		t, err := indexHandler()
		if err != nil {
			log.Fatal(err)
		}

		token := csrfToken()

		// var cookie http.Cookie

		// cookie.Name = "token"
		// cookie.Value = token
		// cookie.MaxAge = 0

		// w.Header().Add("Set-Cookie", cookie.String())
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
			// http.SetCookie(w, shortLink)
			w.Header().Add("Set-Cookie", shortLink.String())
		}
		// fmt.Println(os.Getwd())
		// sLink, _ = os.Getwd()
		l := link{Token: token, ShortLink: sLink}
		t.Execute(w, l)
		// t.ExecuteTemplate()
		// shortLink.MaxAge = -1
	}
}

func csrfToken() string {
	h := md5.New()
	crutime := time.Now().Unix()

	io.WriteString(h, strconv.FormatInt(crutime, 10))
	io.WriteString(h, "MySup9erSecureSalt*/45+`~jhsFh")

	return fmt.Sprintf("%x", h.Sum(nil))
}
