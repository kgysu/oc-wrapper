package gitutils

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/kgysu/oc-wrapper/fileutils"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const gitTemp = "gitTemp"

func LoadFromGitRepo(w io.Writer, toDir, gitSubPath, gitRepoUrl, tagToClone, branchToClone string) error {
	err := fileutils.CreateIfNotExists(toDir + filepath.FromSlash("/"+gitSubPath))
	if err != nil {
		return err
	}

	tempDir := toDir + filepath.FromSlash("/") + gitTemp
	fmt.Fprintf(w, "cloning from=[%s] to=[%s] \n", gitRepoUrl, tempDir)
	defer os.RemoveAll(filepath.FromSlash(tempDir)) // clean up

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
	os.RemoveAll(filepath.FromSlash(tempDir)) // clean up
	return nil
}

func moveFilesFromTempToParent(w io.Writer, tempDir string) error {
	files, err := fileutils.FilePathWalkDir(tempDir)
	if err != nil {
		return err
	}
	for _, file := range files {
		newFilePath := strings.Replace(file, filepath.FromSlash("/")+gitTemp, "", 1)
		parentDir := filepath.Dir(newFilePath)
		err := fileutils.CreateIfNotExists(parentDir)
		if err != nil {
			return err
		}
		err = os.Rename(file, newFilePath)
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "written file to: [%s]\n", newFilePath)
	}
	return nil
}
