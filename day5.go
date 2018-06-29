package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func increment(a int) int {
	return a + 1
}

func decrease3(a int) int {
	if a < 3 {
		return a + 1
	}
	return a - 1
}

func countSteps(arr []int, modify func(a int) int) int {
	// Make a copy
	n := len(arr)
	a := make([]int, n)
	copy(a, arr)

	// Compute
	pos := 0
	steps := 0
	for 0 <= pos && pos < n {
		next := pos + a[pos]
		a[pos] = modify(a[pos])
		steps++
		pos = next
	}
	return steps
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var a []int
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		a = append(a, n)
	}

	// PART 1
	fmt.Println(countSteps(a, increment))

	// PART 2
	fmt.Println(countSteps(a, decrease3))
}
