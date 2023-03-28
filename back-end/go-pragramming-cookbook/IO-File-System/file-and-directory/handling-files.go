package file_and_directory

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func CopyWithCapitalization(src *os.File, dst *os.File) error {
	// validate src file
	if _, err := src.Seek(0, io.SeekStart); err != nil {
		return err
	}

	// read src file with buffer
	contentsBuffer := new(bytes.Buffer)

	// io.Writer 인터페이스 구현체
	// => io-interface 예제와 동일하게 사용 가능
	if _, err := io.Copy(contentsBuffer, src); err != nil {
		return err
	}

	toCapital := strings.ToUpper(contentsBuffer.String())
	contentsReader := strings.NewReader(toCapital)
	if _, err := io.Copy(dst, contentsReader); err != nil {
		return err
	}

	// file close
	if err := src.Close(); err != nil {
		return err
	}
	if err := dst.Close(); err != nil {
		return err
	}

	return nil
}

func CopyWithCapitalizationExample(src, content, dst string, cleanup bool) error {
	fmt.Printf("Create %s...\n", src)
	f1, err := os.Create(src)
	if err != nil {
		return err
	}

	if _, err := f1.Write([]byte(content)); err != nil {
		return err
	}

	fmt.Printf("Create %s...\n", dst)
	f2, err := os.Create(dst)
	if err != nil {
		return nil
	}

	if err := CopyWithCapitalization(f1, f2); err != nil {
		return err
	}

	if cleanup {
		fmt.Printf("Remove %s...\n", src)
		if err := os.Remove(src); err != nil {
			return err
		}

		fmt.Printf("Remove %s...\n", dst)
		if err := os.Remove(dst); err != nil {
			return err
		}
	}

	return nil
}
