package main

import (
	"bufio"
	"fmt"
	"os"
)

func sumAdjacent(s string, n int, step int) int {
	sum := 0
	for i := 0; i < n; i++ {
		if s[i] == s[(i+step)%n] {
			sum += int(s[i]) - '0'
		}
	}
	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	n := len(s)

	// PART 1
	fmt.Println(sumAdjacent(s, n, 1))

	// PART 2
	fmt.Println(sumAdjacent(s, n, n/2))
}
