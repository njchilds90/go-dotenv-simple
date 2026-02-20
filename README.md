# go-dotenv-simple

A simple Go library to load environment variables from a `.env` file, inspired by Python's dotenv. Pure Go, no dependencies.

## Installation

```bash
go get github.com/njchilds90/go-dotenv-simple

Usage
Create a .env file:
text# Comment
EXAMPLE_VAR=hello world
QUOTED="value with spaces"
In your Go code:
Goimport (
	"fmt"
	"os"

	"github.com/njchilds90/go-dotenv-simple/dotenv"
)

func main() {
	err := dotenv.Load(".env")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(os.Getenv("EXAMPLE_VAR"))  // hello world
	fmt.Println(os.Getenv("QUOTED"))      // value with spaces
}
Options

dotenv.LoadNoOverwrite(".env"): Loads but doesn't overwrite existing env vars.

Why This Library?

Simple and lightweight.
Useful for config in development, scripts, or AI agents handling env-based secrets.
Easy to extend.

See example.go for more.
