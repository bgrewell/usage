package usage

import . "github.com/bgrewell/usage/pkg"

type UsageOption func(sage UseSage)

func NewUseSage(options ...UsageOption) UseSage {
	u := &StandardUseSage{}
	for _, opt := range options {
		opt(u)
	}
	return u
}

func WithApplicationName(name string) UsageOption {
	return func(sage UseSage) {
		sage.SetApplicationName(name)
	}
}

func WithApplicationVersion(version string) UsageOption {
	return func(sage UseSage) {
		sage.SetApplicationVersion(version)
	}
}

func WithApplicationBuild(build string) UsageOption {
	return func(sage UseSage) {
		sage.SetApplicationBuild(build)
	}
}

func WithApplicationRevision(revision string) UsageOption {
	return func(sage UseSage) {
		sage.SetApplicationRevision(revision)
	}
}

func WithApplicationBranch(branch string) UsageOption {
	return func(sage UseSage) {
		sage.SetApplicationBranch(branch)
	}
}

func WithApplicationDescription(description string) UsageOption {
	return func(sage UseSage) {
		sage.SetApplicationDescription(description)
	}
}

func NewGroup(title string) Group {
	g := StandardGroup{}
	g.SetTitle(title)
	return &g
}
