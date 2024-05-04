package pkg

import "github.com/bgrewell/usage/internal"

func NewGroup(name, description string) *internal.Group {
	return &internal.Group{
		Name:        name,
		Description: description,
	}
}
