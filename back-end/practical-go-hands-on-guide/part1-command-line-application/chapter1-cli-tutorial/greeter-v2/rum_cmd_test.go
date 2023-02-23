package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestRunCmd(t *testing.T) {
	// io.Reader 및 io.Writer 를 인자로 받는 함수이므로, 모킹해야 함
	tests := []struct {
		c      config
		input  string
		output string
		err    error
	}{
		{
			c:      config{numTimes: 7},
			input:  "",
			output: strings.Repeat("Your name please? Press the Enter key when done.\n", 1),
			err:    errors.New("You didn't enter your name"),
		},
		{
			c:     config{numTimes: 7},
			input: "Yoonjeong Choi",
			output: strings.Repeat("Your name please? Press the Enter key when done.\n", 1) +
				strings.Repeat("Nice to meet you~, Yoonjeong Choi\n", 7),
		},
		// test case for new feature : name input
		{
			c:      config{numTimes: 7, name: "Yoonjeong"},
			input:  "",
			output: strings.Repeat("Nice to meet you~, Yoonjeong\n", 7),
		},
	}

	// io.Writer mocking
	byteBuf := new(bytes.Buffer)
	for _, tc := range tests {
		// io.Reader mocking
		rd := strings.NewReader(tc.input)

		err := runCmd(rd, byteBuf, tc.c)
		if err != nil && tc.err == nil {
			t.Fatalf("Expected nil error, get: %v\n", err)
		}
		if err != nil && err.Error() != tc.err.Error() {
			t.Fatalf("Expected error to be: %v, got: %v\n", tc.err, err)
		}

		gotMsg := byteBuf.String()
		if gotMsg != tc.output {
			t.Errorf("Expected stdout message to be %v, Get: %v\n", tc.output, gotMsg)
		}
		byteBuf.Reset()
	}
}
