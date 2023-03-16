package utils

import "encoding/json"

func GetJsonString(v interface{}) (string, error) {
	vJson, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		return "", err
	}
	return string(vJson), nil
}

func GetJsonStringUnsafe(v interface{}) string {
	result, _ := GetJsonString(v)
	return result
}
