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

	// TrimSuffix removes the extension from the executable name.
	// Adjust the cut set based on expected executable extensions.
	return strings.TrimSuffix(exeName, filepath.Ext(exeName))
}
