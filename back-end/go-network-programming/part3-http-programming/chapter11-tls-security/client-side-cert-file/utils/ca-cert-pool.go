package utils

import (
	"crypto/x509"
	"errors"
	"io/ioutil"
)

// CACertPool 인증서 고정을 위해 인증서 풀을 등록하는 유틸함수
func CACertPool(certFilename string) (*x509.CertPool, error) {
	caCert, err := ioutil.ReadFile(certFilename)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(caCert); !ok {
		return nil, errors.New("failed to add certificate to pool")
	}
	return certPool, nil
}
