package main

import (
	"github.com/kgysu/oc-wrapper/fileutils"
	"github.com/kgysu/oc-wrapper/gitutils"
	"os"
)

func main() {
	currentDir, err := fileutils.GetCurrentDir()
	if err != nil {
		panic(err)
	}
	subPathInGit := "/apps"
	gitRepoUrl := "https://github.com/kgysu/oc-wrapper"
	gitRepoTag := "v0.1.0"
	gitRepoBranch := ""

	err = gitutils.LoadFromGitRepo(os.Stdout, currentDir, subPathInGit, gitRepoUrl, gitRepoTag, gitRepoBranch)
	if err != nil {
		panic(err)
	}

}
