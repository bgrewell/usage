// Package pkg provides factory functions for creating formatters and groups.
package pkg

import (
	"github.com/bgrewell/usage/internal"
	"io"
)

// NewStandardFormatter creates a plain text formatter that outputs usage
// information without color codes. This is suitable for non-terminal outputs
// or environments that don't support ANSI color codes.
//
// Parameters:
//   - output: the writer for normal usage output (typically os.Stdout)
//   - error: the writer for error messages (typically os.Stderr)
//   - config: the configuration containing application and option information
//
// Returns a Formatter that can be used with WithFormatter option.
func NewStandardFormatter(output, error io.Writer, config *internal.Configuration) internal.Formatter {
	return &internal.StandardFormatter{
		Output:        output,
		Error:         error,
		Configuration: config,
	}
}

// NewColorFormatter creates a colored formatter that uses ANSI color codes
// to enhance the visual appearance of usage output. This is the default
// formatter used by NewUsage.
//
// Parameters:
//   - output: the writer for normal usage output (typically os.Stdout)
//   - error: the writer for error messages (typically os.Stderr)
//   - config: the configuration containing application and option information
//
// Returns a Formatter that can be used with WithFormatter option.
func NewColorFormatter(output, error io.Writer, config *internal.Configuration) internal.Formatter {
	return &internal.ColorFormatter{
		Output:        output,
		Error:         error,
		Configuration: config,
	}
}
