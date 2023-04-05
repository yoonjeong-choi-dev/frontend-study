package currency

import "strconv"

type Penny = int64

func ConvertPenniesToDollarString(penny Penny) string {
	// convert int64 -> string with base 10
	result := strconv.FormatInt(penny, 10)

	isNegative := result[0] == '-'
	if isNegative {
		result = result[1:]
	}

	// 1달러 이하인 경우, 문자열 처리를 위해 앞에 0을 붙여준다
	for len(result) < 3 {
		result = "0" + result
	}

	length := len(result)
	result = result[0:length-2] + "." + result[length-2:]

	if isNegative {
		result = "-" + result
	}
	return result
}
