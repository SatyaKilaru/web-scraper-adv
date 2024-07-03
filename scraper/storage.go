package scraper

import (
	"fmt"
	"os"
	"strings"
)

func saveData(outputDir, url, title, body string) {
	
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err = os.MkdirAll(outputDir, 0755)
		if err != nil {
			fmt.Printf("Failed to create output directory: %v\n", err)
			return
		}
	}

	
	filename := fmt.Sprintf("%s/%s.txt", outputDir, url2filename(url))
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Failed to create file: %v\n", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("URL: %s\nTitle: %s\nBody: %s\n", url, title, body))
	if err != nil {
		fmt.Printf("Failed to write to file: %v\n", err)
	}
}

func url2filename(url string) string {
	return strings.ReplaceAll(url, "https://", "")
}
