package main

import (
	"fmt"
	"log"

	"github.com/bgrewell/usage"
)

var (
	version    string
	buildDate  string
	commitHash string
	branch     string
)

func main() {
	// Create a new sage to handle command line arguments
	sage := usage.NewUsage(
		usage.WithApplicationName("bowser"),
		usage.WithApplicationVersion(version),
		usage.WithApplicationBuildDate(buildDate),
		usage.WithApplicationCommitHash(commitHash),
		usage.WithApplicationBranch(branch),
		usage.WithApplicationDescription("It's almost a browser but not quite. Instead it's just a example of how to use the 'usage' package. It is designed to show an example of how to use the package and a sample of what the output would look like."))

	// Add some options using the new error-returning methods
	output_file, err := sage.AddStringOptionE("o", "output", "", "Output filename", "", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Add an options group
	request_group := sage.AddGroup(1, "Request Options", "Options related http request")

	// Use the new error-returning methods with proper error handling
	agent, err := sage.AddStringOptionE("u", "user-agent", "Bowser/0.0.1", "The user agent to use", "", request_group)
	if err != nil {
		log.Fatal(err)
	}

	request_type, err := sage.AddStringOptionE("r", "request-type", "GET", "The type of request to make", "", request_group)
	if err != nil {
		log.Fatal(err)
	}

	follow, err := sage.AddBooleanOptionE("f", "follow", false, "Follow Redirects", "not yet implemented", request_group)
	if err != nil {
		log.Fatal(err)
	}

	timeout, err := sage.AddIntegerOptionE("t", "timeout", 10, "Timeout in seconds", "", request_group)
	if err != nil {
		log.Fatal(err)
	}

	url := sage.AddArgument(1, "url", "The url of the page to retrieve", "Extra")

	parsed := sage.Parse()

	if !parsed {
		sage.PrintUsage()
	}

	fmt.Println("\n\nParsed Values:")
	fmt.Printf("  Output File: %s\n", *output_file)
	fmt.Printf("  User Agent: %s\n", *agent)
	fmt.Printf("  Request Type: %s\n", *request_type)
	fmt.Printf("  Follow Redirects: %t\n", *follow)
	fmt.Printf("  Timeout: %d\n", *timeout)
	fmt.Printf("  URL: %s\n", *url)
}
