package internal

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
)

// StandardFormatter implements the Formatter interface using plain text
// without any color codes. This is suitable for non-terminal outputs or
// environments that don't support ANSI color codes.
type StandardFormatter struct {
	Output        io.Writer      // Writer for normal usage output (defaults to os.Stdout)
	Error         io.Writer      // Writer for error messages (defaults to os.Stderr)
	Configuration *Configuration // Application and option configuration
}

// FormatUsage returns the formatted usage information as a plain text string.
// This method does not print anything or mutate state, making it suitable
// for programmatic usage and testing.
func (f *StandardFormatter) FormatUsage() string {
	var buf bytes.Buffer

	// Print the usage line
	fmt.Fprintf(&buf, "Usage: %s [OPTIONS] [ARGUMENTS]\n\n", f.Configuration.ApplicationName)

	// Print the version information if it is provided
	versionInfoPresent := false
	if f.Configuration.ApplicationVersion != "" {
		fmt.Fprintf(&buf, "Version: %s\n", f.Configuration.ApplicationVersion)
		versionInfoPresent = true
	}

	// Print the build date information if it is provided
	if f.Configuration.ApplicationBuildDate != "" {
		fmt.Fprintf(&buf, "Date: %s\n", f.Configuration.ApplicationBuildDate)
		versionInfoPresent = true
	}

	// Print the commit hash information if it is provided
	if f.Configuration.ApplicationCommitHash != "" {
		fmt.Fprintf(&buf, "Codebase: %s", f.Configuration.ApplicationCommitHash)
		if f.Configuration.ApplicationBranch != "" {
			fmt.Fprintf(&buf, " (%s)", f.Configuration.ApplicationBranch)
		}
		fmt.Fprintln(&buf, "")
		versionInfoPresent = true
	}
	if versionInfoPresent {
		fmt.Fprintln(&buf, "")
	}

	// Print the description if one is provided
	if f.Configuration.ApplicationDescription != "" {
		fmt.Fprintf(&buf, "Description: %s\n\n", f.Configuration.ApplicationDescription)
	}

	// Get all the keys for the groups so we can get the priorities and then sort them
	groupKeys := make([]string, 0, len(f.Configuration.Groups))
	for key := range f.Configuration.Groups {
		groupKeys = append(groupKeys, key)
	}
	sort.Slice(groupKeys, func(i, j int) bool {
		return f.Configuration.Groups[groupKeys[i]].Priority < f.Configuration.Groups[groupKeys[j]].Priority
	})

	fmt.Fprintln(&buf, "Options:")
	for _, key := range groupKeys {
		group := f.Configuration.Groups[key]
		if len(group.Options) == 0 {
			continue
		}
		sw, lw, dvw, dsw := group.CalculateOptionWidths()
		fmt.Fprintf(&buf, "  %s: %s\n", group.Name, group.Description)
		for _, option := range group.Options {
			fmt.Fprintf(&buf, "    -%-*s --%-*s", sw, option.Short, lw, option.Long)
			optionDefault := option.Default
			if optionDefault == "" {
				optionDefault = "-"
			}
			fmt.Fprintf(&buf, "  %-*v", dvw, optionDefault)
			fmt.Fprintf(&buf, "  %-*s\n", dsw, option.Description)
		}
		fmt.Fprintln(&buf, "")
	}

	// Get a count of the number of arguments
	argumentCount := 0
	for _, group := range f.Configuration.Groups {
		argumentCount += len(group.Arguments)
	}

	if argumentCount > 0 {
		fmt.Fprintln(&buf, "Arguments:")
		for _, group := range f.Configuration.Groups {
			arguments := group.Arguments
			// Sort the arguments by position
			sort.Slice(arguments, func(i, j int) bool {
				return arguments[i].Position < arguments[j].Position
			})
			for _, argument := range arguments {
				fmt.Fprintf(&buf, "    %s  %s\n", argument.Name, argument.Description)
			}
		}
	}

	return buf.String()
}

// PrintUsage outputs formatted usage information in plain text.
// It displays application metadata (version, build date, commit hash),
// description, and options grouped by category.
// If Output is nil, it defaults to os.Stdout.
func (f *StandardFormatter) PrintUsage() {
	if f.Output == nil {
		f.Output = os.Stdout
	}
	fmt.Fprint(f.Output, f.FormatUsage())
}

// PrintError outputs the error message followed by the usage information.
// If Error is nil, it defaults to os.Stderr. The error message is displayed
// before the usage information.
func (f *StandardFormatter) PrintError(err error) {
	if f.Error == nil {
		f.Error = os.Stderr
	}

	fmt.Fprintln(f.Error, "Error: ", err)
	f.PrintUsage()
}
