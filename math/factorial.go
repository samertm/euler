package math

var fact map[int]int

func init() {
	fact = make(map[int]int)
}

// Undefined for i <= 0.
func Factorial(i int) int {
	if i <= 1 {
		return 1
	}
	if v, ok := fact[i]; ok {
		return v
	}
	val := i * Factorial(i-1)
	fact[i] = val
	return val
}
