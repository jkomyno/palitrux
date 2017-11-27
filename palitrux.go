package main

import (
	"fmt"
	"os"

	"github.com/jkomyno/palitrux/app"
	"github.com/jkomyno/palitrux/config"
)

func exitWithError(args ...interface{}) {
	fmt.Fprintf(os.Stderr, "%s\n", args)
	os.Exit(1)
}

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		exitWithError(err)
	}

	// starts the server
	err = app.Server(c)

	if err != nil {
		exitWithError("Couldn't start the server", err)
	}
}
