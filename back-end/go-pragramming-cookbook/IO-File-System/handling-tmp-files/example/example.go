package main

import (
	"fmt"
	handling_tmp_files "handling-tmp-files"
	"log"
	"os"
)

func main() {
	fmt.Printf("Temp Directory: %s\n", os.TempDir())
	if err := handling_tmp_files.HandlingTmpExample(); err != nil {
		log.Fatalln(err)
	}
}
