package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

func WordCount(r io.Reader) (map[string]int, error) {
	ret := make(map[string]int)

	// split by space
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		ret[scanner.Text()]++
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.New(fmt.Sprintf("error for reading input: %s", err))
	}

	return ret, nil
}

func main() {
	fmt.Println("Counting words of the left input in the pipe")

	// os.Stdin: 파일 객체로서 파이프 왼쪽 인풋을 읽을 수 있음
	counter, err := WordCount(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	for key, val := range counter {
		fmt.Printf("%s: %d\n", key, val)
	}
}
