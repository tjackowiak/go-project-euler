package main

/**
 * Let d(n) be defined as the sum of proper divisors of n (numbers less than n
 * which divide evenly into n).
 * If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and
 * each of a and b are called amicable numbers.
 *
 * For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110;
 * therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.
 *
 * Evaluate the sum of all the amicable numbers under 10000.
 */

func isAmicable(n int) bool {
	s := sumDivisors(n)
	return n != s && n == sumDivisors(s)
}

func sumDivisors(n int) int {
	sum := 1
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			if i*i == n {
				sum += i
			} else {
				sum += i + n/i
			}
		}
	}
	return sum
}

func main() {
	sum := 0
	for i := 1; i < 10000; i++ {
		if isAmicable(i) {
			sum += i
		}
	}
	println(sum)
}
