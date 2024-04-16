package pkg

type StandardGroup struct {
	title string
}

func (g *StandardGroup) Title() string {
	return g.title
}

func (g *StandardGroup) SetTitle(title string) {
	g.title = title
}
