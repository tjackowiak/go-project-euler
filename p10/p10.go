package main

import "fmt"

/**
 * The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
 *
 * Find the sum of all the primes below two million.
 */

func primeChan() chan int {
	res := make(chan int)

	go func() {
		res <- 2
		for i := 3; ; i += 2 {
			prime := true
			for j := 3; j*j <= i; j += 2 {
				if i%j == 0 {
					prime = false
					break
				}
			}

			if prime {
				res <- i
			}
		}
	}()

	return res
}

func main() {
	pc := primeChan()
	i, sum := 0, 0
	for ; i < 2000000; i, sum = <-pc, sum+i {
	}

	fmt.Println(sum)
}
