package simple_client

import (
	"encoding/json"
	"io"
	"net/http"
)

type PkgData struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func FetchPkgData(url string) ([]PkgData, error) {
	var packages []PkgData

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.Header.Get("Content-Type") != "application/json" {
		return packages, nil
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &packages)
	return packages, err
}
