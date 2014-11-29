package problem94

import (
	"math"
	"math/big"
)

type triangle [3]float64

var numCycles int64 = 3

func cycle() (t triangle, valid bool) {
	base := float64(numCycles / 2)
	if numCycles%2 == 0 {
		t[0], t[1], t[2] = base, base, base+1
	} else {
		t[0], t[1], t[2] = base, base+1, base+1
	}
	numCycles++
	//if t[0]+t[1]+t[2] > 1000000000 {
	if t[0]+t[1]+t[2] > 1000000000 {
		return t, false
	}
	return t, true
}

func perimeter(t triangle) float64 {
	return t[0] + t[1] + t[2]
}

func areaIsIntegral(t triangle) bool {
	p := perimeter(t) / 2.0
	i, f := math.Modf(p * (p - t[0]) * (p - t[1]) * (p - t[2]))
	if f != 0 {
		return false
	}
	_, f = math.Modf(math.Sqrt(i))
	if f != 0 {
		return false
	}
	return true
}

func Solve() string {
	sum := big.NewInt(0)
	for t, valid := cycle(); valid; t, valid = cycle() {
		if areaIsIntegral(t) {
			sum.Add(sum, big.NewInt(int64(perimeter(t))))
		}
	}
	return sum.String()
}
