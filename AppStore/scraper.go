package appstore

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func (b *AppStoreScraper) Fetch() ([]Review, error) {
	var reviews []Review

	token := b._token()
	if token == "" {
		return nil, fmt.Errorf("failed to fetch token")
	}

	b._requestHeaders.Set("Authorization", token)
	b.url = b._requestURL()
	b._requestParams["limit"] = "20" // Set the initial limit to 20

	totalReviews := 0 // Track the total number of reviews fetched
	offset := 0       // Offset for pagination

	for totalReviews < b.limit {
		b._requestParams["offset"] = strconv.Itoa(offset) // Update the offset parameter

		resp, err := b._get(b.url, b._requestHeaders, b._requestParams, 3, 3, []int{401, 404, 429})
		if err != nil {
			return nil, fmt.Errorf("failed to make a request: %s", err.Error())
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %s", err.Error())
		}

		var data map[string]interface{}
		if err := json.Unmarshal(body, &data); err != nil {
			return reviews, fmt.Errorf("failed to parse response body: %s", err.Error())
		}

		reviewsData, ok := data["data"].([]interface{})
		if !ok {
			return reviews, fmt.Errorf("invalid reviews data format")
		}

		// Extract reviews and append them to the reviews slice
		for _, review := range reviewsData {
			reviewMap, ok := review.(map[string]interface{})
			if !ok {
				return reviews, fmt.Errorf("invalid review format")
			}

			attributes, ok := reviewMap["attributes"].(map[string]interface{})
			if !ok {
				return reviews, fmt.Errorf("invalid attributes format")
			}

			// Extract the review details
			reviewText := attributes["review"].(string)
			rating := attributes["rating"].(float64)
			date := attributes["date"].(string)
			title := attributes["title"].(string)
			isEdited := attributes["isEdited"].(bool)
			userName := attributes["userName"].(string)

			// Create a new Review instance
			reviewObj := Review{
				ReviewText: reviewText,
				Rating:     rating,
				Date:       date,
				Title:      title,
				IsEdited:   isEdited,
				UserName:   userName,
			}

			// Append the review to the reviews slice
			reviews = append(reviews, reviewObj)

			totalReviews++
			if totalReviews >= b.limit {
				return reviews, nil // Stop fetching if the desired number of reviews is reached
			}
		}

		if len(reviewsData) == 0 {
			break // No more reviews to fetch, exit the loop
		}

		offset += len(reviewsData) // Increment the offset for the next page
		log.Printf("Fetched.... %s reviews for app `%s` ", strconv.Itoa(len(reviews)), b.appName)
		time.Sleep(time.Duration(rand.Intn(2000)+1000) * time.Millisecond) // Sleep to avoid rate limiting
	}

	return reviews, nil
}
