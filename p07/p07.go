package main

import "fmt"

/**
 * By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13,
 * we can see that the 6th prime is 13.
 *
 * What is the 10 001st prime number?
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

func main() {
	pc := primeChan()

	for i := 0; i < 10000; i++ {
		<-pc
	}

	fmt.Println(<-pc)
}
