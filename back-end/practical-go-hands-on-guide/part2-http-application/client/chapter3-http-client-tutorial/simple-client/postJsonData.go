package simple_client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type PkgRegisterResponse struct {
	ID string `json:"ID"`
}

func RegisterPkgData(url string, data PkgData) (PkgRegisterResponse, error) {
	result := PkgRegisterResponse{}

	reqBody, err := json.Marshal(data)
	if err != nil {
		return result, err
	}

	reader := bytes.NewReader(reqBody)
	res, err := http.Post(url, "application/json", reader)
	if err != nil {
		return result, err
	}

	defer res.Body.Close()
	resData, err := io.ReadAll(res.Body)
	if err != nil {
		return result, err
	}

	if res.StatusCode != http.StatusOK {
		return result, errors.New(string(resData))
	}

	err = json.Unmarshal(resData, &result)
	return result, err
}
