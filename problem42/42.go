package problem42

import (
	"strconv"

	"github.com/samertm/euler/csv"
	"github.com/samertm/euler/math"
)

var offset rune = 'A' - 1

func wordSum(word string) int {
	var sum int
	for _, c := range word {
		sum += int(c - offset)
	}
	return sum
}

func Solve() string {
	contents := csv.Read("problem42/p042_words.txt")
	var triangleWords int
	for _, word := range contents {
		if math.IsTriangleNum(wordSum(word)) {
			triangleWords++
		}
	}
	return strconv.Itoa(triangleWords)
}
