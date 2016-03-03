package main

import "testing"

func TestSmallestDividedByAll(t *testing.T) {

	if got := smallestDividedByAll(10); got != 2520 {
		t.Error("Wanted 2520, got:", got)
	}
}

func TestSmallestDividedByAllLCM(t *testing.T) {

	if got := smallestDividedByAllLCM(10); got != 2520 {
		t.Error("Wanted 2520, got:", got)
	}
}

func TestSmallestDividedByAllLCMBig(t *testing.T) {

	if got := smallestDividedByAllLCMBig(10); got != 2520 {
		t.Error("Wanted 2520, got:", got)
	}
}
