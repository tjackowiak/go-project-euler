package main

import "testing"

func TestPrimeChan(t *testing.T) {
	cases := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}

	fc := primeChan()

	for i, c := range cases {
		if got := <-fc; got != c {
			t.Errorf("%d prime from primeChan() should be %d, got %d", i+1, c, got)
		}
	}
}

func TestP03(t *testing.T) {

	if res := p03(13195); res != 29 {
		t.Errorf("p03(13195)) should be 29, got %d", res)
	}
}
