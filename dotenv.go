// Package dotenv loads .env files into environment variables.
// It provides typed getters, validation, and multi-file support
// with zero external dependencies.
//
// Basic usage:
//
//	dotenv.Load()
//	port := dotenv.GetInt("PORT", 8080)
package dotenv

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Load reads ".env" from the current directory and sets any key that
// is not already present in the environment. It silently succeeds if
// the file does not exist.
func Load(files ...string) error {
	if len(files) == 0 {
		files = []string{".env"}
	}
	return load(false, files...)
}

// LoadFiles is an alias for Load with explicit file paths.
func LoadFiles(files ...string) error {
	if len(files) == 0 {
		files = []string{".env"}
	}
	return load(false, files...)
}

// Overload works like Load but overwrites existing environment variables.
func Overload(files ...string) error {
	if len(files) == 0 {
		files = []string{".env"}
	}
	return load(true, files...)
}

// Require returns an error if any of the named keys are absent or empty.
func Require(keys ...string) error {
	var missing []string
	for _, k := range keys {
		if strings.TrimSpace(os.Getenv(k)) == "" {
			missing = append(missing, k)
		}
	}
	if len(missing) > 0 {
		return fmt.Errorf("go-dotenv: missing required keys: %s", strings.Join(missing, ", "))
	}
	return nil
}

// GetString returns the value for key, or fallback if unset/empty.
func GetString(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// GetInt returns the integer value for key, or fallback on error/absence.
func GetInt(key string, fallback int) int {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	i, err := strconv.Atoi(strings.TrimSpace(v))
	if err != nil {
		return fallback
	}
	return i
}

// GetBool returns the boolean value for key, or fallback on error/absence.
// Accepted true values: 1, t, T, TRUE, true, True.
func GetBool(key string, fallback bool) bool {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	b, err := strconv.ParseBool(strings.TrimSpace(v))
	if err != nil {
		return fallback
	}
	return b
}

// GetFloat returns the float64 value for key, or fallback on error/absence.
func GetFloat(key string, fallback float64) float64 {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	f, err := strconv.ParseFloat(strings.TrimSpace(v), 64)
	if err != nil {
		return fallback
	}
	return f
}

// GetDuration returns the time.Duration value for key, or fallback on error/absence.
// Values are parsed with time.ParseDuration (e.g. "30s", "1m", "2h").
func GetDuration(key string, fallback time.Duration) time.Duration {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	d, err := time.ParseDuration(strings.TrimSpace(v))
	if err != nil {
		return fallback
	}
	return d
}

// ── internal ──────────────────────────────────────────────────────────────────

func load(overwrite bool, files ...string) error {
	for _, f := range files {
		if err := loadFile(f, overwrite); err != nil && !os.IsNotExist(err) {
			return err
		}
	}
	return nil
}

func loadFile(filename string, overwrite bool) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	vars, err := parse(f)
	if err != nil {
		return err
	}
	for k, v := range vars {
		if overwrite || os.Getenv(k) == "" {
			os.Setenv(k, expand(v, vars))
		}
	}
	return nil
}

func parse(f *os.File) (map[string]string, error) {
	out := make(map[string]string)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		idx := strings.IndexByte(line, '=')
		if idx < 0 {
			continue
		}
		key := strings.TrimSpace(line[:idx])
		val := strings.TrimSpace(line[idx+1:])
		val = stripInlineComment(val)
		val = unquote(val)
		out[key] = val
	}
	return out, scanner.Err()
}

func stripInlineComment(s string) string {
	if !strings.Contains(s, " #") {
		return s
	}
	// only strip if not inside quotes
	if len(s) > 0 && (s[0] == '"' || s[0] == '\'') {
		return s
	}
	idx := strings.Index(s, " #")
	return strings.TrimSpace(s[:idx])
}

func unquote(s string) string {
	if len(s) >= 2 {
		if (s[0] == '"' && s[len(s)-1] == '"') ||
			(si f(0) == '\'' && s[len(s)-1] == '\'' ) {
			return s[1 : len(s)-1]
		}
	}
	return s
}

func expand(value string, env map[string]string) string {
	return os.Expand(value, func(key string) string {
		if v, ok := env[key]; ok {
			return v
		}
		return os.Getenv(key)
	})
}