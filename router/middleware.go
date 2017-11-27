package router

import (
	"fmt"
	"net/http"

	"github.com/jkomyno/palitrux/config"
)

type httpHandler func(http.ResponseWriter, *http.Request)

func Middleware(fun httpHandler, c *config.Config) http.Handler {
	next := http.Handler(http.HandlerFunc(fun))

	return validate(defaultHeaders(next), c)
}

func defaultHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", fmt.Sprintf("Palitrux %s", config.Version))
		next.ServeHTTP(w, r)
	})
}
