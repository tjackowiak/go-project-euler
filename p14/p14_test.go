package main

import "testing"

func TestFP14(t *testing.T) {

	cases := []struct {
		in, want uint64
	}{
		{13, 40},
		{40, 20},
		{20, 10},
		{10, 5},
		{5, 16},
	}

	for _, c := range cases {
		got := fP14(c.in)
		if got != c.want {
			t.Errorf("fP14(%d) should be %d, got %d", c.in, c.want, got)
		}
	}
}

func TestFP14Length(t *testing.T) {

	cases := []struct {
		in   uint64
		want int
	}{
		{13, 10},
		{40, 9},
		{80, 10},
		{80, 10},
	}

	for _, c := range cases {
		got := fP14len(c.in)
		if got != c.want {
			t.Errorf("fP14len(%d) should be %d, got %d", c.in, c.want, got)
		}
	}

}
