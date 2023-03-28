package file_and_directory

import (
	"errors"
	"fmt"
	"os"
)

func OperateDirectoryAndFile(dirName, fileName, content string, cleanup bool) error {
	// create a directory if not exists
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		fmt.Printf("Create direactory: %s\n", dirName)
		if err := os.Mkdir(dirName, os.FileMode(0755)); err != nil {
			fmt.Printf("%#v\n", err)
			return err
		}
	}

	// move to dir
	if err := os.Chdir(dirName); err != nil {
		return err
	}

	// create file
	fmt.Printf("Create %s...\n", fileName)
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	// write file with byte array
	b := []byte(content)
	size, err := file.Write(b)
	if err != nil {
		return err
	}

	if size != len(b) {
		return errors.New("incorrect length returned from *os.File.Write")
	}

	// Close file
	if err := file.Close(); err != nil {
		return err
	}

	// Clean up
	if cleanup {
		fmt.Printf("Remove all: rm -r %s\n", dirName)
		if err := os.Chdir(".."); err != nil {
			return err
		}

		if err := os.RemoveAll(dirName); err != nil {
			return err
		}
	}

	return nil
}
