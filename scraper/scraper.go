package scraper

import (
	"web-scraper-adv/config"
	"web-scraper-adv/utils"
)

type Scraper struct {
	config *config.Config
	logger *utils.Logger
}

func NewScraper(config *config.Config, logger *utils.Logger) *Scraper {
	return &Scraper{
		config: config,
		logger: logger,
	}
}

func (s *Scraper) Scrape() {
	for _, website := range s.config.Websites {
		s.scrapeWebsite(website)
	}
}

func (s *Scraper) scrapeWebsite(website config.WebsiteConfig) {
	
	doc, err := fetchWebPage(website.URL)
	if err != nil {
		s.logger.Error("Failed to fetch web page: %v", err)
		return
	}

	
	title := extractTitle(doc, website.TitleXPath)
	body := extractBody(doc, website.BodyXPath)

	
	saveData(s.config.OutputDir, website.URL, title, body)
}
