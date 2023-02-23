package main

import (
	"errors"
	"testing"
)

func TestValidateArgs(t *testing.T) {
	// TestParseArgs 와 다르게 익명 구조체를 이용하여 tc 정의
	tests := []struct {
		c   config
		err error
	}{
		{
			c:   config{},
			err: errors.New("must specify a number greater than 0"),
		},
		{
			c:   config{numTimes: -1},
			err: errors.New("must specify a number greater than 0"),
		},
		{
			c:   config{numTimes: 10},
			err: nil,
		},
	}

	for _, tc := range tests {
		err := validateArgs(tc.c)
		if tc.err != nil && err.Error() != tc.err.Error() {
			t.Errorf("Expected error to be: %v, got: %v\n", tc.err, err)
		}

		if tc.err == nil && err != nil {
			t.Errorf("Expected nil error, get: %v\n", err)
		}
	}
}
