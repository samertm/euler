package problem71

import (
	"fmt"

	"github.com/cznic/mathutil"
)

type fraction struct {
	n   uint32
	d   uint32
	val *float64
}

func val(f *fraction) float64 {
	if f.val == nil {
		v := float64(f.n) / float64(f.d)
		f.val = &v
	}
	return *f.val
}

func Solve() string {
	upper := fraction{n: 3, d: 7}
	lower := fraction{n: 2, d: 7}
	var i uint32
	for i = 8; i <= 1000000; i++ {
		var j uint32
		for j = lower.n; j < i; j++ {
			if mathutil.GCDUint32(i, j) == 1 {
				f := fraction{n: j, d: i}
				if val(&f) < val(&lower) {
					continue
				}
				if val(&f) >= val(&upper) {
					break
				}
				lower = f
			}
		}
	}
	return fmt.Sprintf("%d", lower.n)
}
