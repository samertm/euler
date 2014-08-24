package primes

import (
	"log"
	"math"
)

var cache []int

func init() {
	cache = make([]int, 0)
}

func Gen(ps chan int) {
	update := func(i int) []int {
		ps <- i
		return append(cache, i)
	}
	cache = update(2)
	pp := 3
	cache = update(pp)
	for {
		pp += 2
		divisable := false
		sqrtpp := math.Sqrt(float64(pp))
		for _, a := range cache {
			if float64(a) > sqrtpp {
				break
			}
			if pp%a == 0 {
				divisable = true
			}
		}
		if !divisable {
			cache = update(pp)
		}
	}
}

func Contains(p int) bool {
	// TODO call gen for this condition
	if cache[len(cache)-1] < p {
		log.Fatal("Not enough prime numbers in cache")
	}
	low := 0
	high := len(cache) - 1
	for low <= high {
		mid := (low + high) / 2
		if cache[mid] == p {
			return true
		} else if cache[mid] > p {
			high = mid - 1
		} else { // cache[mid] < p
			low = mid + 1
		}
	}
	return false
}
