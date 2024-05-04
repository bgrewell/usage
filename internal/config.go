package internal

type Configuration struct {
	ApplicationName        string
	ApplicationVersion     string
	ApplicationBuildDate   string
	ApplicationCommitHash  string
	ApplicationBranch      string
	ApplicationDescription string
	Groups                 map[string]*Group
}
