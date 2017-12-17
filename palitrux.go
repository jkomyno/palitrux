package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/jkomyno/palitrux/app"
	"github.com/jkomyno/palitrux/config"
)

func exitWithError(args ...interface{}) {
	log.Fatal("%s\n", args)
	os.Exit(1)
}

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		exitWithError(err)
	}

	// starts the server
	server, err := app.Server(c)

	if err != nil {
		exitWithError("Couldn't start the server", err)
	}

	// Handles graceful shutdown
	interrupt := make(chan os.Signal, 1)
	// relays incoming signals to channel interrupt
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	killSignal := <-interrupt

	switch killSignal {
	case os.Interrupt:
		log.Print("Received SIGINT signal")
	case syscall.SIGTERM:
		log.Print("Received SIGTERM signal")
	}

	log.Print("The microservice is shutting down")
	err = server.Shutdown(context.Background())
	if err != nil {
		exitWithError(err)
	}
	log.Print("The microservice has shut down")
}
