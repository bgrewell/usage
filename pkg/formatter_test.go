package pkg

import (
	"bytes"
	"github.com/bgrewell/usage/internal"
	"testing"
)

func TestNewStandardFormatter(t *testing.T) {
	var output, error bytes.Buffer
	config := &internal.Configuration{}

	formatter := NewStandardFormatter(&output, &error, config)

	if _, ok := formatter.(*internal.StandardFormatter); !ok {
		t.Errorf("NewStandardFormatter() = %T; want *internal.StandardFormatter", formatter)
	}
}

func TestNewColorFormatter(t *testing.T) {
	var output, error bytes.Buffer
	config := &internal.Configuration{}

	formatter := NewColorFormatter(&output, &error, config)

	if _, ok := formatter.(*internal.ColorFormatter); !ok {
		t.Errorf("NewColorFormatter() = %T; want *internal.ColorFormatter", formatter)
	}
}
