package test_code_auto_gen

import "testing"

// go install github.com/cweill/gotests/gotests@latest
// gotests -all -w .
func TestCoverage(t *testing.T) {
	type args struct {
		condition bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// 이 부분만 채워주면 자동으로 테스트
		{"set condition", args{true}, true},
		{"no condition", args{false}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Coverage(tt.args.condition); (err != nil) != tt.wantErr {
				t.Errorf("Coverage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
