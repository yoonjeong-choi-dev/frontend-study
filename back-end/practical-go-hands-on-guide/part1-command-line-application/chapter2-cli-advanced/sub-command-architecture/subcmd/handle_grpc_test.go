package subcmd

import (
	"bytes"
	"errors"
	"testing"
)

func TestHandleGrpc(t *testing.T) {
	usageMessage := `
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
			output: "Executing grpc command with config - subcmd.grpcConfig{server:\"yj.com\", method:\"\", body:\"\"}\n",
		},
		{
			args: []string{"yj.com", "another positional args"},
			err:  ErrNoServerSpecified,
		},
		{
			args:   []string{"yj.com"},
			err:    nil,
			output: "Executing grpc command with config - subcmd.grpcConfig{server:\"yj.com\", method:\"\", body:\"\"}\n",
		},
		{
			args: []string{"yj.com", "another positional args"},
			err:  ErrNoServerSpecified,
		},
		{
			args:   []string{"-method", "POST", "yj.com"},
			err:    nil,
			output: "Executing grpc command with config - subcmd.grpcConfig{server:\"yj.com\", method:\"POST\", body:\"\"}\n",
		},
		{
			args:   []string{"-method", "POST", "-body", "Hello~", "yj.com"},
			err:    nil,
			output: "Executing grpc command with config - subcmd.grpcConfig{server:\"yj.com\", method:\"POST\", body:\"Hello~\"}\n",
		},
	}

	// Mocking io.writer
	byteBuf := new(bytes.Buffer)

	for _, tc := range testConfigs {
		t.Logf("Test : %#v\n", tc.args)
		err := HandleGrpc(byteBuf, tc.args)

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
