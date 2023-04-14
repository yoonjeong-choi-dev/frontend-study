package tweak

import (
	"fmt"
	"strings"
)

type StringTweaker struct{}

type Args struct {
	String  string
	ToUpper bool
	Reverse bool
}

// Tweak method for RPC
// RPC 라이브러리 조건
// 0) 함수 및 매개변수의 타입은 built-in 이거나 public 객체
// 1) 함수는 2개의 포인터 매개변수를 가지며, 에러 타입 반환
// 2) 첫번째 매개변수는 input, 두번째 매개변수는 output
func (s StringTweaker) Tweak(args *Args, resp *string) error {
	ret := string(args.String)
	if args.ToUpper {
		ret = strings.ToUpper(ret)
	}
	if args.Reverse {
		runes := []rune(ret)

		// swap left-side & right-side
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}

		ret = string(runes)
	}

	// save the response
	*resp = ret
	return nil
}

func (s StringTweaker) Echo(req *string, res *string) error {
	*res = fmt.Sprintf("[Echo] %s", *req)
	return nil
}
