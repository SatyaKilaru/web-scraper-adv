package scheduler

import (
	"time"

	"web-scraper-adv/scraper"
	"web-scraper-adv/utils"
)

func RunScheduler(scraper *scraper.Scraper, logger *utils.Logger) {
	logger.Info("Starting scheduler...")

	ticker := time.NewTicker(1 * time.Hour)  
	for range ticker.C {
		logger.Info("Running scraper...")
		scraper.Scrape()
	}
}
