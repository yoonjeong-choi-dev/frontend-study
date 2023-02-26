package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func handleCmdA(w io.Writer, args []string) error {
	var v string
	fs := flag.NewFlagSet("subcmd-a", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.StringVar(&v, "verb", "argument-value", "Argument 1")

	err := fs.Parse(args)
	if err != nil {
		return err
	}

	fmt.Fprintf(w, "Executing command A with verb - \"%s\"\n", v)
	return nil
}

func handleCmdB(w io.Writer, args []string) error {
	var v string
	fs := flag.NewFlagSet("subcmd-b", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.StringVar(&v, "verb", "argument-value", "Argument 1")

	err := fs.Parse(args)
	if err != nil {
		return err
	}

	fmt.Fprintf(w, "Executing command B with verb - \"%s\"\n", v)
	return nil
}

func printUsage(w io.Writer) {
	fmt.Fprintf(w, "Usage: %s [subcmd-a|subcmd-b] -h\n", os.Args[0])

	// 서브 커맨드 도움말은 서브 커맨드 내 flagSet에게 위임
	handleCmdA(w, []string{"-h"})
	handleCmdB(w, []string{"-h"})
}

func main() {
	var err error
	if len(os.Args) < 2 {
		printUsage(os.Stdout)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "subcmd-a":
		err = handleCmdA(os.Stdout, os.Args[2:])
	case "subcmd-b":
		err = handleCmdB(os.Stdout, os.Args[2:])
	default:
		printUsage(os.Stdout)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
