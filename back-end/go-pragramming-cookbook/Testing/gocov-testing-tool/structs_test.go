package gocov_testing_tool

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_exampleStruct_example3(t *testing.T) {
	type fields struct {
		Branch bool
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"no branch", fields{false}, false},
		{"with branch", fields{true}, true},
	}

	for _, tt := range tests {
		Convey(tt.name, t, func() {
			testStruct := &exampleStruct{
				Branch: tt.fields.Branch,
			}
			res := testStruct.exampleFunc3()
			So(res != nil, ShouldEqual, tt.wantErr)
		})
	}
}
