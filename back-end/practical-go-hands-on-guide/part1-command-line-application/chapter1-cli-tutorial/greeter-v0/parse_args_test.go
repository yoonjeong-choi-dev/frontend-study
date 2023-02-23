package main

import (
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
			err:    nil,
			config: config{printUsage: true, numTimes: 0},
		},
		{
			args:   []string{"--help"},
			err:    nil,
			config: config{printUsage: true, numTimes: 0},
		},
		{
			args:   []string{"7"},
			err:    nil,
			config: config{printUsage: false, numTimes: 7},
		},
		{
			args:   []string{"not-number"},
			err:    errors.New("strconv.Atoi: parsing \"not-number\": invalid syntax"),
			config: config{printUsage: false, numTimes: 0},
		},
		{
			args:   []string{},
			err:    errors.New("invalid number of arguments"),
			config: config{printUsage: false, numTimes: 0},
		},
	}

	for _, tc := range tests {
		c, err := parseArgs(tc.args)
		if tc.err != nil && err != nil && err.Error() != tc.err.Error() {
			t.Fatalf("Expected error to be: %v, got: %v\n", tc.err, err)
		}

		if tc.err == nil && err != nil {
			t.Errorf("Expected nil error, get: %v\n", err)
		}

		if c.printUsage != tc.printUsage {
			t.Errorf("Expected printUsage to be: %v, got: %v\n", tc.printUsage, c.printUsage)
		}
		if c.numTimes != tc.numTimes {
			t.Errorf("Expected numTimes to be: %v, got: %v\n", tc.numTimes, c.numTimes)
		}
	}
}
