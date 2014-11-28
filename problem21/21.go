package problem21

import (
	"github.com/samertm/euler/math"
	"strconv"
)

func Solve() string {
	amicable := make([]int, 0, 0)
	for a := 0; a < 10000; a++ {
		b := math.SumList(math.Divisors(a))
		potentialA := math.SumList(math.Divisors(b))
		if a == potentialA && a != b {
			amicable = append(amicable, a)
		}
	}
	return strconv.Itoa(math.SumList(math.ToSet(amicable)))
}
