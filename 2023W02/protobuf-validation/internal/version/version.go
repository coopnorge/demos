package version

import (
	"os"
	"path"
)

var (
	gitDesc       string
	gitCommit     string
	gitRemote     string
	gitCommitDate string
)

// Get returns the version info for this application.
func Get() (info map[string]string, err error) {
	info = map[string]string{
		"executable":      path.Base(os.Args[0]),
		"git_description": gitDesc,
		"git_commit":      gitCommit,
		"git_remote":      gitRemote,
		"git_commit_date": gitCommitDate,
	}
	return info, nil
}
