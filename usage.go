package usage

import (
	"fmt"
	"github.com/bgrewell/usage/internal"
)

type UsageOption func(sage *Usage)

func WithApplicationName(name string) UsageOption {
	return func(u *Usage) {
		u.applicationName = name
	}
}

func WithApplicationVersion(version string) UsageOption {
	return func(u *Usage) {
		u.applicationVersion = version
	}
}

func WithApplicationBuild(build string) UsageOption {
	return func(u *Usage) {
		u.applicationBuild = build
	}
}

func WithApplicationRevision(revision string) UsageOption {
	return func(u *Usage) {
		u.applicationRevision = revision
	}
}

func WithApplicationBranch(branch string) UsageOption {
	return func(u *Usage) {
		u.applicationBranch = branch
	}
}

func WithApplicationDescription(description string) UsageOption {
	return func(u *Usage) {
		u.applicationDescription = description
	}
}

func NewUsage(options ...UsageOption) *Usage {
	u := &Usage{
		applicationName: internal.GetExecutableName(),
	}
	for _, opt := range options {
		opt(u)
	}
	return u
}

type Usage struct {
	applicationName        string
	applicationVersion     string
	applicationBuild       string
	applicationRevision    string
	applicationBranch      string
	applicationDescription string
	groups                 []Group
}

func (s *Usage) ApplicationName() string {
	return s.applicationName
}

func (s *Usage) ApplicationVersion() string {
	return s.applicationVersion
}

func (s *Usage) ApplicationBuild() string {
	return s.applicationBuild
}

func (s *Usage) ApplicationRevision() string {
	return s.applicationRevision
}

func (s *Usage) ApplicationBranch() string {
	return s.applicationBranch
}

func (s *Usage) ApplicationDescription() string {
	return s.applicationDescription
}

func (s *Usage) Usage() {
	fmt.Println("Peanut Butter")
}
