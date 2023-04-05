package main

import (
	"currency"
	"fmt"
)

func main() {
	dollar := "123.73"

	penny, err := currency.ConvertStringDollarsToPennies(dollar)
	if err != nil {
		panic(err)
	}

	fmt.Printf("dollar: %s -> penney: %d\n", dollar, penny)

	penny -= 12300
	dollar = currency.ConvertPenniesToDollarString(penny)
	fmt.Printf("penny: %d -> dollar: %s\n", penny, dollar)
}
