package main

import (
	"log"
	"menteslibres.net/gosexy/cli"
)

func main() {
	var err error

	// Dispatches the command.
	err = cli.Dispatch()

	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

}
