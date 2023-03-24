package main

import (
	"bufio"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"robust-app/utils"
	"strconv"
)

func getUserInputString(r io.Reader, w io.Writer, title string) (string, error) {
	var ret string
	for {
		scanner := bufio.NewScanner(r)
		prompt := fmt.Sprintf("%s: ", title)
		fmt.Fprint(w, prompt)
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			return "", err
		}

		ret = scanner.Text()
		if ret == "" {
			fmt.Fprintf(w, "Please Enter '%s'\n", title)
		} else {
			break
		}
	}
	return ret, nil
}

func getUserInputInt(r io.Reader, w io.Writer, title string) (int, error) {
	for {
		scanner := bufio.NewScanner(r)
		prompt := fmt.Sprintf("%s: ", title)
		fmt.Fprint(w, prompt)
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			return -1, err
		}

		ret, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Fprintf(w, "Please Enter '%s' as Integer\n", title)
		} else {
			return ret, nil
		}
	}
}

func checkError(w io.Writer, errorType string, err error) bool {
	if err != nil {
		fmt.Fprintf(w, "Error for %s: %s\n", errorType, utils.GetJsonStringUnsafe(err))
		return true
	}
	return false
}

func printResponse(w io.Writer, res interface{}, err error) {
	errStatus := status.Convert(err)
	if errStatus.Code() == codes.OK {
		fmt.Fprintf(w, "Response:\n%s\n", utils.GetJsonStringUnsafe(res))
	} else {
		fmt.Fprintf(w, "Request failed: %v - %v\n", errStatus.Code(), errStatus.Message())
	}
}
