package gitutils

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/kgysu/oc-wrapper/fileutils"
	"io"
	"os"
	"strings"
)

const gitTemp = "/gitTemp"

func LoadFromGitRepo(w io.Writer, toDir, gitSubPath, gitRepoUrl, tagToClone, branchToClone string) error {
	tempDir := toDir + gitTemp
	fmt.Fprintf(w, "cloning from=[%s] to=[%s] \n", gitRepoUrl, tempDir)
	defer os.RemoveAll(tempDir) // clean up

	referenceToClone := plumbing.NewBranchReferenceName(branchToClone)
	if tagToClone != "" {
		referenceToClone = plumbing.NewTagReferenceName(tagToClone)
	}

	// Clones the repository into the given dir, just as a normal git clone does
	cloned, err := git.PlainClone(tempDir, false, &git.CloneOptions{
		URL:           gitRepoUrl,
		Progress:      w,
		ReferenceName: referenceToClone,
	})
	if err != nil {
		return err
	}

	// ... retrieves the branch pointed by HEAD
	ref, err := cloned.Head()
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "cloned Head=[%v] \n", ref.Hash())

	err = moveFilesFromTempToParent(w, tempDir+gitSubPath)
	if err != nil {
		return err
	}
	os.RemoveAll(tempDir)
	return nil
}

func moveFilesFromTempToParent(w io.Writer, tempDir string) error {
	files, err := fileutils.FilePathWalkDir(tempDir)
	if err != nil {
		return err
	}
	for _, file := range files {
		newPath := strings.Replace(file, gitTemp, "", 1)
		err := os.Rename(file, newPath)
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "success=[%s]\n", newPath)
	}
	return nil
}
