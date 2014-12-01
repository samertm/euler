package math

func IsPermutation(x int, y int) bool {
	xDigits := make(map[int]int)
	for x != 0 {
		xDigits[x%10] += 1
		x /= 10
	}
	yDigits := make(map[int]int)
	for y != 0 {
		yDigits[y%10] += 1
		y /= 10
	}
	for i := 0; i < 10; i++ {
		if xDigits[i] != yDigits[i] {
			return false
		}
	}
	return true
}

func Phis(limit int) []int {
	phis := make([]int, limit)
	for i := range phis {
		phis[i] = i - 1
	}
	for i := 2; i < limit; i++ {
		if phis[i] == i-1 {
			for j := 2 * i; j < len(phis); j += i {
				phis[j] -= phis[j] / i
			}
		}
	}
	return phis
}
