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

func TestFormatUsage(t *testing.T) {
	sage := usage.NewUsage(
		usage.WithApplicationName("TestApp"),
		usage.WithApplicationDescription("Test Description"),
	)
	formatted := sage.FormatUsage()
	assert.Contains(t, formatted, "TestApp")
	assert.Contains(t, formatted, "Test Description")
	assert.Contains(t, formatted, "Usage:")
}

func TestPrintUsageWithoutExit(t *testing.T) {
	sage := usage.NewUsage(
		usage.WithApplicationName("TestApp"),
	)
	// This should not panic or exit
	sage.PrintUsageWithoutExit()
}

func TestPrintErrorWithoutExit(t *testing.T) {
	sage := usage.NewUsage(
		usage.WithApplicationName("TestApp"),
	)
	// This should not panic or exit
	sage.PrintErrorWithoutExit(assert.AnError)
}

func TestWithExitOnHelp(t *testing.T) {
	// We can't easily test the actual exit behavior, but we can test that the option is set
	sage := usage.NewUsage(
		usage.WithApplicationName("TestApp"),
		usage.WithExitOnHelp(false),
	)
	// Verify it was created successfully
	assert.NotNil(t, sage)
	// The PrintUsageWithoutExit should work fine
	sage.PrintUsageWithoutExit()
}

func TestWithExitOnError(t *testing.T) {
	// We can't easily test the actual exit behavior, but we can test that the option is set
	sage := usage.NewUsage(
		usage.WithApplicationName("TestApp"),
		usage.WithExitOnError(false),
	)
	// Verify it was created successfully
	assert.NotNil(t, sage)
	// The PrintErrorWithoutExit should work fine
	sage.PrintErrorWithoutExit(assert.AnError)
}
