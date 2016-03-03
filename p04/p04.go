package main

import (
	"fmt"
	"strconv"
)

/**
 * A palindromic number reads the same both ways. The largest palindrome made
 * from the product of two 2-digit numbers is 9009 = 91 × 99.
 *
 * Find the largest palindrome made from the product of two 3-digit numbers.
 */

func isPalindrome(number int) bool {
	n := []rune(strconv.Itoa(number))

	for i, r := range n {
		if r != n[len(n)-i-1] {
			return false
		}
	}
	return true
}

func maxPalindrome(start, end int) int {

	max := 0
	for i := start; i <= end; i++ {
		for j := start; j <= end; j++ {
			test := i * j
			if isPalindrome(test) && test > max {
				max = test
			}
		}
	}

	return max
}

func main() {
	fmt.Println(maxPalindrome(100, 999))
}
