# go-dotenv

[![Go Reference](https://pkg.go.dev/badge/github.com/njchilds90/go-dotenv.svg)](https://pkg.go.dev/github.com/njchilds90/go-dotenv)
[![Go Report Card](https://goreportcard.com/badge/github.com/njchilds90/go-dotenv)](https://goreportcard.com/report/github.com/njchilds90/go-dotenv)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

A zero-dependency `.env` file loader for Go with typed getters, validation, and multi-file support.

## Features

- Load one or many `.env` files
- Typed getters: `GetString`, `GetInt`, `GetBool`, `GetDuration`
- Default values and required-key validation
- Expand `${VAR}` references within values
- Does **not** overwrite existing environment variables (safe for production)
- Zero external dependencies

## Install
```bash
go get github.com/njchilds90/go-dotenv-simple
```

## Quick Start
```go
package main

import (
    "fmt"
    "github.com/njchilds90/go-dotenv-simple"
)

func main() {
    dotenv.Load() // loads .env by default

    port := dotenv.GetInt("PORT", 8080)
    debug := dotenv.GetBool("DEBUG", false)
    dsn := dotenv.GetString("DATABASE_URL", "")

    fmt.Println(port, debug, dsn)
}
```

## API

### Loading
```go
dotenv.Load()                          // load .env
dotenv.LoadFiles(".env", ".env.local") // load multiple files
dotenv.Overload(".env")                // overwrite existing env vars
```

### Typed Getters
```go
dotenv.GetString("KEY", "default")
dotenv.GetInt("PORT", 8080)
dotenv.GetBool("DEBUG", false)
dotenv.GetDuration("TIMEOUT", 30*time.Second)
```

### Validation
```go
err := dotenv.Require("DATABASE_URL", "SECRET_KEY")
// returns error if any key is missing or empty
```

## .env File Format
```env
# Comments supported
APP_NAME=my-app
PORT=8080
DEBUG=true
TIMEOUT=30s
BASE_URL=https://example.com
CALLBACK=${BASE_URL}/callback
```

## Contributing

Pull requests welcome. Please run `go test ./...` before submitting.

---

## License

MIT

---

Maintained by @njchilds90