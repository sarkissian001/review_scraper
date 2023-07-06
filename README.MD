# Review Scraper


This is a Go project that scrapes reviews from the Google Play Store and saves them as a JSON file.

---


## Prerequisites

You need to have Go installed on your machine

---

## Getting Started

To get started with this project, do the following:

---

1. Clone this repository:


```shell
     git clone https://github.com/sarkissian001/review_scraper.git
```


2. Navigate to the review_scraper directory:


```shell
     cd review_scraper
```

 ### Usage


```
  -appID int
        ID of the App
  -appName string
        App Name
  -country GB
        Optional Review Country | Default is GB (default "GB")
  -limit int
        Number of reviews to scrape
  -source string
        Source of the reviews (google or appstore)

```
 
 - Download PlayStore reviews

     ```shell
          go run main.go -appName com.spotify.music -limit 30 -source playstore
     ```

- Download AppStore reviews

     ```shell
     go run main.go -appName spotify -appID 324684580 -limit 30 -source appstore
     ```

### Run Tests

```shell
     go test ./...    
```         