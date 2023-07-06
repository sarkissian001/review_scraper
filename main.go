package main

import (
	"flag"
	"log"
	appstore "review_scraper/AppStore"
	playstore "review_scraper/PlayStore"
	utils "review_scraper/Utils"
	"strings"
	"time"

	"github.com/n0madic/google-play-scraper/pkg/reviews"
)

func main() {
	// Command line flags
	source := flag.String("source", "", "Source of the reviews (google or appstore)")
	appName := flag.String("appName", "", "App Name")
	appID := flag.Int("appID", 0, "ID of the App")
	limit := flag.Int("limit", 0, "Number of reviews to scrape")
	country := flag.String("country", "GB", "Optional Review Country | Default is `GB`")

	flag.Parse()

	if *appName == "" {
		log.Fatal("Please provide an appName using the -appName flag.")
	}
	if *source == "" {
		log.Fatal("Please provide a source using the -source flag.")
	}

	var reviewResults interface{}

	var err error

	if *source == "playstore" {
		// Create a GooglePlayReviewScraper instance
		scraper := playstore.NewGooglePlayReviewScraper(*appName, reviews.Options{
			Number: *limit,
		})

		// Scrape reviews
		reviewResults, err = scraper.ScrapeReviews()
		if err != nil {
			log.Fatal(err)
		}

	} else if *source == "appstore" {
		scraper := appstore.NewAppStoreScraper(*country, *appName, *appID, *limit)

		reviewResults, err = scraper.Fetch()
		if err != nil {
			log.Fatal(err)
		}

	} else {
		log.Fatal("Invalid source. Must be either 'playstore' or 'appstore'")
	}

	// Generate the file name with timestamp and modified appID
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	fileName := strings.ReplaceAll(*appName, ".", "_") + "_" + *source + timestamp + "_" + ".json"

	// Save reviews to JSON file
	err = utils.OutputToJSON(reviewResults, fileName)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Reviews saved to %s\n", fileName)
}
