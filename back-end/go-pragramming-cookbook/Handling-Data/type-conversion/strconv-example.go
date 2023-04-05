package main

import (
	"fmt"
	"strconv"
)

func StringConvert() error {
	numStr := "7166"

	// string -> 64 bit integer
	res, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		return err
	}
	fmt.Printf("%s -> (base10)%d (type: %T)\n", numStr, res, res)

	// string -> 64 bit integer based 16
	res, err = strconv.ParseInt(numStr, 16, 64)
	if err != nil {
		return err
	}
	fmt.Printf("%s -> (base16)%d (type: %T)\n", numStr, res, res)

	// string -> bool
	boolStr := "true"
	bRes, err := strconv.ParseBool(boolStr)
	if err != nil {
		return err
	}
	fmt.Printf("%s -> %v (type %T)\n", boolStr, bRes, bRes)
	return nil
}
