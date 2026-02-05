// Package internal provides internal types and implementations for the usage library.
// These types are not intended for direct use by library consumers.
package internal

// Formatter defines the interface for formatting and displaying usage information.
// Implementations can provide different output styles (plain text, colored, etc.).
type Formatter interface {
	// PrintUsage outputs the usage information including application details,
	// options, and arguments to the configured output writer.
	PrintUsage()

	// PrintError outputs an error message along with usage information
	// to the configured error writer.
	PrintError(err error)
}
