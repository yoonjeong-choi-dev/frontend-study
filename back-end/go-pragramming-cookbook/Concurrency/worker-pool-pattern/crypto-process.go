package worker_pool_pattern

import "golang.org/x/crypto/bcrypt"

func hashProcess(r CryptoRequest) CryptoResponse {
	val, err := bcrypt.GenerateFromPassword(r.Text, bcrypt.DefaultCost)
	return CryptoResponse{
		Request: r,
		Result:  val,
		Err:     err,
	}
}

func compareProcess(r CryptoRequest) CryptoResponse {
	matched := true
	err := bcrypt.CompareHashAndPassword(r.Hashed, r.Text)
	if err != nil {
		matched = false
	}

	return CryptoResponse{
		Request: r,
		Matched: matched,
		Err:     err,
	}
}
