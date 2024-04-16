package pkg

import (
	"io"
	"os"
)

type StandardFormatter struct {
	output io.Writer
	error  io.Writer
}

func (f *StandardFormatter) Usage() {
	if f.output == nil {
		f.output = os.Stdout
	}
}

func (f *StandardFormatter) SetOutput(output io.Writer) {
	f.output = output
}
