package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	println("THE ROYAL TOOT!")
	// Send a toot to Mastodon
	println(os.Getenv("MASTODON_ACCESS_TOKEN"))

	exec.Command(
		"cd", "../",
	)

	// Get the latest commit hash
	commitHash, err := exec.Command("git", "rev-parse", "HEAD").Output()
	if err != nil {
		fmt.Println("Error getting latest commit hash:", err)
		return
	}

	// Get the list of files that have been created in the latest commit
	cmd := exec.Command("git", "diff", "--diff-filter=A", "--name-only", strings.TrimSpace(string(commitHash)), "HEAD~1")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error running git diff:", err)
		return
	}

	// Filter the list for files in content/posts/*
	createdFiles := []string{}
	for _, file := range strings.Split(out.String(), "\n") {
		if strings.HasPrefix(file, "content/posts/") {
			createdFiles = append(createdFiles, file)
		}
	}

	// Print the list of newly created files
	fmt.Println("Newly created files in 'content/posts/':")
	for _, file := range createdFiles {
		fmt.Println(file)
	}

}
