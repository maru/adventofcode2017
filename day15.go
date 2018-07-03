package main

import (
	"fmt"
)

func judge(a, b int64) bool {
	if (a & 0xFFFF) == (b & 0xFFFF) {
		return true
	}
	return false
}

func generator(prev int64, factor int64, mult int64) int64 {
	for {
		r := (prev * factor) % int64(2147483647)
		if (r % mult) == 0 {
			return r
		}
		prev = r
	}
}

func countMatches(maxRounds int, multA, multB int64) int64 {
	var startA, startB, factorA, factorB int64
	var count int64

	// Test
	// startA, factorA = 65, 16807
	// startB, factorB = 8921, 48271

	// Input
	startA, factorA = 512, 16807
	startB, factorB = 191, 48271

	count = 0
	for i := 0; i < maxRounds; i++ {
		startA = generator(startA, factorA, multA)
		startB = generator(startB, factorB, multB)
		if judge(startA, startB) {
			count++
		}
	}
	return count
}

func main() {
	// Part 1
	fmt.Println(countMatches(40000000, 1, 1))

	// Part 2
	fmt.Println(countMatches(5000000, 4, 8))
}
