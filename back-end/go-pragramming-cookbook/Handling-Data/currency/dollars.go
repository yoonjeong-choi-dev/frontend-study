package currency

import (
	"errors"
	"strconv"
	"strings"
)

func ConvertStringDollarsToPennies(dollar string) (Penny, error) {
	// check the string can be converted to float64
	_, err := strconv.ParseFloat(dollar, 64)
	if err != nil {
		return 0, err
	}

	tokens := strings.Split(dollar, ".")

	// validate the format
	if len(tokens) > 2 {
		return 0, errors.New("dollar string must have format 'int*.int{2}'")
	}

	decimal := ""
	if len(tokens) == 2 {
		if len(tokens[1]) != 2 {
			return 0, errors.New("dollar string must have format 'int*.int{2}'")
		}
		decimal = tokens[1]
	}

	for len(decimal) < 2 {
		decimal += "0"
	}

	result := tokens[0] + decimal
	return strconv.ParseInt(result, 10, 64)
}
