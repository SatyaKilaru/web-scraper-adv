package config

import (
	"fmt"
	"os"
	"strconv"
)

type WebsiteConfig struct {
	URL        string `json:"url"`
	TitleXPath string `json:"titleXPath"`
	BodyXPath  string `json:"bodyXPath"`
}

type Config struct {
	Websites []WebsiteConfig `json:"websites"`
	OutputDir string         `json:"outputDir"`
}

func LoadConfig(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config

	for {
		fmt.Print("Enter the number of websites to scrape: ")
		var numWebsitesStr string
		_, err = fmt.Scanln(&numWebsitesStr)
		if err != nil {
			return nil, err
		}

		numWebsites, err := strconv.Atoi(numWebsitesStr)
		if err == nil && numWebsites > 0 {
			config.Websites = make([]WebsiteConfig, numWebsites)
			break
		}
		fmt.Println("Invalid input. Please enter a positive integer.")
	}

	for i := 0; i < len(config.Websites); i++ {
		fmt.Printf("Enter URL %d: ", i+1)
		var url string
		_, err := fmt.Scanln(&url)
		if err != nil {
			return nil, err
		}
		config.Websites[i].URL = url

		fmt.Printf("Enter Title XPath %d: ", i+1)
		var titleXPath string
		_, err = fmt.Scanln(&titleXPath)
		if err != nil {
			return nil, err
		}
		config.Websites[i].TitleXPath = titleXPath

		fmt.Printf("Enter Body XPath %d: ", i+1)
		var bodyXPath string
		_, err = fmt.Scanln(&bodyXPath)
		if err != nil {
			return nil, err
		}
		config.Websites[i].BodyXPath = bodyXPath
	}

	fmt.Print("Enter the output directory: ")
	_, err = fmt.Scanln(&config.OutputDir)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
