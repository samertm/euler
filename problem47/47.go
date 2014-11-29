package problem47

import (
	"strconv"

	"github.com/cznic/mathutil"
)

func Solve() string {
	var i uint32
	var times int
	for i = 0; true; i++ {
		fs := mathutil.FactorInt(i)
		if len(fs) == 4 {
			times++
		} else {
			times = 0
		}
		if times == 4 {
			return strconv.Itoa(int(i) - 3)
		}
	}
	panic("unreachable")
	return ""
}
