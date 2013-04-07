package main

import (
	"fmt"
	"github.com/gosexy/cli"
)

func main() {
	var err error

	// Dispatches the command.
	err = cli.Dispatch()

	if err != nil {
		fmt.Fatalf("Error: %s\n", err.Error())
	}

}
