package math

import (
	"math/big"

	"github.com/cznic/mathutil"
)

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

func FractionAdd(f Fraction, s Fraction) Fraction {
	return FractionSimplify(Fraction{N: f.N*s.D + s.N*f.D, D: f.D * s.D})
}

type Fraction64 struct {
	N uint64
	D uint64
}

func FractionToSet64(fs []Fraction64) []Fraction64 {
	hit := make(map[Fraction64]bool)
	set := make([]Fraction64, 0, len(fs))
	for _, n := range fs {
		if !hit[n] {
			set = append(set, n)
			hit[n] = true
		}
	}
	return set
}

func FractionAdd64(f Fraction64, s Fraction64) Fraction64 {
	return FractionSimplify64(Fraction64{N: f.N*s.D + s.N*f.D, D: f.D * s.D})
}

func FractionSimplify64(f Fraction64) Fraction64 {
	for gcd := mathutil.GCDUint64(f.N, f.D); gcd != 1; gcd = mathutil.GCDUint64(f.N, f.D) {
		f.N = f.N / gcd
		f.D = f.D / gcd
	}
	return f
}

type FractionBig struct {
	N big.Int
	D big.Int
}

func FractionAddBig(f FractionBig, s FractionBig) FractionBig {
	return FractionSimplifyBig(FractionBig{
		N: *big.NewInt(0).Add(big.NewInt(0).Mul(&f.N, &s.D), big.NewInt(0).Mul(&s.N, &f.D)),
		D: *big.NewInt(0).Mul(&f.D, &s.D),
	})
}

func FractionSimplifyBig(f FractionBig) FractionBig {
	n := &f.N
	d := &f.D
	for gcd := big.NewInt(0).GCD(nil, nil, n, d); gcd.Cmp(big.NewInt(1)) != 0; gcd = gcd.GCD(nil, nil, n, d) {
		n.Div(n, gcd)
		d.Div(d, gcd)
	}
	f.N = *n
	f.D = *d
	return f
}
