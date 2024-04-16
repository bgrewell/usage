package main

import "github.com/bgrewell/usage"

func main() {
	// Create a new sage to handle command line arguments
	sage := usage.NewUseSage(
		usage.WithApplicationName("Bowser"),
		usage.WithApplicationVersion("0.0.1"),
		usage.WithApplicationBuild("debug"),
		usage.WithApplicationRevision("rev X"),
		usage.WithApplicationBranch("development"),
		usage.WithApplicationDescription("It's almost a browser but not quite"))
	sage.Usage()
}
