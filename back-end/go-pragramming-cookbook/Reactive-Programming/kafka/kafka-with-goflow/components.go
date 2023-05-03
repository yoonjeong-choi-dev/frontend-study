// Package kafka_with_goflow goflow 그래프를 구성하는 컴포넌트 정의
package kafka_with_goflow

import (
	"fmt"
	"github.com/trustmaster/goflow"
	"strings"
)

type Upper struct {
	Val    <-chan string // read-only string type channel
	Result chan<- string // write-only string type channel
}

func (c *Upper) Process() {
	for val := range c.Val {
		c.Result <- strings.ToUpper(val)
	}
}

type Printer struct {
	goflow.Component
	Line <-chan string
}

func (c *Printer) Process() {
	for line := range c.Line {
		fmt.Printf("[GoFlow Printer Component] %s\n", line)
	}
}
