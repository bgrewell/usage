package internal

import (
	"fmt"
	"io"
	"os"
)

type StandardFormatter struct {
	Output        io.Writer
	Error         io.Writer
	Configuration *Configuration
}

func (f *StandardFormatter) PrintUsage() {
	if f.Output == nil {
		f.Output = os.Stdout
	}

	// Print the usage line
	fmt.Printf("Usage: %s [OPTIONS] [ARGUMENTS]\n\n", f.Configuration.ApplicationName)

	// Print the description if one is provided
	if f.Configuration.ApplicationDescription != "" {
		fmt.Printf("Description: %s\n\n", f.Configuration.ApplicationDescription)

	}

	// Print the version information if it is provided
	if f.Configuration.ApplicationVersion != "" {
		fmt.Printf("Version: %s\n", f.Configuration.ApplicationVersion)
	}

	// Print the build date information if it is provided
	if f.Configuration.ApplicationBuildDate != "" {
		fmt.Printf("Date: %s\n", f.Configuration.ApplicationBuildDate)
	}

	// Print the commit hash information if it is provided
	if f.Configuration.ApplicationCommitHash != "" {
		fmt.Printf("Codebase: %s", f.Configuration.ApplicationCommitHash)
		if f.Configuration.ApplicationBranch != "" {
			fmt.Printf(" (%s)", f.Configuration.ApplicationBranch)
		}
		fmt.Println("")
	}
	fmt.Println("")

	fmt.Println("Options:")
	for _, group := range f.Configuration.Groups {
		fmt.Printf("  %s: %s\n", group.Name, group.Description)
		for _, option := range group.Options {
			fmt.Printf("    -%s, --%s\t\t%s\n", option.Short, option.Long, option.Description)
		}
		fmt.Println("")
	}
}

func (f *StandardFormatter) PrintError(err error) {
	if f.Error == nil {
		f.Error = os.Stderr
	}

	fmt.Fprintln(f.Error, "Error: ", err)
	f.PrintUsage()
}

// ---------------------------------------------------------------------------------------------------------------------
//
// <error_message_if_present>
// Usage: <app_name> [OPTIONS] [ARGUMENTS]
//
// Description: <app_description>
//
// Options:
//   <group_name_if_more_than_one>: <group_description_if_present>
//   -h, --help					Display this help message
//   -v, --version				Display the application version
//   -b, --build				Display the application build
//   -r, --revision				Display the application revision
//   -B, --branch				Display the application branch
//
// Examples:
//  <app_name> -x -y -z arg1 arg2
//    This is an example of how to use the application
//
//  <app_name> -a -b arg1
//    This is another example of how to use the application
//
// <post_output_message>
//
//
