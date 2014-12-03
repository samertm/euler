package problem31

import "strconv"

var denoms = []int{200, 100, 50, 20, 10, 5, 2, 1}

func numCoins(x int, ds []int) int {
	if x == 0 {
		return 1
	} else if x < 0 || len(ds) == 0 {
		return 0
	}
	return numCoins(x-ds[0], ds) + numCoins(x, ds[1:])
}

func Solve() string {
	return strconv.Itoa(numCoins(200, denoms))
}
