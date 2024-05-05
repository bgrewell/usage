package internal

type Group struct {
	Name        string
	Description string
	Options     []*Option
	Arguments   []*Argument
}

func (g *Group) AddOption(option *Option) {
	g.Options = append(g.Options, option)
}

func (g *Group) RemoveOption(option *Option) {
	for i, o := range g.Options {
		if o == option {
			g.Options = append(g.Options[:i], g.Options[i+1:]...)
			break
		}
	}
}

func (g *Group) AddArgument(argument *Argument) {
	g.Arguments = append(g.Arguments, argument)
}

func (g *Group) RemoveArgument(argument *Argument) {
	for i, a := range g.Arguments {
		if a == argument {
			g.Arguments = append(g.Arguments[:i], g.Arguments[i+1:]...)
			break
		}
	}
}
