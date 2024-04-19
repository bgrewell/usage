package usage_test

import (
	"github.com/bgrewell/usage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUseSage(t *testing.T) {
	sage := usage.NewUsage()
	assert.NotNil(t, sage)
}

func TestWithApplicationName(t *testing.T) {
	sage := usage.NewUsage(usage.WithApplicationName("TestApp"))
	assert.Equal(t, "TestApp", sage.ApplicationName())
}

func TestWithApplicationVersion(t *testing.T) {
	sage := usage.NewUsage(usage.WithApplicationVersion("1.0.0"))
	assert.Equal(t, "1.0.0", sage.ApplicationVersion())
}

func TestWithApplicationBuild(t *testing.T) {
	sage := usage.NewUsage(usage.WithApplicationBuild("100"))
	assert.Equal(t, "100", sage.ApplicationBuild())
}

func TestWithApplicationRevision(t *testing.T) {
	sage := usage.NewUsage(usage.WithApplicationRevision("abc123"))
	assert.Equal(t, "abc123", sage.ApplicationRevision())
}

func TestWithApplicationBranch(t *testing.T) {
	sage := usage.NewUsage(usage.WithApplicationBranch("main"))
	assert.Equal(t, "main", sage.ApplicationBranch())
}

func TestWithApplicationDescription(t *testing.T) {
	sage := usage.NewUsage(usage.WithApplicationDescription("Test Description"))
	assert.Equal(t, "Test Description", sage.ApplicationDescription())
}

func TestNewUseSageWithMultipleOptions(t *testing.T) {
	sage := usage.NewUsage(
		usage.WithApplicationName("TestApp"),
		usage.WithApplicationVersion("1.0.0"),
		usage.WithApplicationBuild("100"),
		usage.WithApplicationRevision("abc123"),
		usage.WithApplicationBranch("main"),
		usage.WithApplicationDescription("Test Description"),
	)
	assert.Equal(t, "TestApp", sage.ApplicationName())
	assert.Equal(t, "1.0.0", sage.ApplicationVersion())
	assert.Equal(t, "100", sage.ApplicationBuild())
	assert.Equal(t, "abc123", sage.ApplicationRevision())
	assert.Equal(t, "main", sage.ApplicationBranch())
	assert.Equal(t, "Test Description", sage.ApplicationDescription())
}

func TestNewUseSageWithNoOptions(t *testing.T) {
	sage := usage.NewUsage()
	assert.NotEmpty(t, sage.ApplicationName())
	assert.Empty(t, sage.ApplicationVersion())
	assert.Empty(t, sage.ApplicationBuild())
	assert.Empty(t, sage.ApplicationRevision())
	assert.Empty(t, sage.ApplicationBranch())
	assert.Empty(t, sage.ApplicationDescription())
}
