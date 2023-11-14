package main

import (
	"auto-toot/helpers"
	"fmt"
	"os"
	"strings"
)

func main() {

	mastodonUrl := "https://mas.to"
	blogUrl := "https://www.tobyscott.dev"
	accessToken := os.Getenv("MASTODON_ACCESS_TOKEN")

	newFiles, err := helpers.GetNewFilesInLastCommit()
	if err != nil {
		fmt.Printf("Error getting new files: %v\n", err)
		return
	}

	for _, file := range newFiles {

		if strings.HasPrefix(file, "content/posts/") {

			filePath := "../" + file

			hugoPostDetails, err := helpers.ParseHugoPost(filePath, blogUrl)
			if err != nil {
				fmt.Printf("Error parsing Hugo post: %v\n", err)
				return
			}

			hashtagString := hugoPostDetails.GetHashtagString()
			status := fmt.Sprintf("%s\n\n%s\n\n%s", hugoPostDetails.Description, hugoPostDetails.URL, hashtagString)

			helpers.SendToot(mastodonUrl, accessToken, status)
			if err != nil {
				fmt.Printf("Error posting about %s to Mastodon: %v\n", file, err)
			} else {
				fmt.Printf("Successfully posted about %s to Mastodon.\n", file)
			}
		}
	}
}
