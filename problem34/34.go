package problem34

import (
	"log"
	"strconv"

	"github.com/samertm/euler/math"
)

func factEqual(x int) bool {
	var sum int
	for i := x; i > 0; i = i / 10 {
		sum += math.Factorial(i % 10)
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
	max := 100000
	var sum int
	for i := 3; i < max; i++ {
		if factEqual(i) {
			log.Println(i)
			sum += i
		}
	}
	return strconv.Itoa(sum)
}
