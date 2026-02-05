package internal

// Option represents a command-line flag with short and long forms.
// Options are registered with Go's flag package and displayed in usage output.
type Option struct {
	Short       string      // Single-character flag name (e.g., "v" for -v)
	Long        string      // Long flag name (e.g., "verbose" for --verbose)
	Default     interface{} // Default value if the flag is not provided
	Description string      // Help text describing the option
	Extra       string      // Additional information shown in usage output
}
