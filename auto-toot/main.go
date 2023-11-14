package main

import (
	"bufio"
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

	// Use a scanner to read the command output line by line
	scanner := bufio.NewScanner(&out)
	for scanner.Scan() {
		file := scanner.Text()
		// Skip empty lines
		if strings.TrimSpace(file) != "" {
			fmt.Println(file)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading command output: %v\n", err)
	}
}
