package main

import (
	"flag"
	"log"
	scrapers "review_scraper/Scrapers"
	utils "review_scraper/Utils"
	"strings"
	"time"

	"github.com/n0madic/google-play-scraper/pkg/reviews"
)

func main() {
	// Command line flags
	appID := flag.String("appID", "", "App package ID")
	reviewNum := flag.Int("reviewNum", 0, "Number of reviews to scrape")
	flag.Parse()

	if *appID == "" {
		log.Fatal("Please provide an appID using the -appID flag.")
	}

	// Create a GooglePlayReviewScraper instance
	scraper := scrapers.NewGooglePlayReviewScraper(*appID, reviews.Options{
		Number: *reviewNum,
	})

	// Scrape reviews
	reviewResults, err := scraper.ScrapeReviews()
	if err != nil {
		log.Fatal(err)
	}

	// Log the number of records processed
	log.Printf("Total reviews processed: %d\n", len(reviewResults))

	// Generate the file name with timestamp and modified appID
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	fileName := timestamp + "_" + strings.ReplaceAll(*appID, ".", "_") + ".json"

	// Save reviews to JSON file
	err = utils.OutputToJSON(reviewResults, fileName)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Reviews saved to %s\n", fileName)
}
