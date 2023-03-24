package io_reader_writer

import (
	"io"
	"os"
)

func PipeExample() error {
	// 인메모리 파이프로 reader, writer 연결
	// 하나의 데이터 스트림의 생산자 및 소비자 역할
	reader, writer := io.Pipe()

	go func() {
		writer.Write([]byte("This is a Test - "))
		writer.Write([]byte("Pipe Writer\n"))
		writer.Close()
	}()

	if _, err := io.Copy(os.Stdout, reader); err != nil {
		return err
	}

	return nil
}
