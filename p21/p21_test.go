package main

import "testing"

func TestIsAmicable(t *testing.T) {
	cases := []struct {
		in   int
		want bool
	}{
		{6, false},
		{220, true},
		{221, false},
	}

	for _, c := range cases {
		if got := isAmicable(c.in); got != c.want {
			t.Errorf("isAmicable(%d) should return: %t, got %t", c.in, c.want, got)
		}
	}
}
