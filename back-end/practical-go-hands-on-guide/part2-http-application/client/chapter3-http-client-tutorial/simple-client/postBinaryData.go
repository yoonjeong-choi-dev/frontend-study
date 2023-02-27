package simple_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

type PkgFileData struct {
	PkgData
	FileName string
	Bytes    io.Reader
}

type PkgFileDataRegisterResponse struct {
	Id       string `json:"id"`
	Filename string `json:"filename"`
	Size     int64  `json:"size"`
}

func RegisterPkgBinaryData(url string, data PkgFileData) (PkgFileDataRegisterResponse, error) {
	result := PkgFileDataRegisterResponse{}

	payload, contentType, err := createMultiPartMessage(data)
	if err != nil {
		return result, err
	}

	reader := bytes.NewReader(payload)
	res, err := http.Post(url, contentType, reader)
	if err != nil {
		return result, err
	}

	defer res.Body.Close()
	resData, err := io.ReadAll(res.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(resData, &result)
	return result, err
}

// createMultiPartMessage : multipart/form-data 페이로드 생성 함수
func createMultiPartMessage(data PkgFileData) ([]byte, string, error) {
	var buffer bytes.Buffer
	var err error
	var fw io.Writer

	// multipart 폼 데이터를 buffer 에 저장하는 writer 생성
	mw := multipart.NewWriter(&buffer)

	// 폼 데이터의 name 필드
	fw, err = mw.CreateFormField("name")
	if err != nil {
		return nil, "", err
	}
	fmt.Fprintf(fw, data.Name)

	// 폼 데이터의 version 필드
	fw, err = mw.CreateFormField("version")
	if err != nil {
		return nil, "", err
	}
	fmt.Fprintf(fw, data.Version)

	// 바이너리 데이터를 저장하기 위한 폼 필드
	fw, err = mw.CreateFormFile("filedata", data.FileName)
	if err != nil {
		return nil, "", err
	}

	// 파일 데이터(data.Bytes)를 폼 필드에 저장
	_, err = io.Copy(fw, data.Bytes)
	err = mw.Close()
	if err != nil {
		return nil, "", err
	}

	contentType := mw.FormDataContentType()
	return buffer.Bytes(), contentType, nil
}
