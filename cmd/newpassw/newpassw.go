package main

import (
	"fmt"
	"os"

	"github.com/littleroot/passw"
)

func main() {
	pw, err := passw.Generate()
	if err != nil {
		fmt.Fprintf(os.Stderr, "generate password: %s\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "%s\n", pw)
}
