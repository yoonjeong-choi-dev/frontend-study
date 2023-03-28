package main

import (
	file_and_directory "file-and-directory"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Example 1: Handling os.File")
	err := file_and_directory.CopyWithCapitalizationExample(
		"src.txt",
		"this is content in src file.\nIt would be converted to upper cases\nabcde~",
		"dst.txt",
		true,
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("\n\nExample 2: Handling directory and file")
	err = file_and_directory.OperateDirectoryAndFile(
		"yj",
		"test.txt",
		"Test Content to Write into the created file",
		true,
	)
	if err != nil {
		log.Fatal(err.Error())
	}
}
