package main

import "testing"

func TestP01(t *testing.T) {

	if res := p01(10); res != 23 {
		t.Error("p01(10) should be 23, got", res)
	}
}
