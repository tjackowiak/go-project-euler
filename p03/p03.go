package main

import "fmt"

/**
 * The prime factors of 13195 are 5, 7, 13 and 29.
 *
 * What is the largest prime factor of the number 600851475143 ?
 */

func primeChan() chan int {
	var in chan int
	out := make(chan int)
	res := make(chan int)

	sieveX := func(x int, in, out chan int) {
		for {
			i := <-in
			if i%x != 0 {
				out <- i
			}
		}
	}

	// int generator
	go func(out chan int) {
		for i := 2; ; i++ {
			out <- i
		}
	}(out)

	go func() {
		for {
			prime := <-out
			res <- prime
			in = out
			out = make(chan int)

			go sieveX(prime, in, out)
		}
	}()

	return res
}

func p03(maxPrimeOf int) int {
	pc := primeChan()

	maxPrime := 0
	for i := <-pc; i <= maxPrimeOf; i = <-pc {
		if maxPrimeOf%i == 0 {
			maxPrime = i
			maxPrimeOf = maxPrimeOf / i
		}
	}

	return maxPrime
}

func main() {
	fmt.Println(p03(600851475143))
}
