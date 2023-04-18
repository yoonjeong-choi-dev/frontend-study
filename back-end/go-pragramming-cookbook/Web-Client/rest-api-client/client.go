package restapi

import "net/http"

// APIClient custom client
// 프론트엔드에서 설정된 axios 객체 사용하는 것과 비슷한 패턴
type APIClient struct {
	*http.Client
}

// NewAPIClient 인증정보가 담긴 http.Client
func NewAPIClient(username, password string) *APIClient {
	t := http.Transport{}
	return &APIClient{
		Client: &http.Client{
			Transport: &APITransport{
				Transport: &t,
				username:  username,
				password:  password,
			},
		},
	}
}

// API 호출 메서드들 구현

func (c *APIClient) GetGoogle() (int, error) {
	res, err := c.Get("http://www.google.com")
	if err != nil {
		return 0, err
	}

	return res.StatusCode, nil
}
