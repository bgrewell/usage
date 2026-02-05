package internal

import (
	"testing"
)

func TestGroup_AddOption(t *testing.T) {
	tests := []struct {
		name          string
		initialCount  int
		optionsToAdd  int
		expectedCount int
	}{
		{
			name:          "add single option",
			initialCount:  0,
			optionsToAdd:  1,
			expectedCount: 1,
		},
		{
			name:          "add multiple options",
			initialCount:  0,
			optionsToAdd:  3,
			expectedCount: 3,
		},
		{
			name:          "add to existing options",
			initialCount:  2,
			optionsToAdd:  1,
			expectedCount: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			group := &Group{}

			// Add initial options
			for i := 0; i < tt.initialCount; i++ {
				group.AddOption(&Option{})
			}

			// Add test options
			for i := 0; i < tt.optionsToAdd; i++ {
				group.AddOption(&Option{})
			}

			if len(group.Options) != tt.expectedCount {
				t.Errorf("AddOption() resulted in %d options, want %d", len(group.Options), tt.expectedCount)
			}
		})
	}
}

func TestGroup_RemoveOption(t *testing.T) {
	tests := []struct {
		name            string
		setupOptions    []*Option
		removeOption    *Option
		expectedCount   int
		shouldBeRemoved bool
	}{
		{
			name: "remove existing option",
			setupOptions: []*Option{
				{Short: "a"},
				{Short: "b"},
				{Short: "c"},
			},
			removeOption:    &Option{Short: "b"},
			expectedCount:   2,
			shouldBeRemoved: true,
		},
		{
			name: "remove non-existing option",
			setupOptions: []*Option{
				{Short: "a"},
			},
			removeOption:    &Option{Short: "b"},
			expectedCount:   1,
			shouldBeRemoved: false,
		},
		{
			name:            "remove from empty group",
			setupOptions:    []*Option{},
			removeOption:    &Option{Short: "a"},
			expectedCount:   0,
			shouldBeRemoved: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			group := &Group{Options: tt.setupOptions}

			// For the "remove existing" test, use the actual option from the slice
			if tt.shouldBeRemoved && len(tt.setupOptions) > 0 {
				tt.removeOption = tt.setupOptions[1] // Remove the middle one
			}

			group.RemoveOption(tt.removeOption)

			if len(group.Options) != tt.expectedCount {
				t.Errorf("RemoveOption() resulted in %d options, want %d", len(group.Options), tt.expectedCount)
			}
		})
	}
}

func TestGroup_AddArgument(t *testing.T) {
	tests := []struct {
		name          string
		initialCount  int
		argsToAdd     int
		expectedCount int
	}{
		{
			name:          "add single argument",
			initialCount:  0,
			argsToAdd:     1,
			expectedCount: 1,
		},
		{
			name:          "add multiple arguments",
			initialCount:  0,
			argsToAdd:     3,
			expectedCount: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			group := &Group{}

			// Add initial arguments
			for i := 0; i < tt.initialCount; i++ {
				group.AddArgument(&Argument{})
			}

			// Add test arguments
			for i := 0; i < tt.argsToAdd; i++ {
				group.AddArgument(&Argument{})
			}

			if len(group.Arguments) != tt.expectedCount {
				t.Errorf("AddArgument() resulted in %d arguments, want %d", len(group.Arguments), tt.expectedCount)
			}
		})
	}
}

func TestGroup_RemoveArgument(t *testing.T) {
	tests := []struct {
		name            string
		setupArgs       []*Argument
		removeArg       *Argument
		expectedCount   int
		shouldBeRemoved bool
	}{
		{
			name: "remove existing argument",
			setupArgs: []*Argument{
				{Name: "file1"},
				{Name: "file2"},
			},
			removeArg:       &Argument{Name: "file2"},
			expectedCount:   1,
			shouldBeRemoved: true,
		},
		{
			name: "remove non-existing argument",
			setupArgs: []*Argument{
				{Name: "file1"},
			},
			removeArg:       &Argument{Name: "file2"},
			expectedCount:   1,
			shouldBeRemoved: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			group := &Group{Arguments: tt.setupArgs}

			// For the "remove existing" test, use the actual argument from the slice
			if tt.shouldBeRemoved && len(tt.setupArgs) > 0 {
				tt.removeArg = tt.setupArgs[1]
			}

			group.RemoveArgument(tt.removeArg)

			if len(group.Arguments) != tt.expectedCount {
				t.Errorf("RemoveArgument() resulted in %d arguments, want %d", len(group.Arguments), tt.expectedCount)
			}
		})
	}
}

func TestGroup_CalculateOptionWidths(t *testing.T) {
	tests := []struct {
		name                string
		options             []*Option
		expectedShort       int
		expectedLong        int
		expectedDefault     int
		expectedDescription int
	}{
		{
			name:                "empty options",
			options:             []*Option{},
			expectedShort:       0,
			expectedLong:        0,
			expectedDefault:     0,
			expectedDescription: 0,
		},
		{
			name: "single option",
			options: []*Option{
				{
					Short:       "v",
					Long:        "verbose",
					Default:     false,
					Description: "Enable verbose output",
				},
			},
			expectedShort:       1,
			expectedLong:        7,
			expectedDefault:     5, // "false"
			expectedDescription: 21,
		},
		{
			name: "multiple options with varying widths",
			options: []*Option{
				{
					Short:       "v",
					Long:        "verbose",
					Default:     false,
					Description: "Enable verbose",
				},
				{
					Short:       "o",
					Long:        "output-file",
					Default:     "default.txt",
					Description: "Output",
				},
				{
					Short:       "p",
					Long:        "port",
					Default:     8080,
					Description: "Port number to use",
				},
			},
			expectedShort:       1,
			expectedLong:        11, // "output-file"
			expectedDefault:     11, // "default.txt"
			expectedDescription: 18, // "Port number to use" is actually 18 characters
		},
		{
			name: "options with integer defaults",
			options: []*Option{
				{
					Short:       "p",
					Long:        "port",
					Default:     12345,
					Description: "Port",
				},
			},
			expectedShort:       1,
			expectedLong:        4,
			expectedDefault:     5, // "12345"
			expectedDescription: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			group := &Group{Options: tt.options}

			shortW, longW, defaultW, descW := group.CalculateOptionWidths()

			if shortW != tt.expectedShort {
				t.Errorf("CalculateOptionWidths() shortWidth = %d, want %d", shortW, tt.expectedShort)
			}
			if longW != tt.expectedLong {
				t.Errorf("CalculateOptionWidths() longWidth = %d, want %d", longW, tt.expectedLong)
			}
			if defaultW != tt.expectedDefault {
				t.Errorf("CalculateOptionWidths() defaultWidth = %d, want %d", defaultW, tt.expectedDefault)
			}
			if descW != tt.expectedDescription {
				t.Errorf("CalculateOptionWidths() descriptionWidth = %d, want %d", descW, tt.expectedDescription)
			}
		})
	}
}
