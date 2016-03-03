package main

import (
	"math/big"
	"strconv"
)

func main() {
	var z big.Int
	z.SetInt64(2)

	for i := 1; i < 1000; i++ {
		z.Add(&z, &z)
	}

	sum := 0
	for _, num := range z.String() {
		val, _ := strconv.Atoi(string(num))
		sum += val
	}
	println(sum)
}
