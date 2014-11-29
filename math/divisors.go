package math

func Divisors(num int) []int {
	if num < 0 {
		return make([]int, 0, 0)
	}
	divs := make([]int, 0, 0)
	half := num / 2
	for i := 1; i <= half; i++ {
		if num%i == 0 {
			divs = append(divs, i)
		}
	}
	return divs
}

func SumList(list []int) int {
	sum := 0
	for _, n := range list {
		sum += n
	}
	return sum
}

func ToSet(list []int) []int {
	hit := make(map[int]bool)
	set := make([]int, 0, len(list))
	for _, n := range list {
		if !hit[n] {
			set = append(set, n)
			hit[n] = true
		}
	}
	return set
}

func StringToSet(str string) []rune {
	hit := make(map[rune]bool)
	set := make([]rune, 0, len(str))
	for _, n := range str {
		if !hit[n] {
			set = append(set, n)
			hit[n] = true
		}
	}
	return set
}
