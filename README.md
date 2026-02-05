# Usage

[![golangci-lint](https://github.com/bgrewell/usage/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/bgrewell/usage/actions/workflows/golangci-lint.yml)
[![codecov](https://codecov.io/gh/bgrewell/usage/graph/badge.svg?token=MP7QKP53BG)](https://codecov.io/gh/bgrewell/usage)

## About

Streamline your Go applications with our powerful CLI management tool. Designed for developers, by developers, this
project simplifies command-line argument parsing, enhances error handling, and standardizes usage output. With 'usage',
setting up and parsing flags becomes effortless, allowing you to focus more on building robust applications. Optimize
your command-line interface experience and reduce boilerplate with our intuitive and flexible API. Ideal for projects
of all sizes, 'usage' ensures clarity and efficiency from development to deployment. Get started today to make your
CLI tools more user-friendly and maintainable!

## Features

- **Flexible Option Management** - Support for boolean, integer, float, and string options with short and long flags
- **Option Groups** - Organize related options into logical groups with custom priorities
- **Custom Formatters** - Built-in color and standard formatters with support for custom implementations
- **Error Handling** - Error-returning methods with optional exit control
- **Non-Exiting Mode** - Safe for library embedding and long-running applications
- **Positional Arguments** - Support for required and optional positional arguments
- **Automatic Usage Generation** - Beautiful, formatted help output with application metadata
- **Build Information** - Display version, build date, commit hash, and branch information
- **Extensive Testing** - Comprehensive test coverage for reliability

## Coverage

![image](https://codecov.io/gh/bgrewell/usage/graphs/sunburst.svg?token=MP7QKP53BG)

## Quick Start

```go
package main

import (
    "github.com/bgrewell/usage"
)

func main() {
    // Create usage instance
    u := usage.NewUsage(
        usage.WithApplicationName("myapp"),
        usage.WithApplicationVersion("1.0.0"),
        usage.WithApplicationDescription("My awesome CLI application"),
    )

    // Add options
    verbose, _ := u.AddBooleanOptionE("v", "verbose", false, "Enable verbose output", "", nil)
    port, _ := u.AddIntegerOptionE("p", "port", 8080, "Server port", "", nil)

    // Parse command line
    u.Parse()

    // Use the parsed values
    if *verbose {
        println("Verbose mode enabled")
    }
    println("Server will run on port:", *port)
}
```

> **Note:** For library usage where you don't want the program to exit, use `WithExitOnHelp(false)` and `WithExitOnError(false)`.

## Example Output

```text
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

## Exit Control

The library provides fine-grained control over program termination behavior, making it safe for use in libraries, long-running applications, and testing scenarios.

### For CLI Applications (Default Behavior)

By default, the library exits on help and error, maintaining traditional CLI behavior:

```go
// Default behavior - exits on help/error
u := usage.NewUsage(
    usage.WithApplicationName("mycli"),
)

// This exits with code 0
u.PrintUsage()

// This exits with code 1
if err := validate(); err != nil {
    u.PrintError(err)
}
```

### For Library Usage (Non-Exiting)

Disable exit behavior when embedding the library or in long-running applications:

```go
// Disable exit behavior for library embedding
u := usage.NewUsage(
    usage.WithApplicationName("mylib"),
    usage.WithExitOnHelp(false),
    usage.WithExitOnError(false),
)

// These no longer exit - they just print
u.PrintUsage()                    // Prints help, returns normally
u.PrintError(errors.New("fail"))  // Prints error, returns normally

// Or use explicit non-exiting methods
u.PrintUsageWithoutExit()
u.PrintErrorWithoutExit(err)

// Or get formatted string for custom handling
helpText := u.FormatUsage()
return fmt.Errorf("invalid arguments:\n%s", helpText)
```

### For Testing

Testing code that uses the usage library is now straightforward:

```go
func TestMyApp(t *testing.T) {
    u := usage.NewUsage(
        usage.WithApplicationName("test"),
        usage.WithExitOnHelp(false),
        usage.WithExitOnError(false),
    )

    // Add options and test
    verbose, _ := u.AddBooleanOptionE("v", "verbose", false, "Verbose", "", nil)

    // Tests won't exit
    helpText := u.FormatUsage()
    assert.Contains(t, helpText, "test")
}
```

## Error Handling

The library provides two styles of methods for adding options:

### Error-Returning Methods (Recommended)

Use the `*E` variants for proper error handling:

```go
verbose, err := u.AddBooleanOptionE("v", "verbose", false, "Enable verbose output", "", nil)
if err != nil {
    return fmt.Errorf("failed to add option: %w", err)
}

port, err := u.AddIntegerOptionE("p", "port", 8080, "Server port", "", nil)
if err != nil {
    return fmt.Errorf("failed to add option: %w", err)
}
```

### Panic Methods (Deprecated)

The original methods without the `E` suffix are deprecated but still supported:

```go
// Deprecated: Will panic on error
verbose := u.AddBooleanOption("v", "verbose", false, "Enable verbose output", "", nil)
```

**Migration:**
```go
// Old (deprecated):
verbose := u.AddBooleanOption("v", "verbose", false, "Enable verbose output", "", nil)

// New (recommended):
verbose, err := u.AddBooleanOptionE("v", "verbose", false, "Enable verbose output", "", nil)
if err != nil {
    return err
}
```

## Option Groups

Organize related options into logical groups:

```go
u := usage.NewUsage(usage.WithApplicationName("myapp"))

// Create a custom group
serverGroup := u.AddGroup(1, "Server Options", "Options related to server configuration")

// Add options to the group
host, _ := u.AddStringOptionE("h", "host", "localhost", "Server host", "", serverGroup)
port, _ := u.AddIntegerOptionE("p", "port", 8080, "Server port", "", serverGroup)

// Options without a group go to the default group
verbose, _ := u.AddBooleanOptionE("v", "verbose", false, "Enable verbose output", "", nil)
```

## API Reference

### Configuration Options

- `WithApplicationName(name string)` - Set the application name
- `WithApplicationVersion(version string)` - Set the version
- `WithApplicationBuildDate(date string)` - Set the build date
- `WithApplicationCommitHash(hash string)` - Set the git commit hash
- `WithApplicationBranch(branch string)` - Set the git branch
- `WithApplicationDescription(desc string)` - Set the description
- `WithFormatter(formatter internal.Formatter)` - Use a custom formatter

**Exit Control:**
- `WithExitOnHelp(exit bool)` - Control whether PrintUsage exits (default: true)
- `WithExitOnError(exit bool)` - Control whether PrintError exits (default: true)

### Adding Options (Recommended)

- `AddBooleanOptionE(short, long string, defaultValue bool, description, extra string, group *internal.Group) (*bool, error)`
- `AddIntegerOptionE(short, long string, defaultValue int, description, extra string, group *internal.Group) (*int, error)`
- `AddFloatOptionE(short, long string, defaultValue float64, description, extra string, group *internal.Group) (*float64, error)`
- `AddStringOptionE(short, long string, defaultValue, description, extra string, group *internal.Group) (*string, error)`

### Adding Options (Deprecated)

- `AddBooleanOption(...)` - Deprecated: Use AddBooleanOptionE
- `AddIntegerOption(...)` - Deprecated: Use AddIntegerOptionE
- `AddFloatOption(...)` - Deprecated: Use AddFloatOptionE
- `AddStringOption(...)` - Deprecated: Use AddStringOptionE

### Groups

- `AddGroup(priority int, name, description string) *internal.Group` - Create a new option group

### Arguments

- `AddArgument(position int, name, description, extra string) *string` - Add a positional argument

### Parsing and Output

- `Parse() bool` - Parse command-line arguments
- `PrintUsage()` - Print usage and exit (respects WithExitOnHelp)
- `PrintError(err error)` - Print error and usage, then exit (respects WithExitOnError)

**Non-Exiting Methods:**
- `FormatUsage() string` - Get formatted usage string without printing
- `PrintUsageWithoutExit()` - Print usage without calling os.Exit
- `PrintErrorWithoutExit(err error)` - Print error without calling os.Exit

### Accessors

- `ApplicationName() string`
- `ApplicationVersion() string`
- `ApplicationBuildDate() string`
- `ApplicationCommitHash() string`
- `ApplicationBranch() string`
- `ApplicationDescription() string`

## Migration Guide

### Migrating from Deprecated Methods

The non-error-returning methods are deprecated. Migrate to the `*E` variants:

**Before:**
```go
verbose := u.AddBooleanOption("v", "verbose", false, "Enable verbose output", "", nil)
port := u.AddIntegerOption("p", "port", 8080, "Server port", "", nil)
```

**After:**
```go
verbose, err := u.AddBooleanOptionE("v", "verbose", false, "Enable verbose output", "", nil)
if err != nil {
    return fmt.Errorf("failed to add verbose option: %w", err)
}

port, err := u.AddIntegerOptionE("p", "port", 8080, "Server port", "", nil)
if err != nil {
    return fmt.Errorf("failed to add port option: %w", err)
}
```

### Migrating to Non-Exiting Behavior

If you're embedding the usage library in another library or long-running application:

**Before (would exit the program):**
```go
u := usage.NewUsage(usage.WithApplicationName("mylib"))
if needHelp {
    u.PrintUsage()  // Exits with code 0 - crashes host app!
}
```

**After (safe for library usage):**
```go
u := usage.NewUsage(
    usage.WithApplicationName("mylib"),
    usage.WithExitOnHelp(false),
    usage.WithExitOnError(false),
)

if needHelp {
    u.PrintUsageWithoutExit()
    return errors.New("help requested")
}

// Or get the formatted string
helpText := u.FormatUsage()
return fmt.Errorf("invalid args:\n%s", helpText)
```

**For Testing:**
```go
func TestMyApp(t *testing.T) {
    u := usage.NewUsage(
        usage.WithApplicationName("test"),
        usage.WithExitOnHelp(false),
        usage.WithExitOnError(false),
    )

    // Now tests won't exit
    helpText := u.FormatUsage()
    assert.Contains(t, helpText, "test")
}
```

## Installation

```bash
go get github.com/bgrewell/usage
```

## License

This project is licensed under the terms specified in the repository.

## Coming Soon

1. Header above columns (flags, default, description)
2. Columnization across groups
3. Arguments section shouldn't be printed if there are no arguments