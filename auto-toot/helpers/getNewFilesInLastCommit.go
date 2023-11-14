package helpers

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// Returns a slice of strings representing the file paths of newly created files in the last commit.
func GetNewFilesInLastCommit() ([]string, error) {
	// Use 'git diff' to get a list of added files in the last commit
	cmd := exec.Command("git", "diff", "--diff-filter=A", "--name-only", "HEAD~1", "HEAD")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("error running git diff: %v", err)
	}

	// Split the output by new lines to get individual file paths
	files := strings.Split(strings.TrimSpace(out.String()), "\n")
	return files, nil
}
