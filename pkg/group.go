package pkg

import "github.com/bgrewell/usage/internal"

func NewGroup(priority int, name, description string) *internal.Group {
	return &internal.Group{
		Priority:    priority,
		Name:        name,
		Description: description,
	}
}
