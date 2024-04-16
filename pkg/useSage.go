package pkg

type UseSage interface {
	ApplicationName() string
	SetApplicationName(name string)
	ApplicationVersion() string
	SetApplicationVersion(version string)
	ApplicationBuild() string
	SetApplicationBuild(build string)
	ApplicationRevision() string
	SetApplicationRevision(revision string)
	ApplicationBranch() string
	SetApplicationBranch(branch string)
	ApplicationDescription() string
	SetApplicationDescription(description string)
	Groups() []Group
	AddGroup(group Group)
	Usage()
}
