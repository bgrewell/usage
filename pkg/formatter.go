package pkg

import (
	"github.com/bgrewell/usage/internal"
	"io"
)

func NewStandardFormatter(output, error io.Writer, config *internal.Configuration) internal.Formatter {
	return &internal.StandardFormatter{
		Output:        output,
		Error:         error,
		Configuration: config,
	}
}
