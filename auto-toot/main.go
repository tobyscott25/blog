package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func main() {

	// Mastodon API endpoint for posting a status
	// mastodonURL := "https://mas.to/api/v1/statuses"
	// accessToken := os.Getenv("MASTODON_ACCESS_TOKEN")

	// // exec.Command("cd", "../").Run()
	// filePath := "../content/posts/configuring-kubuntu/index.md"

	// hugoPostDetails, err := helpers.ParseHugoPost(filePath)
	// if err != nil {
	// 	fmt.Printf("Error parsing Hugo post: %v\n", err)
	// 	return
	// }

	// hashtagString := hugoPostDetails.GetHashtagString()

	// status := fmt.Sprintf("%s\n\n%s\n\n%s", hugoPostDetails.Description, hugoPostDetails.URL, hashtagString)
	// fmt.Println(status)
	// // SendToot(mastodonURL, accessToken, status)

	// Execute 'git diff --cached --name-only --diff-filter=A' to get the list of added files in the last commit
	cmd := exec.Command("git", "diff", "--cached", "--name-only", "--diff-filter=A")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmdErr := cmd.Run()
	if cmdErr != nil {
		fmt.Printf("Error running git diff: %v\n", cmdErr)
		return
	}

	// Split the output into lines
	changedFiles := strings.Split(out.String(), "\n")

	fmt.Println("iterating through added files")
	fmt.Println(changedFiles)

	// Iterate through the list of changed files and post each one to Mastodon
	for _, file := range changedFiles {
		fmt.Printf("File: %s\n", file)
		// Check if the file is in the 'content/posts/' directory
		if strings.HasPrefix(file, "content/posts/") {

			fmt.Printf("Posting to Mastodon for %s...", file)

			// Prepare the status message
			// status := fmt.Sprintf("New post added: %s", file)

			// helpers.SendToot(mastodonURL, accessToken, status)
			// if err != nil {
			// 	fmt.Printf("Error posting about %s to Mastodon: %v\n", file, err)
			// } else {
			// 	fmt.Printf("Successfully posted about %s to Mastodon.\n", file)
			// }
		}
	}
}
