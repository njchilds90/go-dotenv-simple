package dotenv

import (
	"os"
	"testing"
	"time"
)

func TestGetString(t *testing.T) {
	os.Setenv("TEST_STR", "hello")
	if got := GetString("TEST_STR", ""); got != "hello" {
		t.Fatalf("want hello, got %s", got)
	}
	if got := GetString("TEST_MISSING", "default"); got != "default" {
		t.Fatalf("want default, got %s", got)
	}
}

func TestGetInt(t *testing.T) {
	os.Setenv("TEST_INT", "42")
	if got := GetInt("TEST_INT", 0); got != 42 {
		t.Fatalf("want 42, got %d", got)
	}
	if got := GetInt("TEST_BAD_INT", 7); got != 7 {
		t.Fatalf("want fallback 7, got %d", got)
	}
}

func TestGetBool(t *testing.T) {
	os.Setenv("TEST_BOOL", "true")
	if !GetBool("TEST_BOOL", false) {
		t.Fatal("want true")
	}
}

func TestGetDuration(t *testing.T) {
	os.Setenv("TEST_DUR", "5s")
	if got := GetDuration("TEST_DUR", 0); got != 5*time.Second {
		t.Fatalf("want 5s, got %v", got)
	}
}

func TestParse(t *testing.T) {
	content := `
# comment
APP=myapp
PORT=9000
DEBUG=true
URL=https://example.com
CALLBACK=${URL}/cb
`
	f, _ := os.CreateTemp("", "dotenv*.env")
	f.WriteString(content)
	f.Close()
	defer os.Remove(f.Name())

	if err := Load(f.Name()); err != nil {
		t.Fatal(err)
	}
	if os.Getenv("APP") != "myapp" {
		t.Fatal("APP not loaded")
	}
}

func TestRequire(t *testing.T) {
	os.Setenv("REQ_KEY", "present")
	if err := Require("REQ_KEY"); err != nil {
		t.Fatal(err)
	}
	if err := Require("REQ_KEY", "REQ_MISSING"); err == nil {
		t.Fatal("expected error for missing key")
	}
}

func TestUnquote(t *testing.T) {
	cases := []struct{ in, want string }{
		{`"hello"`, "hello"},
		{`'world'`, "world"},
		{`plain`, "plain"},
	}
	for _, c := range cases {
		if got := unquote(c.in); got != c.want {
			t.Fatalf("unquote(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
