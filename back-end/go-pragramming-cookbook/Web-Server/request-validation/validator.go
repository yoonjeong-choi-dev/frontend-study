package request_validation

import "errors"

type ValidationError struct {
	error
}

type Payload struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func ValidatePayload(p *Payload) error {
	if p.Name == "" {
		return ValidationError{errors.New("name is required")}
	}

	if p.Age < 0 || p.Age > 120 {
		return ValidationError{
			errors.New("age is required and must be a value between 0 and 120"),
		}
	}
	return nil
}
