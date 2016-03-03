package main

import "testing"

/**
 * 3
 * 7 4
 * 2 4 6
 * 8 5 9 3
 */

func TestSumFromBottom(t *testing.T) {
	tab := [][]int{
		{3},
		{7, 4},
		{2, 4, 6},
		{8, 5, 9, 3},
	}

	if got := sumFromBottom(tab); got != 23 {
		t.Errorf("sum should give 23, got: %d", got)
	}
}

func TestSumFromTop(t *testing.T) {
	tab := [][]int{
		{3},
		{7, 4},
		{2, 4, 6},
		{8, 5, 9, 3},
	}

	if got := sumFromTop(tab); got != 23 {
		t.Errorf("sum should give 23, got: %d", got)
	}
}
