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

```shell
     go run main.go -appID <com.hellofresh.androidapp> -reviewNum <NUMBER_OF_REVIEWS_TO_GET>
```

### Run Tests

```shell
     go test ./...    
```         