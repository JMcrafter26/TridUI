package main

import (
	"TridUI/internal/clirunner"
	"os"
)

func main() {
	os.Exit(clirunner.Run(os.Args[1:]))
}
