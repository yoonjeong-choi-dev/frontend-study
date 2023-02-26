package main

import (
	"bytes"
	"testing"
)

func TestHandleCommand(t *testing.T) {
	usageMessage := `Usage: yjnc [http|grpc] -h

http: A HTTP client.

http: <options> server

Options: 
  -method string
    	HTTP method (default "GET")

grpc: A gRPC client.

grpc: <options> server

Options: 
  -body string
    	Body of request
  -method string
    	Method to call
`
	testConfigs := []struct {
		args   []string
		output string
		err    error
	}{
		{
			args:   []string{},
			err:    errInvalidSubCommand,
			output: "invalid sub-command specified\n" + usageMessage,
		},
		{
			args:   []string{"-h"},
			err:    nil,
			output: usageMessage,
		},
		{
			args:   []string{"wrong-command", "ignored"},
			err:    errInvalidSubCommand,
			output: "invalid sub-command specified\n" + usageMessage,
		},
		// Exercise 2.1
		{
			args:   []string{"http", "-method", "POST", "yj.com"},
			err:    nil,
			output: "Executing http command with config - subcmd.httpConfig{url:\"yj.com\", verb:\"POST\"}\n",
		},
	}

	// Mocking io.writer
	byteBuf := new(bytes.Buffer)

	for _, tc := range testConfigs {
		err := handleCommand(byteBuf, tc.args)

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
