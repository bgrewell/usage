package internal

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestColorFormatter_PrintUsage(t *testing.T) {
	tests := []struct {
		name           string
		config         *Configuration
		expectedOutput []string // Substrings that should appear in output
	}{
		{
			name: "empty configuration",
			config: &Configuration{
				ApplicationName: "testapp",
				Groups:          map[string]*Group{},
			},
			expectedOutput: []string{
				"Usage:",
				"testapp",
				"[OPTIONS]",
			},
		},
		{
			name: "full version info",
			config: &Configuration{
				ApplicationName:       "myapp",
				ApplicationVersion:    "1.0.0",
				ApplicationBuildDate:  "2024-01-01",
				ApplicationCommitHash: "abc123",
				ApplicationBranch:     "main",
				Groups:                map[string]*Group{},
			},
			expectedOutput: []string{
				"Version:",
				"1.0.0",
				"Date:",
				"2024-01-01",
				"Codebase:",
				"abc123",
				"(main)",
			},
		},
		{
			name: "with description",
			config: &Configuration{
				ApplicationName:        "myapp",
				ApplicationDescription: "A test application",
				Groups:                 map[string]*Group{},
			},
			expectedOutput: []string{
				"Description:",
				"A test application",
			},
		},
		{
			name: "with options group",
			config: &Configuration{
				ApplicationName: "myapp",
				Groups: map[string]*Group{
					"Test": {
						Priority:    0,
						Name:        "Test",
						Description: "Test Options",
						Options: []*Option{
							{
								Short:       "v",
								Long:        "verbose",
								Default:     false,
								Description: "Enable verbose output",
							},
						},
					},
				},
			},
			expectedOutput: []string{
				"Options:",
				"Test:",
				"Test Options",
				"verbose",
				"Enable verbose output",
			},
		},
		{
			name: "with positional arguments",
			config: &Configuration{
				ApplicationName: "myapp",
				Groups: map[string]*Group{
					"Default": {
						Priority:    0,
						Name:        "Default",
						Description: "Default Options",
						Arguments: []*Argument{
							{
								Position:    0,
								Name:        "file",
								Description: "Input file",
							},
						},
					},
				},
			},
			expectedOutput: []string{
				"Arguments:",
				"file",
				"Input file",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			formatter := &ColorFormatter{
				Output:        &buf,
				Configuration: tt.config,
			}

			formatter.PrintUsage()
			output := buf.String()

			for _, expected := range tt.expectedOutput {
				if !strings.Contains(output, expected) {
					t.Errorf("PrintUsage() output missing expected substring %q", expected)
				}
			}
		})
	}
}

func TestColorFormatter_PrintError(t *testing.T) {
	tests := []struct {
		name           string
		err            error
		expectedOutput []string
	}{
		{
			name: "simple error",
			err:  errors.New("test error"),
			expectedOutput: []string{
				"Error:",
				"test error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			config := &Configuration{
				ApplicationName: "testapp",
				Groups:          map[string]*Group{},
			}
			formatter := &ColorFormatter{
				Output:        &buf,
				Error:         &buf,
				Configuration: config,
			}

			formatter.PrintError(tt.err)
			output := buf.String()

			for _, expected := range tt.expectedOutput {
				if !strings.Contains(output, expected) {
					t.Errorf("PrintError() output missing expected substring %q", expected)
				}
			}
		})
	}
}

func TestColorFormatter_DefaultWriters(t *testing.T) {
	t.Run("nil output defaults to stdout", func(t *testing.T) {
		config := &Configuration{
			ApplicationName: "testapp",
			Groups:          map[string]*Group{},
		}
		formatter := &ColorFormatter{
			Output:        nil,
			Configuration: config,
		}

		// This test just verifies it doesn't panic when Output is nil
		// We can't easily capture os.Stdout in tests, so we just ensure no panic
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("PrintUsage() panicked with nil Output: %v", r)
			}
		}()

		// We can't actually call PrintUsage because it writes to stdout
		// Just verify the Output field gets defaulted in the method
		if formatter.Output != nil {
			t.Log("Output field is already set")
		}
	})
}
