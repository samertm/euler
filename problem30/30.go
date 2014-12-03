package problem30

import (
	"log"
	"math"
	"strconv"
)

func fifthEqual(x int) bool {
	var sum int
	for i := x; i > 0; i = i / 10 {
		sum += int(math.Pow(float64(i%10), 5))
		if sum > x {
			// Short-circuit.
			return false
		}
	}
	if sum == x {
		return true
	}
	return false
}

func Solve() string {
	max := 1000000
	var sum int
	for i := 3; i < max; i++ {
		if fifthEqual(i) {
			log.Println(i)
			sum += i
		}
	}
	return strconv.Itoa(sum)
}
