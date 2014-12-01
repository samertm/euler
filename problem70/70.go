package problem70

import (
	"strconv"

	"github.com/samertm/euler/math"
)

func Solve() string {
	phis := math.Phis(10000000)
	var minN int
	var minFn float64 = 1000.0
	for i := 2; i < len(phis); i++ {
		fn := float64(i) / float64(phis[i])
		if fn < minFn && math.IsPermutation(i, phis[i]) {
			minFn = fn
			minN = i
		}
	}
	return strconv.Itoa(minN)
}
