package dotenv

import (
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	content := `
# Comment
FOO=bar
NAME="John Doe"
export PORT=8080
`

	err := os.WriteFile(".env.test", []byte(content), 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(".env.test")

	vars, err := Read(".env.test")
	if err != nil {
		t.Fatal(err)
	}

	if vars["FOO"] != "bar" {
		t.Error("Expected FOO=bar")
	}
	if vars["NAME"] != "John Doe" {
		t.Error("Expected NAME=John Doe")
	}
	if vars["PORT"] != "8080" {
		t.Error("Expected PORT=8080")
	}
}

func TestLoad(t *testing.T) {
	os.WriteFile(".env.test2", []byte("A=1"), 0644)
	defer os.Remove(".env.test2")

	err := Load(".env.test2")
	if err != nil {
		t.Fatal(err)
	}

	if os.Getenv("A") != "1" {
		t.Error("Expected A=1")
	}
}
