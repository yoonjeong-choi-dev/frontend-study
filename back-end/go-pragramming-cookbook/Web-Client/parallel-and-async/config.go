package async

import "net/http"

// Client 비동기 병렬 처리를 위한 http.Client 구조체
// => 채널을 이용하여 응답 비동기 처리
// 팬아웃 방식: 가장 먼저 응답받은 요청부터 처리
type Client struct {
	*http.Client
	Response chan *http.Response
	Error    chan error
}

func (c *Client) AsyncGet(url string) {
	res, err := c.Get(url)
	if err != nil {
		c.Error <- err
		return
	}

	c.Response <- res
}

func CreateNewClient(client *http.Client, bufferSize int) *Client {
	resCh := make(chan *http.Response, bufferSize)
	errCh := make(chan error, bufferSize)

	return &Client{
		Client:   client,
		Response: resCh,
		Error:    errCh,
	}
}
