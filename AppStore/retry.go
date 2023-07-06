package appstore

import (
	"fmt"
	"net/http"
	"time"
)

type Retry struct {
	total           int
	backoffFactor   int
	statusForcelist []int
}

func NewRetry(total, backoffFactor int, statusForcelist []int) *Retry {
	return &Retry{
		total:           total,
		backoffFactor:   backoffFactor,
		statusForcelist: statusForcelist,
	}
}

func (r *Retry) RoundTrip(req *http.Request) (*http.Response, error) {
	var resp *http.Response
	var err error

	for i := 0; i <= r.total; i++ {
		if i > 0 {
			backoff := time.Duration(i*r.backoffFactor) * time.Second
			time.Sleep(backoff)
		}

		resp, err = http.DefaultTransport.RoundTrip(req)

		if err != nil {
			continue
		}

		if r.shouldRetry(resp.StatusCode) {
			continue
		}

		return resp, nil
	}

	return nil, fmt.Errorf("maximum retries exceeded")
}

func (r *Retry) shouldRetry(statusCode int) bool {
	for _, code := range r.statusForcelist {
		if statusCode == code {
			return true
		}
	}
	return false
}
