package helios

import (
	"net/http"
	"net/http/httptest"
	"net/url"
)

// Testing approach here inspired by the github client from Google:
// https://github.com/google/go-github/blob/master/github/github_test.go

var (
	testMux    *http.ServeMux
	testServer *httptest.Server

	// client is the Helios client used for tests
	client *Client
)

// setup initializes testServer and testMux, along with a helios.Client that talks
// to the test server. Use testMux.handleFunc to return test data, and make sure to
// call teardown() to tear down the testServer when done.
func setup() {
	testMux = http.NewServeMux()
	testServer = httptest.NewServer(testMux)

	url, _ := url.Parse(testServer.URL)
	client = NewClientForURL(url, nil)
}

func teardown() {
	testServer.Close()
}
