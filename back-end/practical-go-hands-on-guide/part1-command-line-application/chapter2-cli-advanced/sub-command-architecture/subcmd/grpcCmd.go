package subcmd

import (
	"flag"
	"fmt"
	"io"
)

type grpcConfig struct {
	server string
	method string
	body   string
}

func HandleGrpc(w io.Writer, args []string) error {
	config := grpcConfig{}
	fs := flag.NewFlagSet("grpc", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.StringVar(&config.method, "method", "", "Method to call")
	fs.StringVar(&config.body, "body", "", "Body of request")

	fs.Usage = func() {
		var usageString = `
grpc: A gRPC client.

grpc: <options> server`
		fmt.Fprintf(w, usageString)

		fmt.Fprintln(w)
		fmt.Fprintln(w)
		fmt.Fprintln(w, "Options: ")
		fs.PrintDefaults()
	}

	err := fs.Parse(args)
	if err != nil {
		return err
	}

	if fs.NArg() != 1 {
		return ErrNoServerSpecified
	}
	config.server = fs.Arg(0)

	fmt.Fprintf(w, "Executing grpc command with config - %#v\n", config)
	return nil
}
