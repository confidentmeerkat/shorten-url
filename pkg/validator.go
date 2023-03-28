package pkg

import (
	"net/url"
	"regexp"
)

func IsValidURL(s string) bool {
	_, err := url.ParseRequestURI(s)
	if err != nil {
		return false
	}

	u, err := url.Parse(s)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	if u.Scheme != "https" && u.Scheme != "http" {
		return false
	}

	r := regexp.MustCompile(`(^[a-z0-9]+\.{1}[a-z]+)`)

	return r.MatchString(u.Host)
}
