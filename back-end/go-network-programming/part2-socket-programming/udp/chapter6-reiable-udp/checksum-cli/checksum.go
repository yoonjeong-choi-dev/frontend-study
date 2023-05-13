package main

import (
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func init() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s file...\n", os.Args)
		flag.PrintDefaults()
	}
}

func checksum(file string) string {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("%x", sha512.Sum512_256(b))
}

func main() {
	flag.Parse()
	for _, file := range flag.Args() {
		fmt.Printf("%s %s\n", checksum(file), file)
	}
}
