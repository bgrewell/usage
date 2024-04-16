package pkg_test

import (
	"github.com/bgrewell/usage/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGroupAddingAndGetting(t *testing.T) {
	sage := pkg.StandardUseSage{}
	group := &pkg.StandardGroup{}
	sage.AddGroup(group)
	assert.Equal(t, group, sage.Groups()[0])
}
