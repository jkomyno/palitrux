package app

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/jkomyno/palitrux/config"
	"github.com/jkomyno/palitrux/router"
	"github.com/rs/cors"
)

// Server launches the microservice with the proper configuration
func Server(c *config.Config) (*http.Server, error) {
	handler := router.NewServerMux(c)

	if c.CorsEnabled {
		corsHandler := cors.New(cors.Options{
			AllowedOrigins: c.CorsAllowedOrigins,
		})

		handler = corsHandler.Handler(handler)
	}

	addr := ":" + strconv.Itoa(c.ServerPort)

	server := &http.Server{
		Addr:           addr,
		Handler:        handler,
		ReadTimeout:    time.Duration(c.HTTPReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(c.HTTPWriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("Listening on *%s\n", addr)
	return server, listenAndServe(server, c)
}

func listenAndServe(s *http.Server, c *config.Config) error {
	/*
		if c.CertFile != "" && c.KeyFile != "" {
			return s.ListenAndServeTLS(c.CertFile, c.KeyFile)
		}
	*/
	return s.ListenAndServe()
}
