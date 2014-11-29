package problem33

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/samertm/euler/math"
)

func FakeSimplify(f math.Fraction) (fr math.Fraction, valid bool) {
	n := strconv.Itoa(int(f.N))
	d := strconv.Itoa(int(f.D))
	if n[1] == '0' && d[1] == '0' {
		return math.Fraction{}, false
	}
	for _, t := range math.StringToSet(n + d) {
		newN := strings.Replace(n, string(t), "", -1)
		if len(newN) != 1 {
			continue
		}
		newD := strings.Replace(d, string(t), "", -1)
		if len(newD) != 1 {
			continue
		}
		numN, _ := strconv.Atoi(newN)
		numD, _ := strconv.Atoi(newD)
		if numD == 0 {
			continue
		}
		if math.FractionEqual(&f, &math.Fraction{N: uint32(numN), D: uint32(numD)}) {
			return f, true
		}
	}
	return math.Fraction{}, false
}

func Solve() string {
	nontriv := []math.Fraction{}
	for denom := 10; denom < 100; denom++ {
		for numer := 10; numer < denom; numer++ {
			if f, valid := FakeSimplify(math.Fraction{N: uint32(numer), D: uint32(denom)}); valid {
				nontriv = append(nontriv, f)
			}
		}
	}
	prod := math.FractionMultiplyList(nontriv)
	return fmt.Sprintf("%d", prod.D)
}
