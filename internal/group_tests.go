package internal

import (
	"testing"
)

func TestGroupAddOption(t *testing.T) {
	group := &Group{}
	option := &Option{}

	group.AddOption(option)

	if len(group.Options) != 1 {
		t.Errorf("AddOption() = %v; want 1", len(group.Options))
	}
}

func TestGroupRemoveOption(t *testing.T) {
	option := &Option{}
	group := &Group{Options: []*Option{option}}

	group.RemoveOption(option)

	if len(group.Options) != 0 {
		t.Errorf("RemoveOption() = %v; want 0", len(group.Options))
	}
}

func TestGroupAddArgument(t *testing.T) {
	group := &Group{}
	argument := &Argument{}

	group.AddArgument(argument)

	if len(group.Arguments) != 1 {
		t.Errorf("AddArgument() = %v; want 1", len(group.Arguments))
	}
}

func TestGroupRemoveArgument(t *testing.T) {
	argument := &Argument{}
	group := &Group{Arguments: []*Argument{argument}}

	group.RemoveArgument(argument)

	if len(group.Arguments) != 0 {
		t.Errorf("RemoveArgument() = %v; want 0", len(group.Arguments))
	}
}
