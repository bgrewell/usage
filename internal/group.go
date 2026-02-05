package internal

import "fmt"

// Group represents a collection of related command-line options and arguments.
// Groups are used to organize options in the usage output and control their
// display order via the Priority field.
type Group struct {
	Priority    int
	Name        string
	Description string
	Options     []*Option
	Arguments   []*Argument
}

// AddOption adds an option to this group.
func (g *Group) AddOption(option *Option) {
	g.Options = append(g.Options, option)
}

// RemoveOption removes the first occurrence of the specified option from this group.
func (g *Group) RemoveOption(option *Option) {
	for i, o := range g.Options {
		if o == option {
			g.Options = append(g.Options[:i], g.Options[i+1:]...)
			break
		}
	}
}

// AddArgument adds a positional argument to this group.
func (g *Group) AddArgument(argument *Argument) {
	g.Arguments = append(g.Arguments, argument)
}

// RemoveArgument removes the first occurrence of the specified argument from this group.
func (g *Group) RemoveArgument(argument *Argument) {
	for i, a := range g.Arguments {
		if a == argument {
			g.Arguments = append(g.Arguments[:i], g.Arguments[i+1:]...)
			break
		}
	}
}

// CalculateOptionWidths calculates the maximum width of each column in the options
// display for proper alignment. Returns the widths for short names, long names,
// default values, and descriptions respectively.
func (g *Group) CalculateOptionWidths() (shortWidth, longWidth, defaultValueWidth, descriptionWidth int) {
	for _, option := range g.Options {
		if len(option.Short) > shortWidth {
			shortWidth = len(option.Short)
		}
		if len(option.Long) > longWidth {
			longWidth = len(option.Long)
		}
		defaultValueStr := fmt.Sprintf("%v", option.Default)
		if len(defaultValueStr) > defaultValueWidth {
			defaultValueWidth = len(defaultValueStr)
		}
		if len(option.Description) > descriptionWidth {
			descriptionWidth = len(option.Description)
		}
	}
	return
}
