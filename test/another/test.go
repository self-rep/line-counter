package main

import (
	"log"
	"math"
	"time"
)

func primeNumbers(max int) []int {
	var primes []int

	for i := 2; i < max; i++ {
		isPrime := true

		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primes = append(primes, i)
		}
	}

	return primes
}

func main() {
	tm := time.Now()
	primeNumbers(100000)
	ts := time.Since(tm)
	log.Println("Took: ", ts.Milliseconds(), "ms To Get to 100000 prime numbers!") // if your wondering for me its 21ms
}
