package internal

// Configuration holds all application metadata and option groups.
// This is the central data structure used by formatters to generate usage output.
type Configuration struct {
	ApplicationName        string            // Name of the application
	ApplicationVersion     string            // Version string
	ApplicationBuildDate   string            // Build date/timestamp
	ApplicationCommitHash  string            // Git commit hash
	ApplicationBranch      string            // Git branch name
	ApplicationDescription string            // Application description
	Groups                 map[string]*Group // Option groups keyed by name
}
