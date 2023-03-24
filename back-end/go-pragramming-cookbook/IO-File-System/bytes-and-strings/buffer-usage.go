package bytes_and_strings

import (
	"bufio"
	"bytes"
	"fmt"
)

func BufferToStringExample(rawString string) error {
	var buffer = Buffer1(rawString)

	// String() 메서드를 이용한 문자열 변환
	fmt.Printf("Buffer -> String: %s\n", buffer.String())

	// bytes.Buffer 는 io.Reader 인터페이스 만족
	fmt.Print("bytes.buffer as io.Reader -> String: ")
	s, err := ReaderToString(buffer)
	if err != nil {
		return err
	}
	fmt.Println(s)

	// 버퍼 대신 bytes reader(io.Reader) 생성하여 이용
	bytesReader := bytes.NewReader([]byte(rawString))

	// 바이트 배열을 읽는 io.Reader 에 대한 스캐너
	scanner := bufio.NewScanner(bytesReader)

	// 스캐너를 통한 토큰화: 공백 기준
	scanner.Split(bufio.ScanWords)
	fmt.Print("Tokenized Result: ")
	for scanner.Scan() {
		fmt.Print(scanner.Text())
		fmt.Print(", ")
	}
	fmt.Println()

	return nil
}
