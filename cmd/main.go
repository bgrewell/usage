package main

import "github.com/bgrewell/usage"

var (
	buildDate  string
	commitHash string
	branch     string
)

func main() {
	// Create a new sage to handle command line arguments
	sage := usage.NewUsage(
		usage.WithApplicationName("bowser"),
		usage.WithApplicationVersion("0.0.1"),
		usage.WithApplicationBuildDate(buildDate),
		usage.WithApplicationCommitHash(commitHash),
		usage.WithApplicationBranch(branch),
		usage.WithApplicationDescription("It's almost a browser but not quite"))

	// Add some options
	output_file := sage.AddStringOption("o", "output", "", "Output filename", "", nil)

	// Add an options group
	request_group := sage.AddGroup("Request Options", "Options related http request")
	agent := sage.AddStringOption("u", "user-agent", "Bowser/0.0.1", "The user agent to use", "", request_group)
	request_type := sage.AddStringOption("r", "request-type", "GET", "The type of request to make", "", request_group)
	follow := sage.AddBooleanOption("f", "follow", false, "Follow Redirects", "not yet implemented", request_group)
	timeout := sage.AddIntegerOption("t", "timeout", 10, "Timeout in seconds", "", request_group)

	url := sage.AddArgument(1, "url", "The url of the page to retrieve", "Extra")

	parsed := sage.Parse()

	if !parsed {
		sage.PrintUsage()
	}

	_ = output_file
	_ = agent
	_ = request_type
	_ = follow
	_ = timeout
	_ = url

}
