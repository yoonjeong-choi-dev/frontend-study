package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"subcommand/yjnc/subcmd"
)

var errInvalidSubCommand = errors.New("invalid sub-command specified")

func printUsage(w io.Writer) {
	fmt.Fprintf(w, "Usage: yjnc [http|grpc] -h\n")
	subcmd.HandleHttp(w, []string{"-h"})
	subcmd.HandleGrpc(w, []string{"-h"})
}

func handleCommand(w io.Writer, args []string) error {
	var err error
	if len(args) < 1 {
		err = errInvalidSubCommand
	} else {
		switch args[0] {
		case "http":
			err = subcmd.HandleHttp(w, args[1:])
		case "grpc":
			err = subcmd.HandleGrpc(w, args[1:])
		case "-h", "--help":
			printUsage(w)
		default:
			err = errInvalidSubCommand
		}
	}

	// check the error is a custom error
	if errors.Is(err, subcmd.ErrNoServerSpecified) ||
		errors.Is(err, errInvalidSubCommand) ||
		errors.Is(err, subcmd.ErrInvalidHTTPMethod) {
		fmt.Fprintln(w, err)
		printUsage(w)
	}
	return err
}

func main() {
	err := handleCommand(os.Stdout, os.Args[1:])
	if err != nil {
		os.Exit(1)
	}
}
