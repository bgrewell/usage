package internal

import (
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/fatih/color"
)

type ColorFormatter struct {
	Output        io.Writer
	Error         io.Writer
	Configuration *Configuration
}

func (f *ColorFormatter) PrintUsage() {
	if f.Output == nil {
		f.Output = os.Stdout
	}

	// Define colors
	lineColor := color.New(color.FgHiWhite)
	usageColor := color.New(color.FgHiBlue, color.Bold)
	headerColor := color.New(color.FgHiBlue, color.Bold)
	optionHeaderColor := color.New(color.FgHiBlue)
	optionDescColor := color.New(color.FgWhite)
	optionColor := color.New(color.FgGreen)
	optionDefaultColor := color.New(color.FgHiCyan)

	// Print the usage line with colors
	usageColor.Fprint(f.Output, "Usage: ")
	lineColor.Fprintf(f.Output, "%s [OPTIONS] [ARGUMENTS]\n\n", f.Configuration.ApplicationName)

	// Print the version information if it is provided
	versionInfoPresent := false
	if f.Configuration.ApplicationVersion != "" {
		headerColor.Fprintf(f.Output, "Version: ")
		lineColor.Fprintf(f.Output, "%s\n", f.Configuration.ApplicationVersion)
		versionInfoPresent = true
	}

	// Print the build date information if it is provided
	if f.Configuration.ApplicationBuildDate != "" {
		headerColor.Fprintf(f.Output, "Date: ")
		lineColor.Fprintf(f.Output, "%s\n", f.Configuration.ApplicationBuildDate)
		versionInfoPresent = true
	}

	// Print the commit hash information if it is provided
	if f.Configuration.ApplicationCommitHash != "" {
		headerColor.Fprintf(f.Output, "Codebase: ")
		lineColor.Fprintf(f.Output, "%s", f.Configuration.ApplicationCommitHash)
		if f.Configuration.ApplicationBranch != "" {
			lineColor.Fprintf(f.Output, " (%s)", f.Configuration.ApplicationBranch)
		}
		fmt.Fprintln(f.Output, "")
		versionInfoPresent = true
	}
	if versionInfoPresent {
		fmt.Fprintln(f.Output, "")
	}

	// Wrap the description text
	wrappedDescription := wrapText(f.Configuration.ApplicationDescription, 60, len("Description: "))

	// Print the description if one is provided
	if len(wrappedDescription) > 0 {
		headerColor.Fprint(f.Output, "Description: ")
		lineColor.Fprintf(f.Output, "%s\n", wrappedDescription[0])
		for _, line := range wrappedDescription[1:] {
			lineColor.Fprintf(f.Output, "  %s\n", line)
		}
		fmt.Fprintln(f.Output, "")
	}

	headerColor.Fprintln(f.Output, "Options:")
	for _, group := range f.Configuration.Groups {
		sw, lw, dvw, dsw := group.CalculateOptionWidths()
		optionHeaderColor.Fprintf(f.Output, "  %s: ", group.Name)
		lineColor.Fprintf(f.Output, "%s\n", group.Description)
		for _, option := range group.Options {
			optionColor.Fprintf(f.Output, "    -%-*s, --%-*s", sw, option.Short, lw, option.Long)
			optionDefault := option.Default
			if optionDefault == "" {
				optionDefault = "-"
			}
			optionDefaultColor.Fprintf(f.Output, "  %-*v", dvw, optionDefault)
			optionDescColor.Fprintf(f.Output, "  %-*s\n", dsw, option.Description)
		}
		fmt.Fprintln(f.Output, "")
	}

	headerColor.Fprintln(f.Output, "Arguments:")
	for _, group := range f.Configuration.Groups {
		arguments := group.Arguments
		// Sort the arguments by position
		sort.Slice(arguments, func(i, j int) bool {
			return arguments[i].Position < arguments[j].Position
		})
		for _, argument := range arguments {
			optionColor.Fprintf(f.Output, "    %s", argument.Name)
			optionDescColor.Fprintf(f.Output, "  %s\n", argument.Description)
		}
	}
}

func (f *ColorFormatter) PrintError(err error) {
	if f.Error == nil {
		f.Error = os.Stderr
	}

	errColor := color.New(color.FgHiRed)
	errColor.Fprintln(f.Error, "[!] Error: ", err)
	f.PrintUsage()
}
