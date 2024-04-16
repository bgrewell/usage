package pkg

import "io"

type Formatter interface {
	Usage()
	SetOutput(output io.Writer)
}
