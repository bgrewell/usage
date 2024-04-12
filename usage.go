package usage

import "github.com/bgrewell/usage/pkg"

func NewUseSage() UseSage {
	return &pkg.StandardUseSage{}
}

type UseSage interface {
}
