package bytes_and_strings

import (
	"bytes"
	"io"
	"io/ioutil"
)

func Buffer1(rawStr string) *bytes.Buffer {
	// 문자열을 바이트로 인코딩
	rawBytes := []byte(rawStr)

	// 방법 1:
	buffer := new(bytes.Buffer)
	buffer.Write(rawBytes)
	return buffer
}

func Buffer2(rawStr string) *bytes.Buffer {
	// 문자열을 바이트로 인코딩
	rawBytes := []byte(rawStr)

	// 방법 2:
	buffer := bytes.NewBuffer(rawBytes)
	return buffer
}

func Buffer3(rawStr string) *bytes.Buffer {
	// 방법 3: 바이트로 인코딩된 변수 필요 X
	return bytes.NewBufferString(rawStr)
}

func ReaderToString(buffer io.Reader) (string, error) {
	b, err := ioutil.ReadAll(buffer)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
