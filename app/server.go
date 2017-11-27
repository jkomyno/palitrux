package app

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/jkomyno/palitrux/config"
	"github.com/jkomyno/palitrux/router"
)

// Server launches the microservice with the proper configuration
func Server(c *config.Config) error {
	handler := router.NewServerMux(c)
	addr := ":" + strconv.Itoa(c.ServerPort)

	server := &http.Server{
		Addr:           addr,
		Handler:        handler,
		ReadTimeout:    time.Duration(c.HTTPReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(c.HTTPWriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("Listening on *%s\n", addr)
	return listenAndServe(server, c)
}

func listenAndServe(s *http.Server, c *config.Config) error {
	/*
		if c.CertFile != "" && c.KeyFile != "" {
			return s.ListenAndServeTLS(c.CertFile, c.KeyFile)
		}
	*/
	return s.ListenAndServe()
}
