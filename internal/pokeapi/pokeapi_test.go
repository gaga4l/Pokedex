package pokeapi

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

var interval = 5 * time.Minute

func TestNewClient(t *testing.T) {
	cases := []struct {
		name    string
		timeout time.Duration
	}{
		{name: "no timeout", timeout: 0},
		{name: "custom timeout", timeout: 5 * time.Second},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			client := NewClient(c.timeout, interval)
			if client == nil {
				t.Fatal("NewClient() returned nil")
			}
			if client.httpClient.Timeout != c.timeout {
				t.Fatalf("expected timeout %v, got %v", c.timeout, client.httpClient.Timeout)
			}
		})
	}
}

type testTransport func(*http.Request) (*http.Response, error)

func (f testTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestListLocations(t *testing.T) {
	pageURL := baseURL + "/location-area?offset=20&limit=20"
	invalidURL := "://invalid"
	requestError := errors.New("request failed")

	cases := []struct {
		name         string
		pageURL      *string
		expectedURL  string
		body         string
		transportErr error
		wantError    bool
		wantResults  bool
		wantRequest  bool
	}{
		{
			name:        "default page",
			expectedURL: baseURL + "/location-area",
			body:        `{"count":2,"next":"https://pokeapi.co/api/v2/location-area?offset=20&limit=20","previous":null,"results":[{"name":"canalave-city-area","url":"https://pokeapi.co/api/v2/location-area/1/"}]}`,
			wantResults: true,
			wantRequest: true,
		},
		{
			name:        "custom page",
			pageURL:     &pageURL,
			expectedURL: pageURL,
			body:        `{"count":2,"next":null,"previous":"https://pokeapi.co/api/v2/location-area","results":[]}`,
			wantRequest: true,
		},
		{
			name:      "invalid URL",
			pageURL:   &invalidURL,
			wantError: true,
		},
		{
			name:         "request failure",
			expectedURL:  baseURL + "/location-area",
			transportErr: requestError,
			wantError:    true,
			wantRequest:  true,
		},
		{
			name:        "invalid JSON",
			expectedURL: baseURL + "/location-area",
			body:        "not JSON",
			wantError:   true,
			wantRequest: true,
		},
		{
			name:        "empty response",
			expectedURL: baseURL + "/location-area",
			wantError:   true,
			wantRequest: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// ListLocations currently creates its own client using the default transport.
			// Keep these tests sequential because they temporarily replace global state.
			originalTransport := http.DefaultTransport
			t.Cleanup(func() {
				http.DefaultTransport = originalTransport
			})

			requested := false
			http.DefaultTransport = testTransport(func(req *http.Request) (*http.Response, error) {
				requested = true
				if req.Method != http.MethodGet {
					t.Errorf("expected GET request, got %s", req.Method)
				}
				if req.URL.String() != c.expectedURL {
					t.Errorf("expected URL %q, got %q", c.expectedURL, req.URL.String())
				}
				if c.transportErr != nil {
					return nil, c.transportErr
				}
				return &http.Response{
					StatusCode: http.StatusOK,
					Header:     make(http.Header),
					Body:       io.NopCloser(strings.NewReader(c.body)),
					Request:    req,
				}, nil
			})

			client := NewClient(time.Second, interval)
			got, err := client.ListLocations(c.pageURL)
			if requested != c.wantRequest {
				t.Errorf("expected request made=%v, got %v", c.wantRequest, requested)
			}
			if (err != nil) != c.wantError {
				t.Fatalf("expected error=%v, got %v", c.wantError, err)
			}
			if c.transportErr != nil && !errors.Is(err, c.transportErr) {
				t.Fatalf("expected error %v, got %v", c.transportErr, err)
			}
			if c.wantError {
				return
			}
			if got.Count != 2 {
				t.Errorf("expected count 2, got %d", got.Count)
			}
			if c.wantResults {
				if got.Next == nil || *got.Next != pageURL {
					t.Errorf("expected next page %q, got %v", pageURL, got.Next)
				}
				if got.Previous != nil {
					t.Errorf("expected nil previous page, got %q", *got.Previous)
				}
				if len(got.Results) != 1 {
					t.Fatalf("expected 1 location, got %d", len(got.Results))
				}
				if got.Results[0].Name != "canalave-city-area" {
					t.Errorf("unexpected location name %q", got.Results[0].Name)
				}
				if got.Results[0].URL != baseURL+"/location-area/1/" {
					t.Errorf("unexpected location URL %q", got.Results[0].URL)
				}
			} else {
				if got.Next != nil {
					t.Errorf("expected nil next page, got %q", *got.Next)
				}
				if got.Previous == nil || *got.Previous != baseURL+"/location-area" {
					t.Errorf("expected previous page URL, got %v", got.Previous)
				}
				if len(got.Results) != 0 {
					t.Errorf("expected no locations, got %d", len(got.Results))
				}
			}
		})
	}
}
