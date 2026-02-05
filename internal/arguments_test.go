package internal

import "testing"

func TestArgument_Creation(t *testing.T) {
	tests := []struct {
		name        string
		position    int
		argName     string
		description string
		extra       string
	}{
		{
			name:        "basic argument",
			position:    0,
			argName:     "file",
			description: "Input file",
			extra:       "",
		},
		{
			name:        "argument with extra",
			position:    1,
			argName:     "output",
			description: "Output file",
			extra:       "optional",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arg := Argument{
				Position:    tt.position,
				Name:        tt.argName,
				Description: tt.description,
				Extra:       tt.extra,
			}

			if arg.Position != tt.position {
				t.Errorf("Position = %d, want %d", arg.Position, tt.position)
			}
			if arg.Name != tt.argName {
				t.Errorf("Name = %q, want %q", arg.Name, tt.argName)
			}
			if arg.Description != tt.description {
				t.Errorf("Description = %q, want %q", arg.Description, tt.description)
			}
			if arg.Extra != tt.extra {
				t.Errorf("Extra = %q, want %q", arg.Extra, tt.extra)
			}
		})
	}
}
