package internal

type Group struct {
	Name        string
	Description string
	Options     []*Option
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
