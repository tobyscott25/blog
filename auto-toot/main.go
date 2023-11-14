package main

import (
	"auto-toot/helpers"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Execute 'git diff --cached --name-only --diff-filter=A' to get the list of added files in the last commit
	cmd := exec.Command("git", "diff", "--cached", "--name-only", "--diff-filter=A")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error running git diff: %v\n", err)
		return
	}

	// Split the output into lines
	changedFiles := strings.Split(out.String(), "\n")

	// Mastodon API endpoint for posting a status
	mastodonURL := "https://mas.to/api/v1/statuses"
	accessToken := os.Getenv("MASTODON_ACCESS_TOKEN")

	filePath := "content/posts/configuring-kubuntu/index.md"

	hugoPostDetails, err := helpers.ParseHugoPost(filePath)
	if err != nil {
		fmt.Printf("Error parsing Hugo post: %v\n", err)
		return
	}

	hashtagString := hugoPostDetails.GetHashtagString()

	// Combine everything into the Mastodon status format
	status := fmt.Sprintf("%s\n\n%s\n\n%s", hugoPostDetails.Description, hugoPostDetails.URL, hashtagString)

	fmt.Println(status)
	// SendToot(mastodonURL, accessToken, status)

	// Iterate through the list of changed files and post each one to Mastodon
	for _, file := range changedFiles {
		// Check if the file is in the 'content/posts/' directory
		if strings.HasPrefix(file, "content/posts/") {

			fmt.Printf("Posting to Mastodon for %s...", file)

			// Prepare the status message
			status := fmt.Sprintf("New post added: %s", file)

			helpers.SendToot(mastodonURL, accessToken, status)
			if err != nil {
				fmt.Printf("Error posting about %s to Mastodon: %v\n", file, err)
			} else {
				fmt.Printf("Successfully posted about %s to Mastodon.\n", file)
			}
		}
	}
}
