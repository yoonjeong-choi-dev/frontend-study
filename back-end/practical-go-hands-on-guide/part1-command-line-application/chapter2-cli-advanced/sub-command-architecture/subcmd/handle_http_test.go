package subcmd

import (
	"bytes"
	"errors"
	"testing"
)

func TestHandleHttp(t *testing.T) {
	usageMessage := `
http: A HTTP client.

http: <options> server

Options: 
  -method string
    	HTTP method (default "GET")
`
	testConfigs := []struct {
		args   []string
		output string
		err    error
	}{
		{
			args: []string{},
			err:  ErrNoServerSpecified,
		},
		{
			args:   []string{"-h"},
			err:    errors.New("flag: help requested"),
			output: usageMessage,
		},
		{
			args:   []string{"yj.com"},
			err:    nil,
			output: "Executing http command with config - subcmd.httpConfig{url:\"yj.com\", verb:\"GET\"}\n",
		},
		{
			args: []string{"yj.com", "another positional args"},
			err:  ErrNoServerSpecified,
		},
		{
			args:   []string{"-method", "POST", "yj.com"},
			err:    nil,
			output: "Executing http command with config - subcmd.httpConfig{url:\"yj.com\", verb:\"POST\"}\n",
		},
		// Exercise 2.2
		{
			args: []string{"-method", "PUT", "yj.com"},
			err:  ErrInvalidHTTPMethod,
		},
	}

	// Mocking io.writer
	byteBuf := new(bytes.Buffer)

	for _, tc := range testConfigs {
		t.Logf("Test : %#v\n", tc.args)
		err := HandleHttp(byteBuf, tc.args)

		if tc.err == nil && err != nil {
			t.Fatalf("Expected nil error, get: %v\n", err)
		}

		if tc.err != nil && tc.err.Error() != err.Error() {
			t.Fatalf("Expected error to be: %v, got: %v\n", tc.err, err)
		}

		if len(tc.output) != 0 {
			gotMsg := byteBuf.String()
			if gotMsg != tc.output {
				t.Errorf("Expected stdout message to be %v, Get: %v\n", tc.output, gotMsg)
			}
		}
		byteBuf.Reset()
	}
}
