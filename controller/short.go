package controller

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

var SHORTLINK map[string]string

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
		sl := shortLink(8)
		fmt.Println(url)

		SHORTLINK := make(map[string]string)
		SHORTLINK[url] = sl

		// template.HTMLEscape(w, []byte(r.FormValue("url")))
		// template.HTMLEscape(w, []byte(url))
		//fmt.Fprintln(w, SHORTLINK[url])
		// fmt.Fprintf(w, url)
		// fmt.Println("test")
		var cookie http.Cookie

		cookie.Name = "shortLink"
		cookie.Value = SHORTLINK[url]

		// w.Header().Add("Set-Cookie", cookie.String())
		fmt.Println("right token")
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		fmt.Println("wrong token")
		http.Redirect(w, r, "/", http.StatusFound)
	}

}

func shortLink(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789"

	sid := rand.New(rand.NewSource(time.Now().UnixNano()))

	link := make([]byte, length)

	for k := range link {
		link[k] = charset[sid.Intn(len(charset))]
	}

	return string(link)

}
