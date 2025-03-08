package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRateLimiter(t *testing.T) {
	// ping end point
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	rr := httptest.NewRecorder()
	handler := rateLimiter(http.HandlerFunc(apiHandler))
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status 200 OK, got %v", status)
	}
	expected := `{"status":"Success","body":"Hi, Your Server is Up and Running"}`
	actual := strings.TrimSpace(rr.Body.String())
	if actual != expected {
		t.Errorf("expected body %v, got %v", expected, actual)
	}
}

// Simulate a request that will exceed the rate limit
func TestRateLimitExceeded(t *testing.T) {

	for i := 0; i < 6; i++ { // Send 6 requests to trigger the limit after 4
		req, err := http.NewRequest("GET", "/ping", nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}
		rr := httptest.NewRecorder()
		handler := rateLimiter(http.HandlerFunc(apiHandler))
		handler.ServeHTTP(rr, req)
		if i < 4 && rr.Code != http.StatusOK {
			t.Errorf("expected status 200 OK, got %v", rr.Code)
		}
		if i >= 4 && rr.Code != http.StatusTooManyRequests {
			t.Errorf("expected status 429 Too Many Requests, got %v", rr.Code)
		}
	}
}
