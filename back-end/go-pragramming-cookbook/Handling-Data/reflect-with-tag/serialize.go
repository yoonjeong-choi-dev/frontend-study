package reflect_with_tag

import (
	"fmt"
	"reflect"
)

const SerializeTagKey = "serialize"

// SerializeStringsStruct 구조체 직렬화
// 구조체 태그가 serialize:"key_name" 형식인 필드들만 직렬화
func SerializeStringsStruct(s interface{}) (string, error) {
	ret := ""

	// reflection 이용하여 타입 가져온다
	sType := reflect.TypeOf(s)

	// 가져온 타입을 이용하여 타입에 대한 인스턴스 생성
	value := reflect.ValueOf(s)

	// 구조체에 대한 포인터가 전달된 경우, 포인터에 대한 값으로 변경
	if sType.Kind() == reflect.Ptr {
		sType = sType.Elem()
		value = value.Elem()
	}

	// 모든 필드에 대한 직렬화 시작
	for i := 0; i < sType.NumField(); i++ {
		field := sType.Field(i)

		// 기본적으로는 필드 이름으로 직렬화
		key := field.Name

		// 구조체 태그가 있는 경우, 해당 태그의 키 이름 설정
		if serializeName, ok := field.Tag.Lookup(SerializeTagKey); ok {
			// 구조체 태그 값이 -인 경우 무시
			if serializeName == "-" {
				continue
			}

			key = serializeName
		}

		// 필드의 각 타입 별로 직렬화
		// 현재는 문자열 및 숫자 타입만 직렳화
		switch value.Field(i).Kind() {
		case reflect.String:
			ret += fmt.Sprintf("%s:%s;", key, value.Field(i).String())
		case reflect.Int:
			ret += fmt.Sprintf("%s:%d;", key, value.Field(i).Int())
		default:
			continue
		}
	}

	return ret, nil
}
