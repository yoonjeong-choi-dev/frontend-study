package worker_pool_pattern

import "errors"

type op string

const (
	Hash    op = "encrypt"
	Compare op = "decrypt"
)

type CryptoRequest struct {
	Op     op
	Text   []byte
	Hashed []byte // optional
}

type CryptoResponse struct {
	Request CryptoRequest
	Result  []byte // Hash 결과
	Matched bool   // Compare 결과
	Err     error
}

func CryptoProcess(r CryptoRequest) CryptoResponse {
	switch r.Op {
	case Hash:
		return hashProcess(r)
	case Compare:
		return compareProcess(r)
	default:
		return CryptoResponse{Err: errors.New("unsupported operation")}
	}
}
