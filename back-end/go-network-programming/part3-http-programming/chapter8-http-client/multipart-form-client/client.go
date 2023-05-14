package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// multipart-form 요청 바디 버퍼
	reqBody := new(bytes.Buffer)
	w := multipart.NewWriter(reqBody)

	// Add form fields
	formFields := map[string]string{
		"date":        time.Now().Format(time.RFC3339),
		"description": "Form Values with Attached Files",
	}
	for k, v := range formFields {
		err := w.WriteField(k, v)
		if err != nil {
			panic(err)
		}
	}

	// Attach files
	filesToAttach := []string{"./tmp/test.txt", "./tmp/multipart.txt"}
	for i, fileName := range filesToAttach {
		// 파일 정보 저장
		filePart, err := w.CreateFormFile(fmt.Sprintf("fileName-field-%d", i+1), filepath.Base(fileName))
		if err != nil {
			panic(err)
		}

		// 파일 바이너리 데이터 읽어서 저장
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		_, err = io.Copy(filePart, file)

		_ = file.Close()
		if err != nil {
			panic(err)
		}
	}

	// multipart writer 닫기
	// => 명시적으로 닫아야 요청 바디에 바운더리를 추가하는 작업이 처리됨
	err := w.Close()
	if err != nil {
		panic(err)
	}

	// multipart writer 이용하여 데이터를 저장한 바이트 버퍼를 이용하여 요청
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://httpbin.org/post", reqBody)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", w.FormDataContentType())

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer func() { _ = res.Body.Close() }()

	fmt.Printf("Response Status: %s\n", res.Status)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %s\n", string(resBody))
}
