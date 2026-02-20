# go-dotenv-simple

A tiny, dependency-free `.env` loader for Go.

Inspired by Python dotenv. Designed to be:
- Simple for humans
- Predictable for AI agents
- Zero dependencies
- Easy to audit

GitHub: https://github.com/njchilds90/go-dotenv-simple

---

## Install

```bash
go get github.com/njchilds90/go-dotenv-simple/dotenv
```

---

## Quick Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/njchilds90/go-dotenv-simple/dotenv"
)

func main() {
	err := dotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB_HOST:", getenv("DB_HOST"))
}
```

---

## Supported Syntax

Supports common dotenv features:

```
# Comments
FOO=bar

# Quoted values
NAME="John Doe"
PASSWORD='abc123'

# Inline comments
PORT=8080 # dev port

# export prefix
export API_KEY=xyz
```

---

## Behavior

### Load(filename)
- Loads variables from file
- Does NOT overwrite existing environment variables

### Overload(filename)
- Loads variables
- DOES overwrite existing environment variables

### Read(filename)
- Parses file and returns map[string]string
- Does NOT modify environment

---

## Error Handling

Errors include:
- File name
- Line number
- Reason for failure

Example:

```
dotenv parse error in .env at line 4: invalid key=value format
```

---

## When To Use

✅ Local development  
✅ CLI tools  
✅ Small services  
✅ AI agent tooling  

⚠️ For production, prefer real environment injection via:
- Docker
- Kubernetes
- CI/CD secrets
- OS environment

---

## Philosophy

- No magic
- No global state beyond os.Setenv
- No reflection
- No third-party dependencies
- Easy to read source
- AI agent friendly

---

## Testing

Run:

```
go test ./...
```

---

## License

MIT

---

Maintained by @njchilds90
