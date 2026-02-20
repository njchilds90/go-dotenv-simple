package dotenv

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Load reads the .env file and sets variables
// without overwriting existing environment variables.
// Returns nil if file doesn't exist.
func Load(filename string) error {
	return loadFile(filename, false)
}

// Overload reads the .env file and sets variables,
// overwriting existing environment variables.
// Returns nil if file doesn't exist.
func Overload(filename string) error {
	return loadFile(filename, true)
}

// Read parses the .env file and returns variables
// without modifying the environment.
// Returns empty map and nil if file doesn't exist.
func Read(filename string) (map[string]string, error) {
	return parseFile(filename)
}

func loadFile(filename string, overwrite bool) error {
	vars, err := parseFile(filename)
	if err != nil {
		return err
	}

	for key, value := range vars {
		if !overwrite {
			if _, exists := os.LookupEnv(key); exists {
				continue
			}
		}
		err := os.Setenv(key, value)
		if err != nil {
			return fmt.Errorf("dotenv: failed setting %s: %w", key, err)
		}
	}

	return nil
}

func parseFile(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		return make(map[string]string), nil // No error if missing
	}
	if err != nil {
		return nil, fmt.Errorf("dotenv: failed to open %s: %w", filename, err)
	}
	defer file.Close()

	vars := make(map[string]string)
	scanner := bufio.NewScanner(file)
	lineNum := 0
	var multiLineValue strings.Builder
	var multiLineKey string
	var inMultiLine bool

	for scanner.Scan() {
		lineNum++
		line := scanner.Text() // Don't trim yet for multi-line

		if inMultiLine {
			if strings.Contains(line, "\"") { // End of multi-line
				multiLineValue.WriteString("\n" + strings.TrimSpace(line[:strings.Index(line, "\"")]))
				vars[multiLineKey] = multiLineValue.String()
				inMultiLine = false
				continue
			}
			multiLineValue.WriteString("\n" + line)
			continue
		}

		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if strings.HasPrefix(line, "export ") {
			line = strings.TrimSpace(strings.TrimPrefix(line, "export "))
		}

		// Remove inline comments
		if idx := strings.Index(line, " #"); idx != -1 {
			line = strings.TrimSpace(line[:idx])
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("dotenv parse error in %s at line %d: invalid key=value format", filename, lineNum)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		if key == "" {
			return nil, fmt.Errorf("dotenv parse error in %s at line %d: empty key", filename, lineNum)
		}

		// Handle quotes and multi-line start
		if len(value) >= 2 {
			quoteChar := value[0]
			if (quoteChar == '"' || quoteChar == '\'') && value[len(value)-1] == quoteChar {
				value = value[1 : len(value)-1]
			} else if quoteChar == '"' || quoteChar == '\'' {
				// Start of multi-line (no closing quote)
				inMultiLine = true
				multiLineKey = key
				multiLineValue.Reset()
				multiLineValue.WriteString(value[1:])
				continue
			}
		}

		vars[key] = value
	}

	if inMultiLine {
		return nil, fmt.Errorf("dotenv parse error in %s at line %d: unclosed multi-line value", filename, lineNum)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("dotenv: error reading %s: %w", filename, err)
	}

	return vars, nil
}
