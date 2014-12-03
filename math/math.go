package math

import (
	"math"
	"math/big"
)

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

func PrimesMap(limit int) map[int]bool {
	nums := make(map[int]bool)
	for i := 2; i < limit; i++ {
		nums[i] = true
	}
	for i := 2; i < limit; i++ {
		if nums[i] {
			for j := 2 * i; j < limit; j += i {
				nums[j] = false
			}
		}
	}
	return nums
}

func Primes(limit int) []int {
	nums := make([]bool, limit)
	for i := range nums {
		nums[i] = true
	}
	primes := make([]int, 0)
	for i := 2; i < limit; i++ {
		if nums[i] {
			primes = append(primes, i)
			for j := 2 * i; j < len(nums); j += i {
				nums[j] = false
			}
		}
	}
	return primes
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

func IsPalindrome(s string) bool {
	for len(s) > 1 {
		if s[0] != s[len(s)-1] {
			return false
		}
		s = s[1 : len(s)-1]
	}
	return true
}

func Reverse(x int) int {
	var reverse int
	p := int(math.Log10(float64(x)))
	for ; x > 0; x /= 10 {
		reverse += int(math.Pow10(p)) * (x % 10)
		p--
	}
	return reverse
}

func ReverseUint64(x uint64) uint64 {
	var reverse uint64
	p := int(math.Log10(float64(x)))
	for ; x > 0; x /= 10 {
		reverse += uint64(math.Pow10(p)) * (x % 10)
		p--
	}
	return reverse
}

func ReverseBig(x *big.Int) *big.Int {
	reverse := big.NewInt(0)
	p := len(x.String()) - 1
	for ; x.Uint64() != 0; x.Div(x, big.NewInt(10)) {
		reverse.Mul(big.NewInt(int64(math.Pow10(p))), big.NewInt(0).Mod(x, big.NewInt(10)))
		p--
	}
	return reverse
}
