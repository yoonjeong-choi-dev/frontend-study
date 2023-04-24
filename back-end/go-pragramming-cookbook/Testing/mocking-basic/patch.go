package mocking

import "reflect"

// Restorer 이전 상태를 복구하는 함수 Restore() 를 가진다
type Restorer func()

// Restore 이전 상태로 복구하는 함수
func (r Restorer) Restore() {
	r()
}

// Patch dest 매개변수에 value 매개변수를 할당
// 반환 함수는 dest 매개변수의 값을 이전 값(원래 dest 값)으로 복원
func Patch(dest, value interface{}) Restorer {
	destValue := reflect.ValueOf(dest).Elem()

	// oldValue: dest 현재 값
	oldValue := reflect.New(destValue.Type()).Elem()
	oldValue.Set(destValue)

	// valueValue: value 객체의 값
	// => dest 에 할당
	valueValue := reflect.ValueOf(value)
	if !valueValue.IsValid() {
		valueValue = reflect.Zero(destValue.Type())
	}

	// value 값으로 업데이트(패치)
	destValue.Set(valueValue)

	// 이전 값으로 복원하는 복원 함수 반환
	return func() {
		destValue.Set(oldValue)
	}
}
