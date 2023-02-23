package main

import (
	"bytes"
	"errors"
	"testing"
)

func TestParseArgs(t *testing.T) {
	// 테스트하는 함수의 인풋과 인풋에 대응하는 아웃풋
	type testConfig struct {
		args []string
		err  error
		config
	}

	tests := []testConfig{
		{
			args:   []string{"-h"},
			err:    errors.New("flag: help requested"),
			config: config{numTimes: 0},
		},
		{
			args:   []string{"--help"},
			err:    errors.New("flag: help requested"),
			config: config{numTimes: 0},
		},
		{
			args:   []string{"-n", "7"},
			err:    nil,
			config: config{numTimes: 7},
		},
		{
			args:   []string{"-n", "not-number"},
			err:    errors.New("invalid value \"not-number\" for flag -n: parse error"),
			config: config{numTimes: 0},
		},
		{
			args:   []string{},
			err:    errors.New("must specify a number greater than 0"),
			config: config{numTimes: 0},
		},
		{
			args:   []string{"-n", "1", "positional argument1", "positional argument2"},
			err:    errors.New("positional arguments specified"),
			config: config{numTimes: 1},
		},
	}

	// io.Writer mocking
	byteBuf := new(bytes.Buffer)
	for _, tc := range tests {
		c, err := parseArgs(byteBuf, tc.args)
		if tc.err == nil && err != nil {
			t.Errorf("Expected nil error, got: %v\n", err)
		}
		if tc.err != nil && err != nil && err.Error() != tc.err.Error() {
			t.Errorf("Expected error to be: %v, got: %v\n", tc.err, err)
		}

		if c.numTimes != tc.numTimes {
			t.Errorf("Expected numTimes to be: %v, got: %v\n", tc.numTimes, c.numTimes)
		}
		byteBuf.Reset()
	}
}
