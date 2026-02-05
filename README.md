# Usage

[![golangci-lint](https://github.com/bgrewell/usage/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/bgrewell/usage/actions/workflows/golangci-lint.yml)
[![codecov](https://codecov.io/gh/bgrewell/usage/graph/badge.svg?token=MP7QKP53BG)](https://codecov.io/gh/bgrewell/usage)

A flexible, feature-rich command-line argument parsing library for Go with support for option groups, custom formatters, and comprehensive error handling.

## Features

- **Functional Options Pattern** - Clean, extensible configuration API
- **Colored Output** - Built-in colored help text with automatic plain-text fallback
- **Organized Option Groups** - Group related options with custom priorities
- **Dual Error Handling** - Both error-returning and panic methods available
- **Automatic Usage Generation** - Beautiful help text with version/build metadata
- **Positional Arguments** - First-class support for positional arguments
- **Lightweight** - Thin wrapper around Go's standard `flag` package

## Installation

```bash
go get github.com/bgrewell/usage
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/bgrewell/usage"
)

func main() {
    // Create usage instance with metadata
    u := usage.NewUsage(
        usage.WithApplicationName("myapp"),
        usage.WithApplicationVersion("1.0.0"),
        usage.WithApplicationDescription("A simple example application"),
    )

    // Add options with error handling
    verbose, err := u.AddBooleanOptionE("v", "verbose", false, "Enable verbose output", "", nil)
    if err != nil {
        panic(err)
    }

    output, err := u.AddStringOptionE("o", "output", "", "Output file path", "", nil)
    if err != nil {
        panic(err)
    }

    // Parse command-line arguments
    u.Parse()

    // Use parsed values
    if *verbose {
        fmt.Println("Verbose mode enabled")
    }
    if *output != "" {
        fmt.Printf("Output file: %s\n", *output)
    }
}
```

## Examples

### Application Metadata

Add version and build information that appears in your help output:

```go
u := usage.NewUsage(
    usage.WithApplicationName("myapp"),
    usage.WithApplicationVersion("1.2.3"),
    usage.WithApplicationBuildDate("2024-01-15"),
    usage.WithApplicationCommitHash("abc1234"),
    usage.WithApplicationBranch("main"),
    usage.WithApplicationDescription("A comprehensive example application"),
)
```

### Option Types

The library supports all common flag types:

```go
// Boolean flags
verbose, err := u.AddBooleanOptionE("v", "verbose", false, "Enable verbose output", "", nil)

// String flags
name, err := u.AddStringOptionE("n", "name", "default", "User name", "", nil)

// Integer flags
count, err := u.AddIntegerOptionE("c", "count", 10, "Number of items", "", nil)

// Float flags
ratio, err := u.AddFloatOptionE("r", "ratio", 1.5, "Scaling ratio", "", nil)
```

### Option Groups

Organize related options into named groups with custom priorities:

```go
// Create an option group (lower priority numbers appear first)
networkGroup := u.AddGroup(1, "Network Options", "Options related to network configuration")

// Add options to the group
host, err := u.AddStringOptionE("h", "host", "localhost", "Server host", "", networkGroup)
port, err := u.AddIntegerOptionE("p", "port", 8080, "Server port", "", networkGroup)
timeout, err := u.AddIntegerOptionE("t", "timeout", 30, "Connection timeout", "", networkGroup)

// Options without a group go to the default group
verbose, err := u.AddBooleanOptionE("v", "verbose", false, "Enable verbose output", "", nil)
```

### Error Handling

The library provides two styles of methods:

**Error-Returning Methods (Recommended):**

```go
verbose, err := u.AddBooleanOptionE("v", "verbose", false, "Enable verbose output", "", nil)
if err != nil {
    // Handle error appropriately
    log.Fatal(err)
}
```

**Panic Methods (Legacy):**

```go
// These methods panic on error - simpler but less flexible
verbose := u.AddBooleanOption("v", "verbose", false, "Enable verbose output", "", nil)
```

> **Note:** Error-returning methods (with `E` suffix) are recommended for better error handling and testing.

### Custom Formatters

Choose between colored and plain-text output:

```go
// Use colored output (default)
u := usage.NewUsage(
    usage.WithApplicationName("myapp"),
    usage.WithFormatter(usage.NewColorFormatter()),
)

// Use plain-text output
u := usage.NewUsage(
    usage.WithApplicationName("myapp"),
    usage.WithFormatter(usage.NewStandardFormatter()),
)
```

You can also implement your own formatter by satisfying the `Formatter` interface:

```go
type Formatter interface {
    FormatUsage(u *Usage) string
    FormatError(err error) string
}
```

### Positional Arguments

Add required or optional positional arguments:

```go
// Single required argument
inputFile := u.AddArgument(1, "input", "Input file path", "")

// Multiple arguments
inputFile := u.AddArgument(1, "input", "Input file path", "")
outputFile := u.AddArgument(2, "output", "Output file path", "")

// Last argument can accumulate remaining values
files := u.AddArgument(1, "files", "Files to process", "Extra")

// After parsing
u.Parse()
fmt.Println("Input file:", *inputFile)
```

## Example Output

Running the example program with `--help`:

```
Usage: bowser [OPTIONS] [ARGUMENTS]

Description: It's almost a browser but not quite. Instead
  it's just a example of how to use the 'usage' package. It is
  designed to show an example of how to use the package and a
  sample of what the output would look like.

Options:
  Default: Default Options
    -o --output  -  Output filename

  Request Options: Options related http request
    -u --user-agent    Bowser/0.0.1  The user agent to use
    -r --request-type  GET           The type of request to make
    -f --follow        false         Follow Redirects
    -t --timeout       10            Timeout in seconds

Arguments:
    url  The url of the page to retrieve
```

With colored output enabled, option groups and flags are highlighted for better readability.

## API Reference

### Core Types

- **`Usage`** - Main struct for managing CLI arguments
- **`Group`** - Container for organizing related options
- **`Formatter`** - Interface for custom output formatting

### Creating a Usage Instance

```go
func NewUsage(options ...UsageOption) *Usage
```

**Functional Options:**
- `WithApplicationName(name string)` - Set application name
- `WithApplicationVersion(version string)` - Set version string
- `WithApplicationBuildDate(date string)` - Set build date
- `WithApplicationCommitHash(hash string)` - Set git commit hash
- `WithApplicationBranch(branch string)` - Set git branch
- `WithApplicationDescription(desc string)` - Set description
- `WithFormatter(formatter Formatter)` - Set custom formatter
- `WithFlagSet(fs *flag.FlagSet)` - Use custom flag.FlagSet

### Adding Options

**Error-Returning Methods (Recommended):**
- `AddBooleanOptionE(short, long string, defaultValue bool, description, note string, group *Group) (*bool, error)`
- `AddStringOptionE(short, long string, defaultValue, description, note string, group *Group) (*string, error)`
- `AddIntegerOptionE(short, long string, defaultValue int, description, note string, group *Group) (*int, error)`
- `AddFloatOptionE(short, long string, defaultValue float64, description, note string, group *Group) (*float64, error)`

**Panic Methods (Legacy):**
- `AddBooleanOption(...)` - Same signature, panics on error
- `AddStringOption(...)` - Same signature, panics on error
- `AddIntegerOption(...)` - Same signature, panics on error
- `AddFloatOption(...)` - Same signature, panics on error

### Other Methods

- `AddGroup(priority int, name, description string) *Group` - Create option group
- `AddArgument(priority int, name, description, extra string) *string` - Add positional argument
- `Parse() bool` - Parse command-line arguments
- `PrintUsage()` - Print formatted help text
- `PrintError(err error)` - Print formatted error message

**Full API Documentation:** https://pkg.go.dev/github.com/bgrewell/usage

## Migration Guide

### Migrating to Error-Returning Methods

The library now provides error-returning methods (with `E` suffix) for better error handling and testability:

**Old (still works, but not recommended):**
```go
verbose := u.AddBooleanOption("v", "verbose", false, "Enable verbose output", "", nil)
port := u.AddIntegerOption("p", "port", 8080, "Server port", "", nil)
```

**New (recommended):**
```go
verbose, err := u.AddBooleanOptionE("v", "verbose", false, "Enable verbose output", "", nil)
if err != nil {
    log.Fatal(err)
}

port, err := u.AddIntegerOptionE("p", "port", 8080, "Server port", "", nil)
if err != nil {
    log.Fatal(err)
}
```

**Why migrate?**
- Better error handling and debugging
- More testable code
- Aligns with Go best practices
- Prevents silent failures

## Comparison to Alternatives

### vs. Standard `flag` Package
- **More features:** Option groups, colored output, metadata support
- **Better organization:** Group related flags together
- **Richer output:** Automatic formatting with colors and structure
- **Same foundation:** Built on top of `flag`, so familiar and reliable

### vs. `cobra` / `urfave/cli`
- **Lighter weight:** Focused solely on flag parsing, not full CLI frameworks
- **Simpler API:** Less boilerplate, easier to learn
- **Functional options:** Clean, extensible configuration
- **No subcommands:** If you need complex command trees, use `cobra`

### Unique Features
- Dual formatter system (colored + plain-text)
- Functional options pattern throughout
- Build metadata integration (version, commit, date)
- Both panic and error-returning methods

## Testing & Coverage

The library has comprehensive test coverage:

- **Overall Coverage:** 21.1% (main package)
- **Internal Package:** 83.7% coverage
- **Pkg Package:** 100% coverage

[![Coverage](https://codecov.io/gh/bgrewell/usage/graphs/sunburst.svg?token=MP7QKP53BG)](https://codecov.io/gh/bgrewell/usage)

**Run tests locally:**
```bash
go test ./...
go test -cover ./...
```

## Contributing

Contributions are welcome! Here's how to get started:

1. **Fork the repository**
2. **Create a feature branch:** `git checkout -b feature/my-feature`
3. **Make your changes**
4. **Run tests:** `go test ./...`
5. **Format code:** `gofmt -w .`
6. **Lint code:** `golangci-lint run ./...` (if available)
7. **Commit changes:** `git commit -m "Add my feature"`
8. **Push to branch:** `git push origin feature/my-feature`
9. **Open a Pull Request**

**Code Style:**
- Follow standard Go conventions (`gofmt`, `golangci-lint`)
- Add tests for new functionality
- Update documentation for API changes
- Keep commits focused and atomic

**Found a bug?** Please [open an issue](https://github.com/bgrewell/usage/issues)!

## License

This project is open source. Please check the repository for license information.

---

**Questions or feedback?** Open an issue or discussion on [GitHub](https://github.com/bgrewell/usage)!
