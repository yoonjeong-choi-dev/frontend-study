package reflect_with_tag

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

// DeserializeStringsStruct 구조체 역직렬화
// 구조체 태그가 serialize:"key_name" 형식인 필드들만 역직렬화
func DeserializeStringsStruct(s string, res interface{}) error {
	resType := reflect.TypeOf(res)

	// 객체 포인터만 전달받을 수 있음(해당 객체에 역직렬화해야 하므로)
	if resType.Kind() != reflect.Ptr {
		return errors.New("res interface must be a pointer")
	}

	resType = resType.Elem()
	value := reflect.ValueOf(res).Elem()

	// 직렬화된 문자열을 ; 기준으로 토근화
	tokens := strings.Split(s, ";")
	valMap := make(map[string]string)
	for _, v := range tokens {
		// key:value 형태
		keyval := strings.Split(v, ":")

		// key:value 형태가 아닌 경우 무시
		if len(keyval) != 2 {
			continue
		}

		valMap[keyval[0]] = keyval[1]
	}

	// 모든 필드에 대한 역직렬화 시작
	for i := 0; i < resType.NumField(); i++ {
		field := resType.Field(i)

		// 구조체 태그가 있는 경우, 역직렬화
		var val string
		var hasVal = false
		if serializeName, ok := field.Tag.Lookup(SerializeTagKey); ok {
			// 구조체 태그 값이 -인 경우 무시
			if serializeName == "-" {
				continue
			}

			mapVal, ok := valMap[serializeName]
			if ok {
				val = mapVal
				hasVal = true
			}
		}

		if !hasVal {
			mapVal, ok := valMap[field.Name]
			if ok {
				val = mapVal
				hasVal = true
			}
		}

		if hasVal {
			// 필드의 각 타입 별로 직렬화
			// 현재는 문자열 타입만 직렳화
			switch value.Field(i).Kind() {
			case reflect.String:
				value.Field(i).SetString(val)
			case reflect.Int:
				intval, err := strconv.ParseInt(val, 10, 64)
				if err != nil {
					return err
				}
				value.Field(i).SetInt(intval)
			default:
				continue
			}
		}
	}

	return nil
}
