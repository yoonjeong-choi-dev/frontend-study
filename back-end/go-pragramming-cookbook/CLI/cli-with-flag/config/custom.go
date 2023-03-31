package config

import (
	"fmt"
	"strconv"
	"strings"
)

// CountTheWays Custom Flag Type
type CountTheWays []int

func (c *CountTheWays) String() string {
	result := ""
	for _, val := range *c {
		if len(result) > 0 {
			result += " ... "
		}
		result += fmt.Sprint(val)
	}
	return result
}

func (c *CountTheWays) Set(value string) error {
	values := strings.Split(value, ",")

	for _, val := range values {
		num, err := strconv.Atoi(val)
		if err != nil {
			return err
		}
		*c = append(*c, num)
	}
	return nil
}
