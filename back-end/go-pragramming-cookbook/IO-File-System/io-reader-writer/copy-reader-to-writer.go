package io_reader_writer

import (
	"fmt"
	"io"
	"os"
)

// Copy :copy from in to out
// io.ReadSeeker 구현체를 이용하여 reader 를 여러 번 읽기 가능
func Copy(in io.ReadSeeker, out io.Writer) error {
	// 표준 출력에도 복사 i.e 표준 출력을 통해 in 데이터 출력
	writer := io.MultiWriter(out, os.Stdout)

	// 복사 1: 전체 복사
	if _, err := io.Copy(writer, in); err != nil {
		// writer 로 복사가 불가능한 경우 에러
		return err
	}

	// 두번째 복사 수행을 위해 reader의 포인터를 0으로 초기화
	in.Seek(0, 0)
	buf := make([]byte, 64)

	// 복사 2: 버퍼 크기만큼 데이터 복사
	if _, err := io.CopyBuffer(writer, in, buf); err != nil {
		return err
	}

	fmt.Println()
	return nil
}
