package subcmd

import (
	"flag"
	"fmt"
	"io"
)

var validMethods = []string{"GET", "POST", "HEAD"}

type httpConfig struct {
	url  string
	verb string
}

func HandleHttp(w io.Writer, args []string) error {
	config := httpConfig{}
	fs := flag.NewFlagSet("http", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.StringVar(&config.verb, "method", "GET", "HTTP method")

	fs.Usage = func() {
		var usageString = `
http: A HTTP client.

http: <options> server`
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

	// Exercise 2.2
	if !Contains(validMethods, config.verb) {
		return ErrInvalidHTTPMethod
	}

	if fs.NArg() != 1 {
		return ErrNoServerSpecified
	}

	config.url = fs.Arg(0)
	fmt.Fprintf(w, "Executing http command with config - %#v\n", config)
	return nil
}
