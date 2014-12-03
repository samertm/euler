package problem112

import (
	"fmt"

	"github.com/samertm/euler/math"
)

func isBouncy(x uint32) bool {
	if x < 100 {
		return false
	}
	var decreasing bool
	var increasing bool
	var next = x % 10
	for x = x / 10; x > 0; x = x / 10 {
		if next > x%10 {
			increasing = true
		} else if next < x%10 {
			decreasing = true
		}
		next = x % 10
		if increasing && decreasing {
			return true
		}
	}
	return increasing && decreasing
}

func Solve() string {
	var bouncy uint32
	max := uint32(3000000)
	var i uint32
	for i = 1; i < max; i++ {
		if isBouncy(i) {
			bouncy++
		}
		f := math.FractionSimplify(math.Fraction{N: bouncy, D: i})
		if f.N == 99 && f.D == 100 {
			return fmt.Sprint(i)
		}
	}
	return "nope"
}
