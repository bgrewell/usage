// Package usage provides a flexible command-line argument and option parsing library
// with support for custom formatters, option groups, and automatic usage generation.
package usage

import (
	"flag"
	"github.com/bgrewell/usage/internal"
	"github.com/bgrewell/usage/pkg"
	"log"
	"os"
)

const (
	// GROUP_DEFAULT is the name of the default group where options are placed
	// when no explicit group is specified.
	GROUP_DEFAULT = "Default"
)

// UsageOption is a functional option for configuring a Usage instance.
// It follows the functional options pattern to provide flexible configuration.
type UsageOption func(sage *Usage)

// WithApplicationName sets the application name displayed in usage output.
// If not provided, the executable name is automatically detected.
func WithApplicationName(name string) UsageOption {
	return func(u *Usage) {
		u.configuration.ApplicationName = name
	}
}

// WithApplicationVersion sets the application version displayed in usage output.
func WithApplicationVersion(version string) UsageOption {
	return func(u *Usage) {
		u.configuration.ApplicationVersion = version
	}
}

// WithApplicationBuildDate sets the build date displayed in usage output.
func WithApplicationBuildDate(date string) UsageOption {
	return func(u *Usage) {
		u.configuration.ApplicationBuildDate = date
	}
}

// WithApplicationCommitHash sets the git commit hash displayed in usage output.
func WithApplicationCommitHash(commitHash string) UsageOption {
	return func(u *Usage) {
		u.configuration.ApplicationCommitHash = commitHash
	}
}

// WithApplicationBranch sets the git branch displayed in usage output.
func WithApplicationBranch(branch string) UsageOption {
	return func(u *Usage) {
		u.configuration.ApplicationBranch = branch
	}
}

// WithApplicationDescription sets the application description displayed in usage output.
func WithApplicationDescription(description string) UsageOption {
	return func(u *Usage) {
		u.configuration.ApplicationDescription = description
	}
}

// WithFormatter sets a custom formatter for usage and error output.
// By default, a ColorFormatter is used. You can provide a StandardFormatter
// or implement your own custom formatter using the internal.Formatter interface.
func WithFormatter(formatter internal.Formatter) UsageOption {
	return func(u *Usage) {
		u.formatter = formatter
	}
}

// NewUsage creates a new Usage instance with the provided functional options.
// By default, it automatically detects the executable name, uses a ColorFormatter,
// and creates a default option group. The function also sets flag.Usage to
// the instance's PrintUsage method.
//
// Example:
//
//	u := usage.NewUsage(
//	    usage.WithApplicationName("myapp"),
//	    usage.WithApplicationVersion("1.0.0"),
//	    usage.WithApplicationDescription("A command-line tool"),
//	)
func NewUsage(options ...UsageOption) *Usage {
	c := &internal.Configuration{
		ApplicationName: internal.GetExecutableName(),
		Groups: map[string]*internal.Group{
			GROUP_DEFAULT: pkg.NewGroup(0, GROUP_DEFAULT, "Default Options"),
		},
	}
	u := &Usage{
		configuration: c,
		formatter:     pkg.NewColorFormatter(os.Stdout, os.Stderr, c),
	}
	for _, opt := range options {
		opt(u)
	}
	flag.Usage = u.PrintUsage
	return u
}

// Usage manages command-line options, arguments, and usage output formatting.
// It wraps Go's standard flag package with additional features like option groups,
// custom formatters, and automatic usage generation.
type Usage struct {
	configuration *internal.Configuration
	formatter     internal.Formatter
	arguments     []*string
}

// ApplicationName returns the configured application name.
func (s *Usage) ApplicationName() string {
	return s.configuration.ApplicationName
}

// ApplicationVersion returns the configured application version.
func (s *Usage) ApplicationVersion() string {
	return s.configuration.ApplicationVersion
}

// ApplicationBuildDate returns the configured application build date.
func (s *Usage) ApplicationBuildDate() string {
	return s.configuration.ApplicationBuildDate
}

// ApplicationCommitHash returns the configured git commit hash.
func (s *Usage) ApplicationCommitHash() string {
	return s.configuration.ApplicationCommitHash
}

// ApplicationBranch returns the configured git branch.
func (s *Usage) ApplicationBranch() string {
	return s.configuration.ApplicationBranch
}

// ApplicationDescription returns the configured application description.
func (s *Usage) ApplicationDescription() string {
	return s.configuration.ApplicationDescription
}

// AddGroup creates a new option group for organizing related options.
// Groups are displayed in order of priority (lower numbers first).
// The name must be unique, and the description is shown in the usage output.
// Returns the created group which can be passed to Add*Option methods.
func (s *Usage) AddGroup(priority int, name string, description string) *internal.Group {
	group := pkg.NewGroup(priority, name, description)
	s.configuration.Groups[name] = group
	return group
}

func (s *Usage) addOption(short string, long string, defaultValue interface{}, description string, extra string, group *internal.Group) {
	o := internal.Option{
		Short:       short,
		Long:        long,
		Default:     defaultValue,
		Description: description,
		Extra:       extra,
	}

	if group == nil {
		group = s.configuration.Groups[GROUP_DEFAULT]
	}

	if g, ok := s.configuration.Groups[group.Name]; ok {
		g.AddOption(&o)
	} else {
		log.Fatalf("Group %s does not exist", group.Name)
	}
}

// AddBooleanOption adds a boolean command-line flag.
// Parameters:
//   - short: single-character flag name (e.g., "v" for -v), or empty string to skip
//   - long: long flag name (e.g., "verbose" for --verbose), or empty string to skip
//   - defaultValue: the default value if the flag is not provided
//   - description: help text describing the option
//   - extra: additional information shown in usage output
//   - group: the group to add this option to, or nil for GROUP_DEFAULT
//
// Returns a pointer to the boolean value that will be populated by flag.Parse().
// The method registers the flag with Go's standard flag package.
func (s *Usage) AddBooleanOption(short string, long string, defaultValue bool, description string, extra string, group *internal.Group) *bool {
	var flagBool bool
	if short != "" {
		flag.BoolVar(&flagBool, short, defaultValue, description)
	}
	if long != "" {
		flag.BoolVar(&flagBool, long, defaultValue, description)
	}
	s.addOption(short, long, defaultValue, description, extra, group)
	return &flagBool
}

// AddIntegerOption adds an integer command-line flag.
// Parameters:
//   - short: single-character flag name (e.g., "p" for -p), or empty string to skip
//   - long: long flag name (e.g., "port" for --port), or empty string to skip
//   - defaultValue: the default value if the flag is not provided
//   - description: help text describing the option
//   - extra: additional information shown in usage output
//   - group: the group to add this option to, or nil for GROUP_DEFAULT
//
// Returns a pointer to the integer value that will be populated by flag.Parse().
// The method registers the flag with Go's standard flag package.
func (s *Usage) AddIntegerOption(short string, long string, defaultValue int, description string, extra string, group *internal.Group) *int {
	var flagInt int
	if short != "" {
		flag.IntVar(&flagInt, short, defaultValue, description)
	}
	if long != "" {
		flag.IntVar(&flagInt, long, defaultValue, description)
	}
	s.addOption(short, long, defaultValue, description, extra, group)
	return &flagInt
}

// AddFloatOption adds a float64 command-line flag.
// Parameters:
//   - short: single-character flag name (e.g., "r" for -r), or empty string to skip
//   - long: long flag name (e.g., "rate" for --rate), or empty string to skip
//   - defaultValue: the default value if the flag is not provided
//   - description: help text describing the option
//   - extra: additional information shown in usage output
//   - group: the group to add this option to, or nil for GROUP_DEFAULT
//
// Returns a pointer to the float64 value that will be populated by flag.Parse().
// The method registers the flag with Go's standard flag package.
func (s *Usage) AddFloatOption(short string, long string, defaultValue float64, description string, extra string, group *internal.Group) *float64 {
	var flagFloat float64
	if short != "" {
		flag.Float64Var(&flagFloat, short, defaultValue, description)
	}
	if long != "" {
		flag.Float64Var(&flagFloat, long, defaultValue, description)
	}
	s.addOption(short, long, defaultValue, description, extra, group)
	return &flagFloat
}

// AddStringOption adds a string command-line flag.
// Parameters:
//   - short: single-character flag name (e.g., "o" for -o), or empty string to skip
//   - long: long flag name (e.g., "output" for --output), or empty string to skip
//   - defaultValue: the default value if the flag is not provided
//   - description: help text describing the option
//   - extra: additional information shown in usage output
//   - group: the group to add this option to, or nil for GROUP_DEFAULT
//
// Returns a pointer to the string value that will be populated by flag.Parse().
// The method registers the flag with Go's standard flag package.
func (s *Usage) AddStringOption(short string, long string, defaultValue string, description string, extra string, group *internal.Group) *string {
	var flagString string
	if short != "" {
		flag.StringVar(&flagString, short, defaultValue, description)
	}
	if long != "" {
		flag.StringVar(&flagString, long, defaultValue, description)
	}
	s.addOption(short, long, defaultValue, description, extra, group)
	return &flagString
}

// AddArgument adds a positional command-line argument.
// Positional arguments are non-flag arguments that must appear in order.
// If this is the last declared argument, it will accumulate all remaining
// command-line arguments joined by spaces.
//
// Parameters:
//   - position: the expected position of this argument (0-indexed)
//   - name: the name of the argument shown in usage output
//   - description: help text describing the argument
//   - extra: additional information shown in usage output
//
// Returns a pointer to the string value that will be populated by Parse().
func (s *Usage) AddArgument(position int, name string, description string, extra string) *string {
	var argString string
	s.arguments = append(s.arguments, &argString)

	// Create argument
	a := internal.Argument{
		Position:    position,
		Name:        name,
		Description: description,
	}
	s.configuration.Groups[GROUP_DEFAULT].AddArgument(&a)

	return &argString
}

// Parse wraps flag.Parse() and populates the positional arguments.
// This method should be called after all options and arguments have been added.
// The last declared argument will accumulate all remaining command-line arguments.
// Returns true if parsing was successful (same as flag.Parsed()).
func (s *Usage) Parse() bool {
	flag.Parse()

	// Populate arguments
	for i, arg := range flag.Args() {
		// If this is the last argument in s.arguments then accumulate the rest of the arguments joined by a space
		if i >= len(s.arguments) {
			*s.arguments[len(s.arguments)-1] += " " + arg
		} else {
			*s.arguments[i] = arg
		}
	}

	return flag.Parsed()
}

// PrintUsage prints the usage information to the configured output writer
// and calls os.Exit(0). This is automatically set as flag.Usage in NewUsage.
func (s *Usage) PrintUsage() {
	s.formatter.PrintUsage()
	os.Exit(0)
}

// PrintError prints the error message and usage information to the configured
// error writer and calls os.Exit(1). This is typically used when command-line
// parsing or validation fails.
func (s *Usage) PrintError(err error) {
	s.formatter.PrintError(err)
	os.Exit(1)
}
