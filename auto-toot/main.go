package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	// Execute 'git show --pretty="" --name-only' to get the list of files changed in the last commit
	cmd := exec.Command("git", "show", "--pretty=format:", "--name-only", "--diff-filter=A")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error running git show: %v\n", err)
		return
	}

	// Split the output into lines
	changedFiles := strings.Split(out.String(), "\n")

	// Iterate through the list of changed files and print only the added ones in 'content/posts/'
	fmt.Println("Added files in 'content/posts/':")
	for _, file := range changedFiles {
		// Check if the file is in the 'content/posts/' directory
		if strings.HasPrefix(file, "content/posts/") {
			fmt.Println(file)
		}
	}
}
