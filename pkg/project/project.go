package project

var (
	buildTimestamp string
	gitSHA         string
	version        = "0.2.0"
)

func BuildTimestamp() string {
	return buildTimestamp
}

func GitSHA() string {
	return gitSHA
}

func Version() string {
	return version
}