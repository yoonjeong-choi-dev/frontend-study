package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDecodeHandler(t *testing.T) {
	const jsonStream = `
	{"user_ip": "172.121.19.21", "event": "click_on_add_cart"}
  {"user_ip": "172.121.19.22", "event": "click_on_checkout"}
`
	body := strings.NewReader(jsonStream)

	req := httptest.NewRequest("POST", "http://test.com/", body)
	res := httptest.NewRecorder()

	decodeLogStreamHandler(res, req)
}
