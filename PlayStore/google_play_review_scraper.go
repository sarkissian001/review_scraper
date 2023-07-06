package playstore

import (
	"log"
	"strconv"

	"github.com/n0madic/google-play-scraper/pkg/reviews"
)

// GooglePlayReviewScraper represents a Google Play Store review scraper
type GooglePlayReviewScraper struct {
	AppID   string
	Options reviews.Options
}

// NewGooglePlayReviewScraper creates a new GooglePlayReviewScraper instance
func NewGooglePlayReviewScraper(appID string, options reviews.Options) *GooglePlayReviewScraper {
	return &GooglePlayReviewScraper{
		AppID:   appID,
		Options: options,
	}
}

func (scraper *GooglePlayReviewScraper) ScrapeReviews() (reviews.Results, error) {
	r := reviews.New(scraper.AppID, scraper.Options)

	err := r.Run()
	if err != nil {
		return nil, err
	}

	log.Printf("Fetched.... %s reviews for app `%s` ", strconv.Itoa(len(r.Results)), scraper.AppID)

	return r.Results, nil
}
