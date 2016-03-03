package main

import "fmt"

func main() {
	var m [21][21]int

	m[0][0] = 1

	for i := 0; i <= 20; i++ {
		for j := 0; j <= 20; j++ {
			if j < 20 {
				m[i][j+1] += m[i][j]
			}
			if i < 20 {
				m[i+1][j] += m[i][j]
			}
		}
	}

	fmt.Println(m[20][20])
}
