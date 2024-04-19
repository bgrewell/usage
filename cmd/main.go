package main

import "github.com/bgrewell/usage"

func main() {
	// Create a new sage to handle command line arguments
	sage := usage.NewUsage(
		usage.WithApplicationName("Bowser"),
		usage.WithApplicationVersion("0.0.1"),
		usage.WithApplicationBuild("debug"),
		usage.WithApplicationRevision("rev X"),
		usage.WithApplicationBranch("development"),
		usage.WithApplicationDescription("It's almost a browser but not quite"))

	default_group := sage.AddGroup("Default Options", "These are the default options")

	timeout := default_group.AddIntegerOption("t", "timeout", 10, "Timeout in seconds", "", default_group)
	agent := default_group.AddStringOption("u", "user-agent", "Bowser/0.0.1", "The user agent to use", "")
	request_type := default_group.AddStringOption("r", "request-type", "GET", "The type of request to make", "")
	follow := default_group.AddBooleanOption("f", "follow", false, "Follow Redirects", "not yet implemented")
	url := default_group.AddArgument(1, "url", "The url of the page to retrieve", "Extra")

	sage.AddGroup(default_group)

	parsed := sage.Parse()

	if !parsed {
		sage.PrintUsage()
	}

}
