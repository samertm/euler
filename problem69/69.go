package problem69

import "fmt"

var phis [1000001]int

func Solve() string {
	for i := range phis {
		phis[i] = i - 1
	}
	var maxN int
	var maxFn float64
	for i := 2; i < 1000000; i++ {
		fn := float64(i) / float64(phis[i])
		if fn > maxFn {
			maxFn = fn
			maxN = i
		}
		if phis[i] == i-1 {
			for j := 2 * i; j < len(phis); j += i {
				phis[j] -= phis[j] / i
			}
		}
	}
	return fmt.Sprintln(maxN, maxFn, phis[maxN])
}
