package helpers

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"
)

type HugoPost struct {
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	URL         string   `json:"url"`
}

func ParseHugoPost(filePath string, blogUrl string) (HugoPost, error) {

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return HugoPost{}, err
	}
	defer file.Close()

	var fileContent strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileContent.WriteString(scanner.Text() + "\n")
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return HugoPost{}, err
	}

	// Convert the builder to a string for regex processing
	contentString := fileContent.String()

	// Convert the file path to a blog post URL
	postSlug := strings.TrimPrefix(filePath, "../content/posts/")
	postSlug = strings.TrimSuffix(postSlug, ".md")
	postSlug = strings.TrimSuffix(postSlug, "/index") // Remove '/index' if present

	// Use url.QueryEscape to encode spaces and other characters in the post slug
	postSlug = url.QueryEscape(postSlug)

	blogPostURL := fmt.Sprintf(blogUrl+"/posts/%s", postSlug)

	// Use regex to find the description block
	descriptionRegex := regexp.MustCompile(`description: "(.*?)"`)
	descriptionMatch := descriptionRegex.FindStringSubmatch(contentString)

	description := ""
	if len(descriptionMatch) > 1 {
		description = descriptionMatch[1]
	}

	// Use regex to find the tags block (taking new lines into account)
	tagsRegex := regexp.MustCompile(`(?s)tags:\s+\[(.*?)\]`)
	tagsMatch := tagsRegex.FindStringSubmatch(contentString)

	// Process tags
	var tags []string
	if len(tagsMatch) > 1 {
		// Remove the square brackets and split the string by comma
		tagsStr := strings.Trim(tagsMatch[1], "[]")
		tagsStr = strings.Replace(tagsStr, "    \"", "\"", -1) // Remove 4-space indentations
		tagsStr = strings.Replace(tagsStr, "\"", "", -1)       // Remove quotes
		tagsStr = strings.Replace(tagsStr, "\n", "", -1)       // Remove newlines
		tagsStr = strings.Replace(tagsStr, " ", "", -1)        // Remove spaces

		// Remove trailing comma
		if strings.HasSuffix(tagsStr, ",") {
			tagsStr = tagsStr[:len(tagsStr)-1]
		}
		tags = strings.Split(tagsStr, ",")

	}

	post := HugoPost{
		Description: description,
		Tags:        tags,
		URL:         blogPostURL,
	}

	return post, nil
}

func (p HugoPost) GetHashtagString() string {
	var hashtags []string
	for _, tag := range p.Tags {
		hashtags = append(hashtags, fmt.Sprintf("#%s", tag))
	}
	hashtagsStr := strings.Join(hashtags, " ")
	hashtagsStr = strings.TrimSpace(hashtagsStr) // Trim the trailing space
	return hashtagsStr
}
