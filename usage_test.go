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
	sage := usage.NewUsage(usage.WithApplicationBuildDate("100"))
	assert.Equal(t, "100", sage.ApplicationBuildDate())
}

func TestWithApplicationRevision(t *testing.T) {
	sage := usage.NewUsage(usage.WithApplicationCommitHash("abc123"))
	assert.Equal(t, "abc123", sage.ApplicationCommitHash())
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
		usage.WithApplicationBuildDate("100"),
		usage.WithApplicationCommitHash("abc123"),
		usage.WithApplicationBranch("main"),
		usage.WithApplicationDescription("Test Description"),
	)
	assert.Equal(t, "TestApp", sage.ApplicationName())
	assert.Equal(t, "1.0.0", sage.ApplicationVersion())
	assert.Equal(t, "100", sage.ApplicationBuildDate())
	assert.Equal(t, "abc123", sage.ApplicationCommitHash())
	assert.Equal(t, "main", sage.ApplicationBranch())
	assert.Equal(t, "Test Description", sage.ApplicationDescription())
}

func TestNewUseSageWithNoOptions(t *testing.T) {
	sage := usage.NewUsage()
	assert.NotEmpty(t, sage.ApplicationName())
	assert.Empty(t, sage.ApplicationVersion())
	assert.Empty(t, sage.ApplicationBuildDate())
	assert.Empty(t, sage.ApplicationCommitHash())
	assert.Empty(t, sage.ApplicationBranch())
	assert.Empty(t, sage.ApplicationDescription())
}
