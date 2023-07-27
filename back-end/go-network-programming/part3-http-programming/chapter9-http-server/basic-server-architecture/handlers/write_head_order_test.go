package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteHeaderFirst(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		// Write Head First
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
	}

	r := httptest.NewRequest(http.MethodGet, "http://dummyhost", nil)
	w := httptest.NewRecorder()
	handler(w, r)

	t.Logf("Response status: %q", w.Result().Status)

	body, err := ioutil.ReadAll(w.Result().Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Response body: %s\n", string(body))
}

func TestWriteBodyFirst(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		// Write Body First
		// => Status Code must be 200
		w.Write([]byte("bad request"))
		w.WriteHeader(http.StatusBadRequest)
	}

	r := httptest.NewRequest(http.MethodGet, "http://dummyhost", nil)
	w := httptest.NewRecorder()
	handler(w, r)

	t.Logf("Response status: %q", w.Result().Status)

	body, err := ioutil.ReadAll(w.Result().Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Response body: %s\n", string(body))
}
