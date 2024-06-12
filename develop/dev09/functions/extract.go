package functions

import (
	"fmt"
	"net/url"
	"regexp"
)

// extractFiles extracts all JS and CSS file URLs from the given HTML content.
func ExtractFiles(baseURL, html string) []string {
	var files []string
	// Regex for finding JS and CSS files
	re := regexp.MustCompile(`<(script|link)[^>]+(src|href)="([^"]+\.(js|css))"`)
	matches := re.FindAllStringSubmatch(html, -1)

	base, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println("Invalid base URL:", err)
		return files
	}

	for _, match := range matches {
		fileURL := match[3]
		u, err := url.Parse(fileURL)
		if err != nil {
			fmt.Println("Invalid file URL:", fileURL, err)
			continue
		}
		if !u.IsAbs() {
			u = base.ResolveReference(u)
		}
		files = append(files, u.String())
	}

	return files
}
