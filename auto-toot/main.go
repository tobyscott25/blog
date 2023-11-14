package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

// Structure to parse the response from Mastodon API
type MastodonStatusResponse struct {
	URL string `json:"url"`
}

func sendToot(mastodonURL string, accessToken string, status string) error {
	data := url.Values{}
	data.Set("status", status)

	// Create a new request
	req, err := http.NewRequest("POST", mastodonURL, strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return err
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error posting to Mastodon: %v\n", err)
		return err
	}
	defer resp.Body.Close()

	// Check the response
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error response from Mastodon: %v\n", resp.Status)
		return err
	}

	// Read response body
	respBody, error := io.ReadAll(resp.Body)
	if error != nil {
		fmt.Println(error)
	}

	// Parse the JSON response
	var mastodonResp MastodonStatusResponse
	if err := json.Unmarshal(respBody, &mastodonResp); err != nil {
		fmt.Printf("Error parsing JSON response: %v\n", err)
		return err
	}

	fmt.Printf("Successfully posted to Mastodon: %s \n", mastodonResp.URL)

	return nil
}

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

	sendToot(mastodonURL, accessToken, "test")

	// Iterate through the list of changed files and post each one to Mastodon
	for _, file := range changedFiles {
		// Check if the file is in the 'content/posts/' directory
		if strings.HasPrefix(file, "content/posts/") {

			fmt.Printf("Posting to Mastodon for %s", file)

			// Prepare the status message
			status := fmt.Sprintf("New post added: %s", file)

			sendToot(mastodonURL, accessToken, status)
			if err != nil {
				fmt.Printf("Successfully posted about %s to Mastodon.\n", file)
			}
		}
	}
}
