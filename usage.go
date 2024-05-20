package usage

import (
	"flag"
	"github.com/bgrewell/usage/internal"
	"github.com/bgrewell/usage/pkg"
	"log"
	"os"
)

const (
	GROUP_DEFAULT = "Default"
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
			GROUP_DEFAULT: pkg.NewGroup(GROUP_DEFAULT, "Default Options"),
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

func (s *Usage) AddGroup(name string, description string) *internal.Group {
	group := pkg.NewGroup(name, description)
	s.configuration.Groups[name] = group
	return group
}

func (s *Usage) AddBooleanOption(short string, long string, default_value bool, description string, extra string, group *internal.Group) *bool {
	// Add flags
	var flagBool bool
	flag.BoolVar(&flagBool, short, default_value, description)
	flag.BoolVar(&flagBool, long, default_value, description)

	// Create option
	o := internal.Option{
		Short:       short,
		Long:        long,
		Default:     default_value,
		Description: description,
		Extra:       extra,
	}

	// Add option to group
	if group == nil {
		group = s.configuration.Groups[GROUP_DEFAULT]
	}

	if g, ok := s.configuration.Groups[group.Name]; ok {
		g.AddOption(&o)
	} else {
		log.Fatalf("Group %s does not exist", group.Name)
	}

	// Return the flag
	return &flagBool
}

func (s *Usage) AddIntegerOption(short string, long string, default_value int, description string, extra string, group *internal.Group) *int {
	var flagInt int
	flag.IntVar(&flagInt, short, default_value, description)
	flag.IntVar(&flagInt, long, default_value, description)

	// Create option
	o := internal.Option{
		Short:       short,
		Long:        long,
		Default:     default_value,
		Description: description,
		Extra:       extra,
	}

	// Add option to group
	if group == nil {
		group = s.configuration.Groups[GROUP_DEFAULT]
	}

	if g, ok := s.configuration.Groups[group.Name]; ok {
		g.AddOption(&o)
	} else {
		log.Fatalf("Group %s does not exist", group.Name)
	}

	return &flagInt
}

func (s *Usage) AddFloatOption(short string, long string, default_value float64, description string, extra string, group *internal.Group) *float64 {
	var flagFloat float64
	flag.Float64Var(&flagFloat, short, default_value, description)
	flag.Float64Var(&flagFloat, long, default_value, description)

	// Create option
	o := internal.Option{
		Short:       short,
		Long:        long,
		Default:     default_value,
		Description: description,
		Extra:       extra,
	}

	// Add option to group
	if group == nil {
		group = s.configuration.Groups[GROUP_DEFAULT]
	}

	if g, ok := s.configuration.Groups[group.Name]; ok {
		g.AddOption(&o)
	} else {
		log.Fatalf("Group %s does not exist", group.Name)
	}

	return &flagFloat
}

func (s *Usage) AddStringOption(short string, long string, default_value string, description string, extra string, group *internal.Group) *string {
	var flagString string
	flag.StringVar(&flagString, short, default_value, description)
	flag.StringVar(&flagString, long, default_value, description)

	// Create option
	o := internal.Option{
		Short:       short,
		Long:        long,
		Default:     default_value,
		Description: description,
		Extra:       extra,
	}

	// Add option to group
	if group == nil {
		group = s.configuration.Groups[GROUP_DEFAULT]
	}

	if g, ok := s.configuration.Groups[group.Name]; ok {
		g.AddOption(&o)
	} else {
		log.Fatalf("Group %s does not exist", group.Name)
	}

	return &flagString
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
