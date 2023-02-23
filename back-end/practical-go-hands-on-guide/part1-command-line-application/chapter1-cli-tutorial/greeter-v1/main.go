package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type config struct {
	numTimes int
}

// flag 패키지 사용으로 인한 변경점
func parseArgs(w io.Writer, args []string) (config, error) {
	c := config{}

	flagSet := flag.NewFlagSet("greeter", flag.ContinueOnError)
	flagSet.SetOutput(w)
	flagSet.IntVar(&c.numTimes, "n", 0, "Number of times to greet")

	err := flagSet.Parse(args)
	if err != nil {
		return c, err
	}

	if flagSet.NArg() != 0 {
		return c, errors.New("positional arguments specified")
	}

	return c, nil
}

func validateArgs(c config) error {
	if c.numTimes <= 0 {
		return errors.New("must specify a number greater than 0")
	}
	return nil
}

func runCmd(r io.Reader, w io.Writer, c config) error {
	name, err := getName(r, w)
	if err != nil {
		return err
	}

	greetUser(c, name, w)
	return nil
}

func getName(r io.Reader, w io.Writer) (string, error) {
	msg := "Your name please? Press the Enter key when done.\n"
	fmt.Fprintf(w, msg)

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	name := scanner.Text()
	if len(name) == 0 {
		return "", errors.New("You didn't enter your name")
	}

	return name, nil
}
func greetUser(c config, name string, w io.Writer) {
	msg := fmt.Sprintf("Nice to meet you~, %s\n", name)
	for i := 0; i < c.numTimes; i++ {
		fmt.Fprintf(w, msg)
	}
}

func main() {
	c, err := parseArgs(os.Stderr, os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}

	err = validateArgs(c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}

	err = runCmd(os.Stdin, os.Stdout, c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}
