package auto_mock_gen

import (
	"auto-mock-gen/internal"
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestController_GetThenSet(t *testing.T) {
	testKey := "testKey"

	// 모킹 인터페이스 객채의 Get 및 Set 메소드의 파라미터 및 반환 값 설정
	tests := []struct {
		name         string
		getReturnVal string
		getReturnErr error
		setReturnErr error
		wantErr      bool
	}{
		{"get error", "value", errors.New("failed"), nil, true},
		{"value match", "value", nil, nil, false},
		{"no errors", "not set", nil, nil, false},
		{"set error", "not set", nil, errors.New("failed"), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockGetSetter := internal.NewMockGetSetter(ctrl)

			// Mocking Get&Set method
			mockGetSetter.EXPECT().Get(testKey).AnyTimes().Return(tt.getReturnVal, tt.getReturnErr)
			mockGetSetter.EXPECT().Set(testKey, gomock.Any()).AnyTimes().Return(tt.setReturnErr)

			c := &Controller{
				GetSetter: mockGetSetter,
			}
			if err := c.GetThenSet(testKey, "value"); (err != nil) != tt.wantErr {
				t.Errorf("Controller.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
