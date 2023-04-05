package main

import (
	"fmt"
	"math"
)

func mathPackageExample() {
	n := 25
	ret := math.Sqrt(float64(n))
	fmt.Printf("math.Sqrt(%d) = %f\n", n, ret)

	f := 9.5
	ret = math.Ceil(f)
	fmt.Printf("math.Ceil(%f) = %f\n", f, ret)

	ret = math.Floor(f)
	fmt.Printf("math.Floor(%f) = %f\n", f, ret)

	fmt.Println("Basic Constant in math package")
	fmt.Printf("math.Pi: %f, math.E: %f\n", math.Pi, math.E)

}
