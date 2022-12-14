package notionapi_test

import (
	"context"
	"fmt"
	"github.com/conduitio-labs/notionapi"
	"net/http"
	"os"
	"testing"
)

// RoundTripFunc .
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// newTestClient returns *http.Client with Transport replaced to avoid making real calls
func newTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

// newMockedClient returns *http.Client which responds with content from given file
func newMockedClient(t *testing.T, requestMockFile string, statusCode int) *http.Client {
	return newTestClient(func(*http.Request) *http.Response {
		b, err := os.Open(requestMockFile)
		if err != nil {
			t.Fatal(err)
		}

		resp := &http.Response{
			StatusCode: statusCode,
			Body:       b,
			Header:     make(http.Header),
		}
		return resp
	})
}

func TestRateLimitRetryError(t *testing.T) {
	requests := 0
	maxRetries := 7

	c := newTestClient(func(*http.Request) *http.Response {
		requests++
		return &http.Response{
			StatusCode: http.StatusTooManyRequests,
			Header:     http.Header{"Retry-After": []string{"0"}},
		}
	})
	client := notionapi.NewClient(
		"irrelevant_token",
		notionapi.WithHTTPClient(c),
		notionapi.WithMaxRetries(maxRetries),
	)

	_, err := client.Page.Get(context.Background(), "some_block_id")
	if err == nil {
		t.Errorf("Get() error = %v", err)
	}

	wantErr := fmt.Sprintf("max number of retries (%v) reached", maxRetries)
	if err.Error() != wantErr {
		t.Errorf("Get() error = %v, wantErr %s", err, wantErr)
	}

	if requests != maxRetries+1 {
		t.Errorf("expected %v requests but got %v", maxRetries, requests)
	}
}
