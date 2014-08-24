package main

import (
	"fmt"

	"github.com/samertm/euler/primes"
)

func main() {
	ch := make(chan int, 10)
	go primes.Gen(ch)
	for <-ch < 1000 {

	}
	for i := 0; i < 100; i++ {
		fmt.Printf("%d is prime: %t\n", i, primes.Contains(i))
	}
}
