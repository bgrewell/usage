package pkg

import (
	"fmt"
	"github.com/bgrewell/usage/internal"
)

type StandardUseSage struct {
	applicationName     string
	applicationVersion  string
	applicationBuild    string
	applicationRevision string
	applicationBranch   string
	groups              []Group
}

func (s *StandardUseSage) ApplicationName() string {
	if s.applicationName != "" {
		return s.applicationName
	} else {
		return internal.GetExecutableName()
	}
}

func (s *StandardUseSage) SetApplicationName(name string) {
	s.applicationName = name
}
func (s *StandardUseSage) ApplicationVersion() string {
	return s.applicationVersion
}
func (s *StandardUseSage) SetApplicationVersion(version string) {
	s.applicationVersion = version
}
func (s *StandardUseSage) ApplicationBuild() string {
	return s.applicationBuild
}
func (s *StandardUseSage) SetApplicationBuild(build string) {
	s.applicationBuild = build
}
func (s *StandardUseSage) ApplicationRevision() string {
	return s.applicationRevision
}
func (s *StandardUseSage) SetApplicationRevision(revision string) {
	s.applicationRevision = revision
}
func (s *StandardUseSage) ApplicationBranch() string {
	return s.applicationBranch
}
func (s *StandardUseSage) SetApplicationBranch(branch string) {
	s.applicationBranch = branch
}

func (s *StandardUseSage) Groups() []Group {
	return s.groups
}

func (s *StandardUseSage) AddGroup(group Group) {
	s.groups = append(s.groups, group)
}

func (s *StandardUseSage) Usage() {
	fmt.Printf("Application: %s\n", s.ApplicationName())
	fmt.Printf("Version: %s\n", s.ApplicationVersion())
	fmt.Printf("Build: %s\n", s.ApplicationBuild())
	fmt.Printf("Revision: %s\n", s.ApplicationRevision())
	fmt.Printf("Branch: %s\n", s.ApplicationBranch())
}
