package appstore

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type AppStoreScraper struct {
	_scheme         string
	_landingHost    string
	_requestHost    string
	_landingPath    string
	_requestPath    string
	_userAgents     []string
	country         string
	appName         string
	appID           int
	limit           int
	url             string
	reviews         []Review
	reviewsCount    int
	_interval       float64
	timer           float64
	_fetchedCount   int
	_requestOffset  int
	_requestHeaders http.Header
	_requestParams  map[string]interface{}
}

func NewAppStoreScraper(country, appName string, appID int, limit int) *AppStoreScraper {

	new := &AppStoreScraper{
		_scheme:      "https",
		_landingHost: "apps.apple.com",
		_requestHost: "amp-api.apps.apple.com",
		_landingPath: "{country}/app/{app_name}/id{app_id}",
		_requestPath: "v1/catalog/{country}/apps/{app_id}/reviews",
		_userAgents: []string{
			"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; .NET CLR 1.1.4322)",
			"Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko",
			"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)",
		},
		country:        strings.ToLower(country),
		appName:        regexp.MustCompile(`[\W_]+`).ReplaceAllString(strings.ToLower(appName), "-"),
		appID:          appID,
		limit:          limit,
		reviews:        make([]Review, 0),
		reviewsCount:   0,
		timer:          0,
		_fetchedCount:  0,
		_requestOffset: 0,

		_requestHeaders: http.Header{
			"Accept":       []string{"application/json"},
			"Connection":   []string{"keep-alive"},
			"Content-Type": []string{"application/x-www-form-urlencoded; charset=UTF-8"},
		},

		_requestParams: map[string]interface{}{
			"l":                   "en-GB",
			"offset":              0,
			"limit":               20,
			"platform":            "web",
			"additionalPlatforms": "appletv,ipad,iphone,mac",
		},
	}

	new.url = new._landingURL()
	new._requestHeaders.Set("Origin", new._scheme+"://"+new._requestHost)
	new._requestHeaders.Set("Referer", new.url)
	new._requestHeaders.Set("User-Agent", new._randomUserAgent())

	fmt.Printf("Initialized: %T('%s', '%s', %d)\n", new, new.country, new.appName, new.appID)
	fmt.Printf("Ready to fetch reviews from: %s\n", new.url)

	return new
}

func (b *AppStoreScraper) _landingURL() string {
	landingURL := fmt.Sprintf("%s://%s/%s", b._scheme, b._landingHost, b._landingPath)
	return strings.NewReplacer("{country}", b.country, "{app_name}", b.appName, "{app_id}", strconv.Itoa(b.appID)).Replace(landingURL)
}

func (b *AppStoreScraper) _requestURL() string {
	requestURL := fmt.Sprintf("%s://%s/%s", b._scheme, b._requestHost, b._requestPath)
	return strings.NewReplacer("{country}", b.country, "{app_id}", strconv.Itoa(b.appID)).Replace(requestURL)
}

func (b *AppStoreScraper) _get(url string, headers http.Header, params map[string]interface{}, total int, backoffFactor int, statusForcelist []int) (*http.Response, error) {
	retries := NewRetry(total, backoffFactor, statusForcelist)
	client := &http.Client{Transport: retries}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header = headers

	// Set query parameters
	query := req.URL.Query()
	for key, value := range params {
		query.Set(key, fmt.Sprintf("%v", value))
	}
	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *AppStoreScraper) _token() string {
	resp, err := b._get(b.url, nil, nil, 3, 3, []int{404, 429})
	if err != nil {
		fmt.Printf("Failed to make a request: %s\n", err.Error())
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read response body: %s\n", err.Error())
		return ""
	}

	tags := strings.Split(string(body), "\n")
	for _, tag := range tags {
		if strings.Contains(tag, "<meta") && strings.Contains(tag, "web-experience-app/config/environment") {
			re := regexp.MustCompile(`token%22%3A%22(.+?)%22`)
			tokenMatches := re.FindStringSubmatch(tag)
			if len(tokenMatches) > 1 {
				token := tokenMatches[1]
				log.Printf("found app store token: %.10s...", token)
				return "bearer " + token
			}
		}
	}

	return ""
}

func (b *AppStoreScraper) _randomUserAgent() string {
	rand.Seed(time.Now().UnixNano())
	return b._userAgents[rand.Intn(len(b._userAgents))]
}
