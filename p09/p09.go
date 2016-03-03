package main

import "fmt"

/**
 * A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
 * a^2 + b^2 = c^2
 *
 * For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
 *
 * There exists exactly one Pythagorean triplet for which a + b + c = 1000.
 * Find the product abc.
 */

func generateTriplets(limit int) {
	for a := 1; a < limit; a++ {
		for b := a; b < limit; b++ {
			c := limit - a - b
			if c <= b || c <= a {
				break
			}
			if a*a+b*b == c*c {
				fmt.Println(a, b, c, a*b*c)
			}
		}
	}
}

func main() {
	generateTriplets(1000)
}
