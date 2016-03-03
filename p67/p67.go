package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func max(nums []int) int {
	max := nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}

func sumFromBottom(t [][]int) int {

	for i := len(t) - 2; i >= 0; i-- {
		row := t[i+1]

		for j, _ := range t[i] {
			t[i][j] += max(row[:2])
			row = row[1:]
		}
	}

	return t[0][0]
}

func main() {
	triangle := make([][]int, 0)

	file, err := os.Open("triangle.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		triangle = append(triangle, make([]int, i+1))
		for j, number := range strings.Split(scanner.Text(), " ") {
			triangle[i][j], _ = strconv.Atoi(number)
		}
		i++
	}

	fmt.Println(sumFromBottom(triangle))
}
