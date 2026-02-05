package pkg

import "github.com/bgrewell/usage/internal"

// NewGroup creates a new option group with the specified priority, name, and description.
// Groups are used to organize related options in the usage output and are displayed
// in order of priority (lower priority numbers appear first).
//
// Parameters:
//   - priority: determines the display order (lower values appear first)
//   - name: unique identifier for the group
//   - description: help text describing the purpose of the group
//
// Returns a Group that can be passed to Add*Option methods.
func NewGroup(priority int, name, description string) *internal.Group {
	return &internal.Group{
		Priority:    priority,
		Name:        name,
		Description: description,
	}
}
