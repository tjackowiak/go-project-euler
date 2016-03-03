package main

import (
	"fmt"
	"math"
	"math/big"
)

/**
 * 2520 is the smallest number that can be divided by each of the numbers from 1 to 10
 * without any remainder.
 *
 * What is the smallest positive number that is evenly divisible by all of the numbers
 * from 1 to 20?
 */

func lcm(m, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	r := (m * n) / gcd(m, n)
	return int(math.Abs(float64(r)))
}

func gcd(a, b int) int {
	for b != 0 {
		b, a = a%b, b
	}

	return a
}

func smallestDividedByAll(upper int) int {

	num := upper
	for ; ; num++ {
		divisible := true
		for i := 1; i <= upper; i++ {
			if num%i != 0 {
				divisible = false
				break
			}
		}
		if divisible {
			return num
		}
	}
	return 0
}

func smallestDividedByAllLCMBig(upper int) int64 {

	var z, x, y big.Int
	z.SetInt64(1)

	for i := 2; i <= upper; i++ {
		x.SetInt64(int64(i))
		y.Set(&z)
		z.Mul(z.Div(&x, z.GCD(nil, nil, &x, &y)), &y)
	}

	return z.Int64()
}

func smallestDividedByAllLCM(upper int) int {

	res := 1
	for i := 2; i <= upper; i++ {
		res = lcm(res, i)
	}

	return res
}

func main() {
	// fmt.Println(smallestDividedByAll(20))
	fmt.Println(smallestDividedByAllLCM(20))
}
