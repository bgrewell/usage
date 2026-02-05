package internal

// Argument represents a positional command-line argument.
// Positional arguments are non-flag arguments that must appear in order.
type Argument struct {
	Position    int    // Expected position of this argument (0-indexed)
	Name        string // Name of the argument shown in usage output
	Description string // Help text describing the argument
	Extra       string // Additional information shown in usage output
}
