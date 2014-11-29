package problem71

import (
	"fmt"
	"sort"

	"github.com/cznic/mathutil"
)

type fraction struct {
	n   uint32
	d   uint32
	val *float64
}

type fractions []fraction

func (f fractions) Len() int {
	return len(f)
}

func (f fractions) Less(i, j int) bool {
	return val(f[i]) < val(f[j])
}

func (f fractions) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func val(f fraction) float64 {
	if f.val == nil {
		v := float64(f.n) / float64(f.d)
		f.val = &v
	}
	return *f.val
}

func Solve() string {
	fracts := make(fractions, 0, 0)
	var mil uint32 = 100000
	target := fraction{n: 3, d: 7}
	var i uint32
	for i = 1; i <= mil; i++ {
		var j uint32
		for j = 1; j < i; j++ {
			if mathutil.GCDUint32(i, j) == 1 {
				f := fraction{n: j, d: i}
				if val(f) > val(target) {
					break
				}
				fracts = append(fracts, fraction{n: j, d: i})
			}
		}
	}
	sort.Sort(fracts)
	for i, n := range fracts {
		if n.n == 3 && n.d == 7 {
			return fmt.Sprintf("%v", fracts[i-1])
		}
	}
	panic("unreachable")
	return "nope"
}
