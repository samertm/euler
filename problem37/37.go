package main

import (
	"fmt"
	"math"

	"github.com/samertm/euler/primes"
)

func euler37() int {
	t := make([]int, 0)
	ps := make(chan int, 10)
	go primes.Gen(ps)
	for p := <-ps; ; p = <-ps {
		if p == 2 || p == 3 || p == 5 || p == 7 {
			continue
		}
		if isTruncatable(p) {
			t = append(t, p)
		}
		if len(t) == 11 {
			fmt.Println(t)
			sum := 0
			for _, num := range t {
				sum += num
			}
			return sum
		}
	}
	return 0
}

func isTruncatable(p int) bool {
	// ABC -> AB -> A
	for i := p / 10; i > 0; i = i / 10 {
		if !primes.Contains(i) {
			return false
		}
	}
	// ABC -> BC -> C
	for i := p; i > 0; {
		if !primes.Contains(i) {
			return false
		}
		digits := int(math.Floor(math.Log10(float64(i))))
		scale := int(math.Pow(10, float64(digits)))
		// if scale == 10 {
		// 	break
		// }
		i -= (i / scale) * scale
	}
	return true
}
