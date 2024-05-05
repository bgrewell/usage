package internal

type Formatter interface {
	PrintUsage()
	PrintError(err error)
}
