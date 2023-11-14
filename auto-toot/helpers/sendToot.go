package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Structure to parse the response from Mastodon API
type MastodonStatusResponse struct {
	URL string `json:"url"`
}

func SendToot(mastodonURL string, accessToken string, status string) error {
	data := url.Values{}
	data.Set("status", status)

	// Create a new request
	req, err := http.NewRequest("POST", mastodonURL+"/api/v1/statuses", strings.NewReader(data.Encode()))
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
