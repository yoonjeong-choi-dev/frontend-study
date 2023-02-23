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
	name     string
}

// Custom Error
var errPosArgsSpecified = errors.New("more than one positional arguments specified")

// flag 패키지 사용으로 인한 변경점
func parseArgs(w io.Writer, args []string) (config, error) {
	c := config{}

	flagSet := flag.NewFlagSet("greeter", flag.ContinueOnError)
	flagSet.SetOutput(w)

	// custom usage message
	flagSet.Usage = func() {
		var usageString = `
A greeter application which prints the name you entered a specified number of times.

Usage of %s: <options> [name]`

		fmt.Fprintf(w, usageString, flagSet.Name())
		fmt.Fprintln(w)
		fmt.Fprintln(w)
		fmt.Fprintln(w, "Options: ")
		// flag 패키지 기본 도움말 메시지 출력
		flagSet.PrintDefaults()
	}

	flagSet.IntVar(&c.numTimes, "n", 0, "Number of times to greet")

	err := flagSet.Parse(args)
	if err != nil {
		return c, err
	}

	if flagSet.NArg() > 1 {
		return c, errPosArgsSpecified
	}
	if flagSet.NArg() == 1 {
		c.name = flagSet.Arg(0)
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
	// 위치 인수로 이름을 받지 않은 경우에만 실행
	var err error
	if len(c.name) == 0 {
		c.name, err = getName(r, w)
		if err != nil {
			return err
		}
	}

	greetUser(c, w)
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
func greetUser(c config, w io.Writer) {
	msg := fmt.Sprintf("Nice to meet you~, %s\n", c.name)
	for i := 0; i < c.numTimes; i++ {
		fmt.Fprintf(w, msg)
	}
}

func main() {
	c, err := parseArgs(os.Stderr, os.Args[1:])
	if err != nil {
		// 커스텀 에러인 경우에만 출력(나머지는 flag 패키지에서 출력해줌)
		if errors.Is(err, errPosArgsSpecified) {
			fmt.Fprintln(os.Stdout, err)
		}

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
