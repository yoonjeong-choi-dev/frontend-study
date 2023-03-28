package handling_tmp_files

import (
	"fmt"
	"io/ioutil"
	"os"
)

func HandlingTmpExample() error {
	// 임시 폴더 생성
	tempDir, err := ioutil.TempDir("", "tmp")
	if err != nil {
		return err
	}

	// 함수 종료 후, 임시 폴더 전체 삭제
	defer os.RemoveAll(tempDir)

	tempFile, err := ioutil.TempFile(tempDir, "tmp")
	if err != nil {
		return err
	}

	fmt.Printf("Temp Directory Path: %s\n", tempDir)
	fmt.Printf("Temp File Path: %s\n", tempFile.Name())

	return nil
}
