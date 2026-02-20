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
export PORT=8080 # inline comment
PASSWORD='secret'
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

	if len(vars) != 4 {
		t.Errorf("Expected 4 vars, got %d", len(vars))
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
	if vars["PASSWORD"] != "secret" {
		t.Error("Expected PASSWORD=secret")
	}
}

func TestLoad(t *testing.T) {
	os.WriteFile(".env.test2", []byte("A=1\nB=2"), 0644)
	defer os.Remove(".env.test2")

	err := Load(".env.test2")
	if err != nil {
		t.Fatal(err)
	}

	if os.Getenv("A") != "1" {
		t.Error("Expected A=1")
	}
	if os.Getenv("B") != "2" {
		t.Error("Expected B=2")
	}
}

func TestOverload(t *testing.T) {
	os.Setenv("C", "old")
	os.WriteFile(".env.test3", []byte("C=new"), 0644)
	defer os.Remove(".env.test3")

	err := Overload(".env.test3")
	if err != nil {
		t.Fatal(err)
	}

	if os.Getenv("C") != "new" {
		t.Error("Expected C=new after overload")
	}
}

func TestMissingFile(t *testing.T) {
	vars, err := Read("nonexistent.env")
	if err != nil {
		t.Error("Expected no error for missing file")
	}
	if len(vars) != 0 {
		t.Error("Expected empty map for missing file")
	}

	err = Load("nonexistent.env")
	if err != nil {
		t.Error("Expected no error for Load on missing file")
	}
}

func TestParseErrors(t *testing.T) {
	content := "INVALID"
	os.WriteFile(".env.test4", []byte(content), 0644)
	defer os.Remove(".env.test4")

	_, err := Read(".env.test4")
	if err == nil {
		t.Error("Expected parse error")
	}
}

func TestMultiLine(t *testing.T) {
	content := `MULTI="line1
line2
line3"`
	os.WriteFile(".env.test5", []byte(content), 0644)
	defer os.Remove(".env.test5")

	vars, err := Read(".env.test5")
	if err != nil {
		t.Fatal(err)
	}
	expected := "line1\nline2\nline3"
	if vars["MULTI"] != expected {
		t.Errorf("Expected %q, got %q", expected, vars["MULTI"])
	}
}
