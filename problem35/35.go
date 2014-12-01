package problem35

import (
	"strconv"

	"github.com/samertm/euler/math"
)

func cycle(n int) chan int {
	origN := n
	cs := make(chan int)
	if n < 10 {
		close(cs)
		return cs
	}
	go func() {
		for {
			s := strconv.Itoa(n)
			n, _ = strconv.Atoi(s[len(s)-1:] + s[0:len(s)-1])
			if n == origN {
				close(cs)
				break
			} else {
				cs <- n
			}
		}
	}()
	return cs
}

func Solve() string {
	lim := 100
	primes := math.PrimesMap(lim)
	var count int
	for p := 2; p < lim; p++ {
		if !primes[p] {
			continue
		}
		cyclesArePrime := true
		for c := range cycle(p) {
			if !primes[c] {
				cyclesArePrime = false
				goto NextPrime
			}
		}
	NextPrime:
		if cyclesArePrime {
			count++
		}
	}
	return strconv.Itoa(count)
}
