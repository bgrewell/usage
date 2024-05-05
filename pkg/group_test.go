package pkg

import (
	"testing"
)

func NewGroupCreation(t *testing.T) {
	name := "testGroup"
	description := "This is a test group"
	group := NewGroup(name, description)

	if group.Name != name {
		t.Errorf("NewGroup() = %v; want %v", group.Name, name)
	}

	if group.Description != description {
		t.Errorf("NewGroup() = %v; want %v", group.Description, description)
	}
}

func NewGroupEmptyName(t *testing.T) {
	name := ""
	description := "This is a test group"
	group := NewGroup(name, description)

	if group.Name != name {
		t.Errorf("NewGroup() = %v; want %v", group.Name, name)
	}
}

func NewGroupEmptyDescription(t *testing.T) {
	name := "testGroup"
	description := ""
	group := NewGroup(name, description)

	if group.Description != description {
		t.Errorf("NewGroup() = %v; want %v", group.Description, description)
	}
}
