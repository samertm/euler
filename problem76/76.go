package problem76

import "strconv"

type tuple struct{ x, y int }

var cache = make(map[tuple]int)

func cycle(x int, i int) (num int) {
	if v, ok := cache[tuple{x, i}]; ok {
		return v
	} else if x == 0 {
		return 1
	} else if x < 0 || i <= 0 {
		return 0
	}
	val := cycle(x-i, i) + cycle(x, i-1)
	cache[tuple{x, i}] = val
	return val
}

func Solve() string {
	return strconv.Itoa(cycle(100, 99))
}
