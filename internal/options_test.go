package internal

import "testing"

func TestOption_Creation(t *testing.T) {
	tests := []struct {
		name         string
		short        string
		long         string
		defaultValue interface{}
		description  string
		extra        string
	}{
		{
			name:         "boolean option",
			short:        "v",
			long:         "verbose",
			defaultValue: false,
			description:  "Enable verbose mode",
			extra:        "",
		},
		{
			name:         "integer option",
			short:        "p",
			long:         "port",
			defaultValue: 8080,
			description:  "Port number",
			extra:        "1-65535",
		},
		{
			name:         "string option",
			short:        "o",
			long:         "output",
			defaultValue: "output.txt",
			description:  "Output file",
			extra:        "required",
		},
		{
			name:         "float option",
			short:        "r",
			long:         "rate",
			defaultValue: 1.5,
			description:  "Sample rate",
			extra:        "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opt := Option{
				Short:       tt.short,
				Long:        tt.long,
				Default:     tt.defaultValue,
				Description: tt.description,
				Extra:       tt.extra,
			}

			if opt.Short != tt.short {
				t.Errorf("Short = %q, want %q", opt.Short, tt.short)
			}
			if opt.Long != tt.long {
				t.Errorf("Long = %q, want %q", opt.Long, tt.long)
			}
			if opt.Default != tt.defaultValue {
				t.Errorf("Default = %v, want %v", opt.Default, tt.defaultValue)
			}
			if opt.Description != tt.description {
				t.Errorf("Description = %q, want %q", opt.Description, tt.description)
			}
			if opt.Extra != tt.extra {
				t.Errorf("Extra = %q, want %q", opt.Extra, tt.extra)
			}
		})
	}
}

func TestOption_EmptyNames(t *testing.T) {
	t.Run("empty short name", func(t *testing.T) {
		opt := Option{
			Short:       "",
			Long:        "verbose",
			Default:     false,
			Description: "Verbose output",
		}

		if opt.Short != "" {
			t.Errorf("Short should be empty, got %q", opt.Short)
		}
		if opt.Long != "verbose" {
			t.Errorf("Long = %q, want %q", opt.Long, "verbose")
		}
	})

	t.Run("empty long name", func(t *testing.T) {
		opt := Option{
			Short:       "v",
			Long:        "",
			Default:     false,
			Description: "Verbose output",
		}

		if opt.Short != "v" {
			t.Errorf("Short = %q, want %q", opt.Short, "v")
		}
		if opt.Long != "" {
			t.Errorf("Long should be empty, got %q", opt.Long)
		}
	})
}
