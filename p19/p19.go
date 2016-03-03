package main

import (
	"fmt"
	"time"
)

func main() {
	sundays := 0
	for i := 1901; i <= 2000; i++ {
		for j := 1; j <= 12; j++ {
			t := time.Date(i, time.Month(j), 1, 1, 0, 0, 0, time.UTC)
			if t.Weekday() == time.Sunday {
				sundays++
			}
		}
	}

	fmt.Println(sundays)
}
