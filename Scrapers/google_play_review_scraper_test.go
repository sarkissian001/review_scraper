package scrapers

import (
	"testing"

	"github.com/n0madic/google-play-scraper/pkg/reviews"
)

func TestNewGooglePlayReviewScraper(t *testing.T) {
	appID := "com.example.app"
	options := reviews.Options{}
	scraper := NewGooglePlayReviewScraper(appID, options)

	if scraper == nil {
		t.Error("Expected GooglePlayReviewScraper instance, got nil")
	}
	if scraper.AppID != appID {
		t.Errorf("Expected AppID to be %s, but got %s", appID, scraper.AppID)
	}
	if scraper.Options != options {
		t.Error("Expected Options to be equal to input options")
	}
}

func TestScrapeReviews(t *testing.T) {
	// This is a simple test and doesn't handle cases where the Run function has side effects or depends on external resources
	appID := "com.example.app"
	options := reviews.Options{}
	scraper := NewGooglePlayReviewScraper(appID, options)
	_, err := scraper.ScrapeReviews()
	if err != nil {
		t.Errorf("Expected nil, got error: %s", err.Error())
	}
}
