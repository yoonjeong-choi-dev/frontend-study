package simple_client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func createMockHttpGetServer() *httptest.Server {
	mockData := `[
{"name": "package1", "version": "1.1"},
{"name": "package2", "version": "1.2"}
]`

	ts := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprintf(w, mockData)
			}))

	return ts
}

func TestFetchPkgData(t *testing.T) {
	mock := createMockHttpGetServer()
	defer mock.Close()

	packages, err := FetchPkgData(mock.URL)

	if err != nil {
		t.Fatal(err)
	}

	if len(packages) != 2 {
		t.Fatalf("Expected 2 packages, Got back: %d\n", len(packages))
	}

	t.Logf("\nResponse : %#v\n", packages)
}
