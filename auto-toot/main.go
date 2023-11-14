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

	fileContent := `
---
date: 2023-11-07T10:35:52+11:00
title: "Configuring Kubuntu"
description: "Until recently, Arch Linux has been my daily driver. Here's how I configure my new Kubuntu installation."
tags: [
    "Linux",
    "Kubuntu",
    "Ubuntu",
    "KDE Plasma",
    "Flatpak",
    "Flathub",
    "Discover",
    "Defaults",
    "Editor",
    "Vim",
    "Nano",
    "Shell",
    "Bash",
    "Zsh",
    "OhMyZsh",
  ]
# author: ["Toby Scott", "Other example contributor"]
hidden: false
draft: false
---
`
	// Assume we get the filename from somewhere, for example:
	filePath := "content/posts/hello-world.md"

	hugoPostDetails, err := helpers.ParseHugoPost(filePath, fileContent)
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

			fmt.Printf("Posting to Mastodon for %s", file)

			// Prepare the status message
			status := fmt.Sprintf("New post added: %s", file)

			helpers.SendToot(mastodonURL, accessToken, status)
			if err == nil {
				fmt.Printf("Successfully posted about %s to Mastodon.\n", file)
			}
		}
	}
}
