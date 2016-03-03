package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		number     int
		palindrome bool
	}{
		{1, true},
		{11, true},
		{121, true},
		{123, false},
		{9009, true},
	}

	for _, c := range cases {
		if got := isPalindrome(c.number); got != c.palindrome {
			t.Errorf("isPalindrome(%d) = %v, got %v", c.number, c.palindrome, got)
		}
	}
}

func TestMaxPalindrome(t *testing.T) {

	if got := maxPalindrome(10, 99); got != 9009 {
		t.Errorf("maxPalindrome(10, 99) = 9009, got %d", got)
	}
}
