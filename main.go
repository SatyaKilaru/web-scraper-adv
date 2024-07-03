package main

import (
	"web-scraper-adv/config"
	"web-scraper-adv/scheduler"
	"web-scraper-adv/scraper"
	"web-scraper-adv/utils"
)

func main() {
	config, err := config.LoadConfig("config/config.json")
	if err != nil {
		panic(err)
	}
	logger := utils.NewLogger()
	scraper := scraper.NewScraper(config, logger)
	scheduler.RunScheduler(scraper, logger)
}
