package usage

func NewGroup(title string) *Group {
	return &Group{
		title: title,
	}
}

type Group struct {
	title string
}
