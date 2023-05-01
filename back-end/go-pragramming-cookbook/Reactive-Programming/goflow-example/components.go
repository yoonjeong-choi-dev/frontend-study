package goflow_example

import (
	"encoding/base64"
	"fmt"
)

// Encoder Val 값을 Base64로 인코딩하는 컴포넌트
type Encoder struct {
	Val           <-chan string // 문자열을 저장하는 read-only channel
	ResultMessage chan<- string // 문자열을 저장하는 write-only channel
}

func (e *Encoder) Process() {

	for val := range e.Val {
		encoded := base64.StdEncoding.EncodeToString([]byte(val))
		e.ResultMessage <- fmt.Sprintf("%s => %s", val, encoded)
	}
}

// Printer 표준출력으로 출력하는 프린터 컴포넌트
type Printer struct {
	Line <-chan string
}

func (p *Printer) Process() {
	for line := range p.Line {
		fmt.Println(line)
	}
}
