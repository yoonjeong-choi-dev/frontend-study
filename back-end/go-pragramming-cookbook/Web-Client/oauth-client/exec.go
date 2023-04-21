package oauth_client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getJsonString(v interface{}) (string, error) {
	vJson, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		return "", err
	}
	return string(vJson), nil
}

func GetUserInfo(client *http.Client) error {
	url := fmt.Sprintf("https://api.github.com/user")

	res, err := client.Get(url)
	if err != nil {
		return err
	}
	defer func() { _ = res.Body.Close() }()

	fmt.Printf("Status Code from '%s': %d\n", url, res.StatusCode)
	fmt.Println("Response Body:")

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var jsonDecoded map[string]interface{}
	err = json.Unmarshal(data, &jsonDecoded)
	if err != nil {
		return err
	}

	parsed, err := getJsonString(jsonDecoded)
	if err != nil {
		return err
	}

	fmt.Println(parsed)

	return nil
}
