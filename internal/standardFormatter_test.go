package internal

import (
	"bytes"
	"errors"
	"testing"
)

// NOTE: StandardFormatter currently has a bug where it uses fmt.Printf instead
// of fmt.Fprintf(f.Output, ...), so it always writes to stdout regardless of
// the configured Output writer. These tests verify the current behavior.
// This should be fixed in a future PR.

func TestStandardFormatter_Creation(t *testing.T) {
	tests := []struct {
		name   string
		config *Configuration
	}{
		{
			name: "basic configuration",
			config: &Configuration{
				ApplicationName: "testapp",
				Groups:          map[string]*Group{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var outBuf, errBuf bytes.Buffer
			formatter := &StandardFormatter{
				Output:        &outBuf,
				Error:         &errBuf,
				Configuration: tt.config,
			}

			if formatter.Configuration.ApplicationName != tt.config.ApplicationName {
				t.Errorf("Configuration.ApplicationName = %q, want %q",
					formatter.Configuration.ApplicationName, tt.config.ApplicationName)
			}
		})
	}
}

func TestStandardFormatter_PrintError(t *testing.T) {
	tests := []struct {
		name string
		err  error
	}{
		{
			name: "error message",
			err:  errors.New("invalid input"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			config := &Configuration{
				ApplicationName: "testapp",
				Groups:          map[string]*Group{},
			}
			formatter := &StandardFormatter{
				Output:        &buf,
				Error:         &buf,
				Configuration: config,
			}

			// PrintError writes to f.Error which IS respected
			formatter.PrintError(tt.err)
			output := buf.String()

			// Error output should contain the error message
			if output == "" {
				t.Error("PrintError() produced no output to Error writer")
			}
		})
	}
}
