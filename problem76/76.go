package problem76

import "log"

func cycle(x int, i int) (num int) {
	if x == 0 {
		return 1
	} else if x < 0 || i <= 0 {
		return 0
	}
	return cycle(x-i, i) + cycle(x, i-1)
}

func Solve() string {
	log.Fatal(cycle(100, 99))
	return ""
}
