package main

import "testing"

func TestFib(t *testing.T) {

	cases := []struct {
		in, want int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{10, 89},
	}

	for _, c := range cases {
		got := fib(c.in)
		if got != c.want {
			t.Errorf("fib(%d) should be %d, got %d", c.in, c.want, got)
		}
	}
}

func TestFibChan(t *testing.T) {
	cases := []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89}

	fc := fibChan()

	for i, c := range cases {
		if got := <-fc; got != c {
			t.Errorf("fib(%d) should be %d, got %d", i+1, c, got)
		}
	}

}

func TestP02(t *testing.T) {

	if res := p02(100); res != (2 + 8 + 34) {
		t.Errorf("p02(10) should be %d, got %d", (2 + 8 + 34), res)
	}
}
