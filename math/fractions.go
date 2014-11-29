package math

import "github.com/cznic/mathutil"

type Fraction struct {
	N   uint32
	D   uint32
	Val *float64
}

func FractionVal(f *Fraction) float64 {
	if f.Val == nil {
		v := float64(f.N) / float64(f.D)
		f.Val = &v
	}
	return *f.Val
}

func FractionEqual(f *Fraction, s *Fraction) bool {
	return FractionVal(f) == FractionVal(s)
}

func FractionSimplify(f Fraction) Fraction {
	for gcd := mathutil.GCDUint32(f.N, f.D); gcd != 1; gcd = mathutil.GCDUint32(f.N, f.D) {
		f.N = f.N / gcd
		f.D = f.D / gcd
	}
	return f
}

func FractionMultiply(f Fraction, s Fraction) Fraction {
	return FractionSimplify(Fraction{N: f.N * s.N, D: f.D * s.D})
}

func FractionMultiplyList(fs []Fraction) Fraction {
	prod := Fraction{N: 1, D: 1}
	for _, f := range fs {
		prod = FractionMultiply(prod, f)
	}
	return prod
}
