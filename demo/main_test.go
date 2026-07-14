package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
)

func resetCounter() {
	mu.Lock()
	defer mu.Unlock()
	counter = 0
}

func TestLandingPageHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	landing_page_handler(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", res.StatusCode)
	}

	body := rec.Body.String()

	tests := []string{
		"GoHTML",
		"Build HTML with Go — not templates.",
		"Server-rendered counter",
		"Increment",
		"theme-toggle",
		"counter-value",
		"htmlbuilder",
	}

	for _, expected := range tests {
		if !strings.Contains(body, expected) {
			t.Errorf("expected response to contain %q", expected)
		}
	}
}

func TestCounterFragment(t *testing.T) {
	node := counterFragment(42)

	html := node.String()

	tests := []string{
		`id="counter-value"`,
		"42",
		"font-weight",
		"1.4rem",
		"var(--accent)",
	}

	for _, expected := range tests {
		if !strings.Contains(html, expected) {
			t.Errorf("expected fragment to contain %q, got:\n%s", expected, html)
		}
	}
}

func TestIncrementHandler(t *testing.T) {
	resetCounter()

	req := httptest.NewRequest(http.MethodPost, "/api/counter/increment", nil)
	rec := httptest.NewRecorder()

	increment_handler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	body := rec.Body.String()

	if !strings.Contains(body, "1") {
		t.Fatalf("expected counter value 1, got %s", body)
	}

	contentType := rec.Header().Get("Content-Type")
	if contentType != "text/html" {
		t.Errorf("expected content type text/html, got %q", contentType)
	}
}

func TestIncrementHandlerMultipleCalls(t *testing.T) {
	resetCounter()

	for i := 1; i <= 5; i++ {
		req := httptest.NewRequest(http.MethodPost, "/api/counter/increment", nil)
		rec := httptest.NewRecorder()

		increment_handler(rec, req)

		if !strings.Contains(rec.Body.String(), string(rune('0'+i))) {
			t.Errorf(
				"expected counter to contain %d, got %s",
				i,
				rec.Body.String(),
			)
		}
	}
}

func TestIncrementHandlerConcurrency(t *testing.T) {
	resetCounter()

	var wg sync.WaitGroup

	requests := 100

	for i := 0; i < requests; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			req := httptest.NewRequest(
				http.MethodPost,
				"/api/counter/increment",
				nil,
			)

			rec := httptest.NewRecorder()

			increment_handler(rec, req)
		}()
	}

	wg.Wait()

	mu.Lock()
	defer mu.Unlock()

	if counter != requests {
		t.Fatalf(
			"expected counter to be %d after concurrent increments, got %d",
			requests,
			counter,
		)
	}
}

func TestLandingPageContainsFeatures(t *testing.T) {
	html := landing_page()

	features := []string{
		"Node tree, not strings",
		"Escaping by default",
		"Typed CSS, every property",
		"Go talks to your HTML",
	}

	for _, feature := range features {
		if !strings.Contains(html, feature) {
			t.Errorf("landing page missing feature %q", feature)
		}
	}
}
