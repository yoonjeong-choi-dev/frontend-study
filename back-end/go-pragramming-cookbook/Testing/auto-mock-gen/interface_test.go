package auto_mock_gen

import (
	"auto-mock-gen/internal"
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestInterfaceExample(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGetSetter := internal.NewMockGetSetter(ctrl)

	// Set mocking 'Get' method
	var k string
	keyToTestMocking := "we can put any value here."
	mockGetSetter.EXPECT().Get(keyToTestMocking).Do(func(key string) {
		// 모킹이 제대로 되는지 확인하기 위한 작업
		k = key
	}).Return("", nil)

	customErr := errors.New("mock error")
	mockGetSetter.EXPECT().Get(gomock.Any()).Return("", customErr)

	// Test
	if _, err := mockGetSetter.Get(keyToTestMocking); err != nil {
		t.Errorf("got: %#v, expected: %#v\n", err, nil)
	}
	if k != keyToTestMocking {
		t.Errorf("got: %s, expected: %s\n", k, keyToTestMocking)
	}

	if _, err := mockGetSetter.Get("any key for error"); !errors.Is(err, customErr) {
		t.Errorf("got: %#v, expected: %#v\n", err, customErr)
	}
}
