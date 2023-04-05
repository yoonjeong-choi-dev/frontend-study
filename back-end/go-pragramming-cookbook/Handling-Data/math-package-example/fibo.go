package main

import "math/big"

var cache map[int]*big.Int

func init() {
	cache = make(map[int]*big.Int)
}

func Fib(n int) *big.Int {
	if n < 0 {
		return big.NewInt(1)
	}

	if n < 2 {
		cache[n] = big.NewInt(1)
	}

	// check that the value is already calculated
	if val, ok := cache[n]; ok {
		return val
	}

	cache[n] = big.NewInt(0)
	cache[n].Add(cache[n], Fib(n-1))
	cache[n].Add(cache[n], Fib(n-2))
	return cache[n]
}
