package main

func fP14(x uint64) uint64 {
	if x%2 == 0 {
		return x / 2
	}
	return 3*x + 1
}

func fP14len(x uint64) int {

	var len int
	for len = 1; x != 1; {
		l, ok := m[x]
		if ok == true {
			return l + len
		}

		x = fP14(x)
		len++
	}
	return len
}

var m map[uint64]int

func init() {
	m = make(map[uint64]int)
}

func main() {
	max := 0
	for i := 1; i < 1000000; i++ {
		if len := fP14len(uint64(i)); len > max {
			max = i
		}
	}

	println(max, fP14len(uint64(max)))
}
