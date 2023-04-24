package mocking

import (
	"errors"
	"testing"
)

// TODO
func TestDoSomeStuff(t *testing.T) {
	tests := []struct {
		name       string
		DoStuff    error // DoStuff(input string) error 반환값
		ThrowError error
		wantErr    bool
	}{
		{"base-case", nil, nil, false},
		{"DoStuff error", nil, errors.New("failed"), true},
		{"ThrowError error", nil, errors.New("failed"), true},
	}

	for i, test := range tests {
		// 모의 객체를 이용하여 DoSomeStuff 테스트
		mock := MockDoStuffer{}
		mock.MockDoStuff = func(string) error {
			return test.DoStuff
		}

		//defer Patch(&ThrowError, func() error {
		//	return test.ThrowError
		//}).Restore()

		restore := Patch(&ThrowError, func() error {
			return test.ThrowError
		})

		err := DoSomeStuff(&mock)
		if (err != nil) != test.wantErr {
			t.Errorf("DoSomeStuff() error = %v, wantErr = %v\n", err, test.wantErr)
		}

		t.Logf("%d Test has err: %v\n", i, err)
		restore()
		t.Logf("%d Test has err after restore: %v\n", i, err)
	}
}
