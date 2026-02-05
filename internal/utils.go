package internal

import (
	"os"
	"path/filepath"
	"strings"
)

// GetExecutableName returns the name of the current executable without extension.
// It extracts the basename from the executable path and removes any file extension.
// This is used as the default application name if not explicitly provided.
func GetExecutableName() string {
	// Executable returns the path to the current executable.
	exePath, err := os.Executable()
	if err != nil {
		// Fallback to os.Args[0] if available
		if len(os.Args) > 0 {
			exePath = os.Args[0]
		} else {
			return "unknown"
		}
	}

	// Base returns the last element of the path.
	exeName := filepath.Base(exePath)

	// TrimSuffix removes the extension from the executable Name.
	// Adjust the cut set based on expected executable extensions.
	return strings.TrimSuffix(exeName, filepath.Ext(exeName))
}

// wrapText wraps the given text to the specified width and returns a slice of lines.
// It splits text on word boundaries to fit within the specified width, accounting
// for a prefix length on the first line. Subsequent lines use the full width.
// This is used internally by formatters to wrap long descriptions.
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
