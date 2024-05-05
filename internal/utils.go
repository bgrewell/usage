package internal

import (
	"os"
	"path/filepath"
	"strings"
)

func GetExecutableName() string {
	// Executable returns the path to the current executable.
	exePath, err := os.Executable()
	if err != nil {
		panic(err) // Handle error according to your needs.
	}

	// Base returns the last element of the path.
	exeName := filepath.Base(exePath)

	// TrimSuffix removes the extension from the executable Name.
	// Adjust the cut set based on expected executable extensions.
	return strings.TrimSuffix(exeName, filepath.Ext(exeName))
}

// wrapText wraps the given text to the specified width and returns a slice of lines.
// It automatically indents the second and subsequent lines by the specified number of spaces.
func wrapText(text string, width int, prefixLen int) []string {
	if text == "" {
		return []string{""}
	}

	var lines []string
	words := strings.Fields(text)
	line := ""
	lineLength := width - prefixLen
	for _, word := range words {
		if len(line)+len(word)+1 > lineLength {
			lines = append(lines, line)
			line = ""
			if len(lines) > 0 {
				lineLength = width
			}
		}
		if line == "" {
			line = word
		} else {
			line += " " + word
		}
	}
	if line != "" {
		lines = append(lines, line)
	}
	return lines
}
