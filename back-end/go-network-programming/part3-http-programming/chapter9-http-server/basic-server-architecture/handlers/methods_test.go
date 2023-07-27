package handlers

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

const DummyHost = "http://dummyhost"

func TestDefaultMethodsHandler_GET(t *testing.T) {
	handler := DefaultMethodsHandler()

	r := httptest.NewRequest(http.MethodGet, DummyHost, nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, r)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected: %d, got: %d\n", http.StatusOK, res.StatusCode)
	}
	t.Logf("Response Status: %s\n", res.Status)

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	_ = res.Body.Close()

	expected := "Hello, anonymous!"
	got := string(data)
	if expected != got {
		t.Fatalf("expected: %s, got: %s\n", expected, got)
	}
	t.Logf("Response Body: %s\n", got)
}

func TestDefaultMethodsHandler_POST(t *testing.T) {
	handler := DefaultMethodsHandler()
	reqBody := bytes.NewBufferString("<Tag Test>")

	r := httptest.NewRequest(http.MethodPost, DummyHost, reqBody)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, r)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected: %d, got: %d\n", http.StatusOK, res.StatusCode)
	}
	t.Logf("Response Status: %s\n", res.Status)

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	_ = res.Body.Close()

	expected := "Hello, &lt;Tag Test&gt;!"
	got := string(data)
	if expected != got {
		t.Fatalf("expected: %s, got: %s\n", expected, got)
	}
	t.Logf("Response Body: %s\n", got)
}

func TestDefaultMethodsHandler_Options(t *testing.T) {
	handler := DefaultMethodsHandler()

	r := httptest.NewRequest(http.MethodOptions, DummyHost, nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, r)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected: %d, got: %d\n", http.StatusOK, res.StatusCode)
	}
	t.Logf("Response Status: %s\n", res.Status)

	allows := res.Header.Get("Allow")
	if allows == "" {
		t.Fatalf("got empty allow header")
	}
	t.Logf("Response Allow Header: %s\n", allows)
}

func TestDefaultMethodsHandler_Head(t *testing.T) {
	handler := DefaultMethodsHandler()

	r := httptest.NewRequest(http.MethodHead, DummyHost, nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, r)

	res := w.Result()
	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Fatalf("expected: %d, got: %d\n", http.StatusMethodNotAllowed, res.StatusCode)
	}
	t.Logf("Response Status: %s\n", res.Status)

	allows := res.Header.Get("Allow")
	if allows == "" {
		t.Fatalf("got empty allow header")
	}
	t.Logf("Response Allow Header: %s\n", allows)
}
