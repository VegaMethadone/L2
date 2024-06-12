package main

import (
	"fmt"
	"io"
	"manwget/arguments"
	"manwget/functions"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func init() {
	fmt.Println("-r <1> ")
	fmt.Println("EXAMPLE: go run main.go -r 1 google.com")
}

func createDirIfNotExists(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	args, err := arguments.GetArgs()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	pageURL := args.Url[0]

	resp, err := http.Get(pageURL)
	if err != nil {
		fmt.Println("Error fetching the URL:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return
	}

	html := string(body)
	parsedURL, err := url.Parse(pageURL)
	if err != nil {
		fmt.Println("Error parsing the URL:", err)
		return
	}
	domain := strings.TrimPrefix(parsedURL.Hostname(), "www.")

	mainDir := domain
	err = createDirIfNotExists(mainDir)
	if err != nil {
		fmt.Println("Error creating main directory:", err)
		return
	}

	htmlFilePath := filepath.Join(mainDir, "index.html")
	err = os.WriteFile(htmlFilePath, []byte(html), os.ModePerm)
	if err != nil {
		fmt.Println("Error saving HTML file:", err)
		return
	}

	cssDir := filepath.Join(mainDir, "static")
	jsDir := filepath.Join(mainDir, "js")
	err = createDirIfNotExists(cssDir)
	if err != nil {
		fmt.Println("Error creating CSS directory:", err)
		return
	}
	err = createDirIfNotExists(jsDir)
	if err != nil {
		fmt.Println("Error creating JS directory:", err)
		return
	}

	files := functions.ExtractFiles(pageURL, html)
	for _, fileURL := range files {
		fmt.Println("Downloading:", fileURL)

		filename := path.Base(fileURL)
		var savePath string
		if strings.HasSuffix(filename, ".css") {
			savePath = filepath.Join(cssDir, filename)
		} else if strings.HasSuffix(filename, ".js") {
			savePath = filepath.Join(jsDir, filename)
		}

		err := functions.DownloadFile(fileURL, savePath)
		if err != nil {
			fmt.Println("Error downloading the file:", err)
		}
	}
}
