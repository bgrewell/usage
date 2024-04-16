package pkg_test

import (
	"github.com/bgrewell/usage"
	"github.com/bgrewell/usage/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUseSage(t *testing.T) {
	sage := usage.NewUseSage()
	assert.NotNil(t, sage)
}

func TestWithApplicationName(t *testing.T) {
	sage := usage.NewUseSage(usage.WithApplicationName("TestApp"))
	assert.Equal(t, "TestApp", sage.ApplicationName())
}

func TestWithApplicationVersion(t *testing.T) {
	sage := usage.NewUseSage(usage.WithApplicationVersion("1.0.0"))
	assert.Equal(t, "1.0.0", sage.ApplicationVersion())
}

func TestWithApplicationBuild(t *testing.T) {
	sage := usage.NewUseSage(usage.WithApplicationBuild("100"))
	assert.Equal(t, "100", sage.ApplicationBuild())
}

func TestWithApplicationRevision(t *testing.T) {
	sage := usage.NewUseSage(usage.WithApplicationRevision("abc123"))
	assert.Equal(t, "abc123", sage.ApplicationRevision())
}

func TestWithApplicationBranch(t *testing.T) {
	sage := usage.NewUseSage(usage.WithApplicationBranch("main"))
	assert.Equal(t, "main", sage.ApplicationBranch())
}

func TestNewGroup(t *testing.T) {
	group := usage.NewGroup("TestGroup")
	assert.Equal(t, "TestGroup", group.Title())
}

func TestApplicationNameSettingAndGetting(t *testing.T) {
	sage := pkg.StandardUseSage{}
	sage.SetApplicationName("TestApp")
	assert.Equal(t, "TestApp", sage.ApplicationName())
}

func TestApplicationVersionSettingAndGetting(t *testing.T) {
	sage := pkg.StandardUseSage{}
	sage.SetApplicationVersion("1.0.0")
	assert.Equal(t, "1.0.0", sage.ApplicationVersion())
}

func TestApplicationBuildSettingAndGetting(t *testing.T) {
	sage := pkg.StandardUseSage{}
	sage.SetApplicationBuild("100")
	assert.Equal(t, "100", sage.ApplicationBuild())
}

func TestApplicationRevisionSettingAndGetting(t *testing.T) {
	sage := pkg.StandardUseSage{}
	sage.SetApplicationRevision("abc123")
	assert.Equal(t, "abc123", sage.ApplicationRevision())
}

func TestApplicationBranchSettingAndGetting(t *testing.T) {
	sage := pkg.StandardUseSage{}
	sage.SetApplicationBranch("main")
	assert.Equal(t, "main", sage.ApplicationBranch())
}
