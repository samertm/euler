package problem52

import "strconv"

func sameDigits(x, y int) bool {
	digits := make(map[int]int)
	for ; x != 0; x /= 10 {
		digits[x%10] += 1
	}
	for ; y != 0; y /= 10 {
		digits[y%10] -= 1
	}
	for i := 0; i < 9; i++ {
		if digits[i] != 0 {
			return false
		}
	}
	return true
}

func Solve() string {
	for i := 1; i < 1000000; i++ {
		if sameDigits(i, 2*i) &&
			sameDigits(i, 3*i) &&
			sameDigits(i, 4*i) &&
			sameDigits(i, 5*i) &&
			sameDigits(i, 6*i) {
			return strconv.Itoa(i)
		}
	}
	return ""
}
