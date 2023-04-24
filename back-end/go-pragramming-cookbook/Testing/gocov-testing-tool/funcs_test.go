package gocov_testing_tool

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_example1(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"function call test"},
	}

	for _, tt := range tests {
		Convey(tt.name, t, func() {
			res := exampleFunc1()
			So(res, ShouldBeNil)
		})
	}
}

func Test_example2(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"function type variable test"},
	}

	for _, tt := range tests {
		Convey(tt.name, t, func() {
			res := exampleFunc2()
			So(res, ShouldEqual, 7)
		})
	}
}
