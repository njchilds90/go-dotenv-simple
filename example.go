// example.go demonstrates usage of the dotenv package.

package main

import (
	"fmt"
	"os"

	"github.com/njchilds90/go-dotenv-simple/dotenv"
)

func main() {
	err := dotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env:", err)
		os.Exit(1)
	}

	value := os.Getenv("EXAMPLE_VAR")
	fmt.Println("EXAMPLE_VAR:", value)
}