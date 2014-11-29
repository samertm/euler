package problem57

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/samertm/euler/math"
)

func Solve() string {
	exp := math.FractionBig{N: *big.NewInt(3), D: *big.NewInt(2)}
	bigOne := math.FractionBig{N: *big.NewInt(1), D: *big.NewInt(1)}
	var count int
	for i := 1; i < 1000; i++ {
		// First, we add one to exp.
		exp = math.FractionAddBig(bigOne, exp)
		// Then we get the reciprocal.
		exp = math.FractionBig{N: exp.D, D: exp.N}
		// Then we add one again.
		exp = math.FractionAddBig(bigOne, exp)
		if len((&exp.N).String()) > len((&exp.D).String()) {
			count++
			fmt.Println(exp)
		}
	}
	return strconv.Itoa(count)
}
