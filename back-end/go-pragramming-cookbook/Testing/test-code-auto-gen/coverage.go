package test_code_auto_gen

import "errors"

func Coverage(condition bool) error {
	if condition {
		return errors.New("condition is true")
	}
	return nil
}
