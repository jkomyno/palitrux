package router

import (
	"mime"
	"net/http"
	"strings"

	"github.com/jkomyno/palitrux/config"
	"github.com/jkomyno/palitrux/errors"
)

func validate(next http.Handler, c *config.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" && r.Method != "POST" {
			errors.ReplyWithError(w, errors.ErrorMethodNotAllowed)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func validateImage(next http.Handler, c *config.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)
	})
}

func acceptMimeType(accept string) string {
	for _, v := range strings.Split(accept, ",") {
		mediatype, _, _ := mime.ParseMediaType(v)
		if mediatype == "image/png" {
			return "png"
		} else if mediatype == "image/jpeg" {
			return "jpeg"
		}
	}
	// default
	return ""
}
