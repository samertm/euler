package problem36

import "strconv"

func isPalindrome(s string) bool {
	for len(s) > 1 {
		if s[0] != s[len(s)-1] {
			return false
		}
		s = s[1 : len(s)-1]
	}
	return true
}

func Solve() string {
	var sum int
	for i := 1; i < 1000000; i++ {
		if isPalindrome(strconv.Itoa(i)) && isPalindrome(strconv.FormatInt(int64(i), 2)) {
			sum += i
		}
	}
	return strconv.Itoa(sum)
}
