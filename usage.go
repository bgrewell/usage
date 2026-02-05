package usage

import (
	"errors"
	"flag"
	"fmt"
	"github.com/bgrewell/usage/internal"
	"github.com/bgrewell/usage/pkg"
	"log"
	"os"
)

const (
	GROUP_DEFAULT = "Default"
)

var (
	// ErrGroupNotFound is returned when attempting to add an option to a non-existent group.
	ErrGroupNotFound = errors.New("group does not exist")
)

type UsageOption func(sage *Usage)

func WithApplicationName(name string) UsageOption {
	return func(u *Usage) {
		u.configuration.ApplicationName = name
	}
}

func WithApplicationVersion(version string) UsageOption {
	return func(u *Usage) {
		u.configuration.ApplicationVersion = version
	}
}

func WithApplicationBuildDate(date string) UsageOption {
	return func(u *Usage) {
		u.configuration.ApplicationBuildDate = date
	}
}

func WithApplicationCommitHash(commitHash string) UsageOption {
	return func(u *Usage) {
		u.configuration.ApplicationCommitHash = commitHash
	}
}

func WithApplicationBranch(branch string) UsageOption {
	return func(u *Usage) {
		u.configuration.ApplicationBranch = branch
	}
}

func WithApplicationDescription(description string) UsageOption {
	return func(u *Usage) {
		u.configuration.ApplicationDescription = description
	}
}

func WithFormatter(formatter internal.Formatter) UsageOption {
	return func(u *Usage) {
		u.formatter = formatter
	}
}

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

type Usage struct {
	configuration *internal.Configuration
	formatter     internal.Formatter
	arguments     []*string
}

func (s *Usage) ApplicationName() string {
	return s.configuration.ApplicationName
}

func (s *Usage) ApplicationVersion() string {
	return s.configuration.ApplicationVersion
}

func (s *Usage) ApplicationBuildDate() string {
	return s.configuration.ApplicationBuildDate
}

func (s *Usage) ApplicationCommitHash() string {
	return s.configuration.ApplicationCommitHash
}

func (s *Usage) ApplicationBranch() string {
	return s.configuration.ApplicationBranch
}

func (s *Usage) ApplicationDescription() string {
	return s.configuration.ApplicationDescription
}

func (s *Usage) AddGroup(priority int, name string, description string) *internal.Group {
	group := pkg.NewGroup(priority, name, description)
	s.configuration.Groups[name] = group
	return group
}

// addOptionE adds an option to a group and returns an error if the group doesn't exist.
// This is the error-returning version of addOption.
func (s *Usage) addOptionE(short string, long string, defaultValue interface{}, description string, extra string, group *internal.Group) error {
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
		return nil
	}
	return fmt.Errorf("%w: %s", ErrGroupNotFound, group.Name)
}

// addOption is a backward-compatible wrapper that panics on error.
// Deprecated: Use addOptionE for proper error handling.
func (s *Usage) addOption(short string, long string, defaultValue interface{}, description string, extra string, group *internal.Group) {
	err := s.addOptionE(short, long, defaultValue, description, extra, group)
	if err != nil {
		log.Fatalf("%v", err)
	}
}

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

// AddBooleanOptionE adds a boolean command-line flag and returns an error if the group doesn't exist.
// This is the error-returning version of AddBooleanOption that allows proper error handling.
// Parameters are the same as AddBooleanOption.
func (s *Usage) AddBooleanOptionE(short string, long string, defaultValue bool, description string, extra string, group *internal.Group) (*bool, error) {
	var flagBool bool
	if short != "" {
		flag.BoolVar(&flagBool, short, defaultValue, description)
	}
	if long != "" {
		flag.BoolVar(&flagBool, long, defaultValue, description)
	}
	if err := s.addOptionE(short, long, defaultValue, description, extra, group); err != nil {
		return nil, err
	}
	return &flagBool, nil
}

// AddIntegerOptionE adds an integer command-line flag and returns an error if the group doesn't exist.
// This is the error-returning version of AddIntegerOption that allows proper error handling.
// Parameters are the same as AddIntegerOption.
func (s *Usage) AddIntegerOptionE(short string, long string, defaultValue int, description string, extra string, group *internal.Group) (*int, error) {
	var flagInt int
	if short != "" {
		flag.IntVar(&flagInt, short, defaultValue, description)
	}
	if long != "" {
		flag.IntVar(&flagInt, long, defaultValue, description)
	}
	if err := s.addOptionE(short, long, defaultValue, description, extra, group); err != nil {
		return nil, err
	}
	return &flagInt, nil
}

// AddFloatOptionE adds a float64 command-line flag and returns an error if the group doesn't exist.
// This is the error-returning version of AddFloatOption that allows proper error handling.
// Parameters are the same as AddFloatOption.
func (s *Usage) AddFloatOptionE(short string, long string, defaultValue float64, description string, extra string, group *internal.Group) (*float64, error) {
	var flagFloat float64
	if short != "" {
		flag.Float64Var(&flagFloat, short, defaultValue, description)
	}
	if long != "" {
		flag.Float64Var(&flagFloat, long, defaultValue, description)
	}
	if err := s.addOptionE(short, long, defaultValue, description, extra, group); err != nil {
		return nil, err
	}
	return &flagFloat, nil
}

// AddStringOptionE adds a string command-line flag and returns an error if the group doesn't exist.
// This is the error-returning version of AddStringOption that allows proper error handling.
// Parameters are the same as AddStringOption.
func (s *Usage) AddStringOptionE(short string, long string, defaultValue string, description string, extra string, group *internal.Group) (*string, error) {
	var flagString string
	if short != "" {
		flag.StringVar(&flagString, short, defaultValue, description)
	}
	if long != "" {
		flag.StringVar(&flagString, long, defaultValue, description)
	}
	if err := s.addOptionE(short, long, defaultValue, description, extra, group); err != nil {
		return nil, err
	}
	return &flagString, nil
}

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

func (s *Usage) PrintUsage() {
	s.formatter.PrintUsage()
	os.Exit(0)
}

func (s *Usage) PrintError(err error) {
	s.formatter.PrintError(err)
	os.Exit(1)
}
