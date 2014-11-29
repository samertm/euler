package problem40

import "strconv"

// 1000000
func Solve() string {
	fract := make([]byte, 0, 1000000)
	for i := 1; len(fract) < 1000000; i++ {
		fract = append(fract, []byte(strconv.Itoa(i))...)
	}
	one, _ := strconv.Atoi(string(fract[0]))
	two, _ := strconv.Atoi(string(fract[9]))
	three, _ := strconv.Atoi(string(fract[99]))
	four, _ := strconv.Atoi(string(fract[999]))
	five, _ := strconv.Atoi(string(fract[9999]))
	six, _ := strconv.Atoi(string(fract[99999]))
	seven, _ := strconv.Atoi(string(fract[999999]))
	return strconv.Itoa(one * two * three * four * five * six * seven)
}
