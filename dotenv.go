// Package dotenv provides a simple way to load environment variables from a .env file.
// Inspired by Python's dotenv library. It parses key-value pairs and sets them in the os environment.
//
// Usage:
//   err := dotenv.Load(".env")
//   if err != nil {
//     log.Fatal(err)
//   }
//   value := os.Getenv("MY_VAR")
package dotenv

import (
	"bufio"
	"os"
	"strings"
)

// Load reads the .env file and sets the environment variables.
// If the file does not exist, it returns nil (no error).
// It skips lines starting with '#' (comments) and empty lines.
// Values can be quoted with double quotes to include spaces.
func Load(filename string) error {
	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		return nil // No error if file missing
	}
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue // Skip empty or comments
		}

		// Split on first '='
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue // Invalid line
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Handle quoted values
		if strings.HasPrefix(value, "\"") && strings.HasSuffix(value, "\"") {
			value = value[1 : len(value)-1]
		}

		// Set env (overwrite if exists)
		os.Setenv(key, value)
	}

	return scanner.Err()
}

// LoadNoOverwrite is like Load but does not overwrite existing env vars.
func LoadNoOverwrite(filename string) error {
	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		if strings.HasPrefix(value, "\"") && strings.HasSuffix(value, "\"") {
			value = value[1 : len(value)-1]
		}

		// Set only if not already set
		if _, exists := os.LookupEnv(key); !exists {
			os.Setenv(key, value)
		}
	}

	return scanner.Err()
}
