package simple_client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func createMockHttpPostServer(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(generateMockServerHandler(t)))
}

func generateMockServerHandler(t *testing.T) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		t.Log("Server Open")
		if r.Method == "POST" {
			reqData := PkgData{}
			resData := PkgRegisterResponse{}

			defer r.Body.Close()
			req, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			err = json.Unmarshal(req, &reqData)
			if err != nil || len(reqData.Name) == 0 || len(reqData.Version) == 0 {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}

			resData.ID = fmt.Sprintf("%s-%s", reqData.Name, reqData.Version)
			jsonData, err := json.Marshal(resData)
			if err != nil {
				t.Logf("Error from marshaling response body: %#v\n", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			p := PkgRegisterResponse{}
			json.Unmarshal(jsonData, &p)

			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, string(jsonData))
		} else {
			http.Error(w, "Invalid HTTP method specified", http.StatusMethodNotAllowed)
			return
		}
	}

}

func TestRegisterPkgDataPositive(t *testing.T) {
	mock := createMockHttpPostServer(t)
	defer mock.Close()

	expected := "TestPackage-1.7.0"
	req := PkgData{Name: "TestPackage", Version: "1.7.0"}

	res, err := RegisterPkgData(mock.URL, req)
	if err != nil {
		t.Fatalf("%#v\n", err)
		t.Fatal(err)
	}

	if res.ID != expected {
		t.Errorf("Expected: %s, Got: %s\n", expected, res.ID)
	}
	t.Logf("Response : %#v\n", res)
}

func TestRegisterPkgDataNegative(t *testing.T) {
	mock := createMockHttpPostServer(t)
	defer mock.Close()

	req := PkgData{}

	res, err := RegisterPkgData(mock.URL, req)
	if err == nil {
		t.Fatal("Expected error to be non-nil, got nil")
	}

	if len(res.ID) != 0 {
		t.Errorf("Expected package ID to be empty, got: %s\n", res.ID)
	}
	t.Logf("\nResponse : %#v\nError: %#v\n", res, err)
}
