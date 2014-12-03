package problem92

import "strconv"

func cycle(x uint64) uint64 {
	var sum uint64
	var i uint64
	for i = x; i > 0; i /= 10 {
		sum += (i % 10) * (i % 10)
	}
	return sum
}

var cache = make(map[uint64]bool)

func Solve() string {
	var count int
	var i uint64
	for i = 2; i < 10000000; i++ {
		num := i
		for num != 1 && num != 89 {
			num = cycle(num)
		}
		if num == 89 {
			count++
		}
	}
	return strconv.Itoa(count)
}
