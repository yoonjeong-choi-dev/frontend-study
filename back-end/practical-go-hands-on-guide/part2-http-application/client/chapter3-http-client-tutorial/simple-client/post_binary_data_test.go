package simple_client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func createMockHttpMultipartFormServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(mockHttpMultipartPostServerHandler))
}

func mockHttpMultipartPostServerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		resData := PkgFileDataRegisterResponse{}

		// 요청 데이터의 사이즈는 최대 5000 바이트
		err := r.ParseMultipartForm(5000)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// multipart/form-data 페이로드
		form := r.MultipartForm

		// filedata : multipart/form-data 페이로드에서 바이너리 데이터를 저장한 필드 이름
		fileData := form.File["filedata"][0]
		resData.Id = fmt.Sprintf("%s-%s", form.Value["name"][0], form.Value["version"][0])
		resData.Filename = fileData.Filename
		resData.Size = fileData.Size

		jsonData, err := json.Marshal(resData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(jsonData))
	} else {
		http.Error(w, "Invalid HTTP method specified", http.StatusMethodNotAllowed)
		return
	}
}

func TestRegisterPkgBinaryData(t *testing.T) {
	mock := createMockHttpMultipartFormServer()
	defer mock.Close()

	reqData := PkgFileData{
		PkgData: PkgData{
			Name:    "test-package",
			Version: "1.8.1",
		},
		FileName: "test-package-1.8.1.tar.gz",
		Bytes:    strings.NewReader("This is a test binary data"),
	}

	res, err := RegisterPkgBinaryData(mock.URL, reqData)
	if err != nil {
		t.Fatal(err)
	}

	expectedId := fmt.Sprintf("%s-%s", reqData.Name, reqData.Version)
	expectedFileName := reqData.FileName

	if res.Id != expectedId {
		t.Errorf("Expected Id to be %s, got: %s\n", expectedId, res.Id)
	}

	if res.Filename != expectedFileName {
		t.Errorf("Expected FileName to be %s, got: %s\n", expectedFileName, res.Filename)
	}

	t.Logf("\nResponse: %#v\n", res)
}
