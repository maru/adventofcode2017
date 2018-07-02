package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func fewestSteps(count map[string]int) {
	for ok := true; ok; {
		ok = false
		// Cancel both
		for _, v := range [][]string{{"ne", "sw"}, {"se", "nw"}, {"s", "n"}} {
			d := min(count[v[0]], count[v[1]])
			if d > 0 {
				count[v[0]] -= d
				count[v[1]] -= d
				ok = true
			}
		}

		// Reduce to 1
		for _, v := range [][]string{{"ne", "s", "se"}, {"ne", "nw", "n"},
			{"se", "sw", "s"}, {"se", "n", "ne"}, {"s", "nw", "sw"}, {"sw", "n", "nw"}} {
			d := min(count[v[0]], count[v[1]])
			if d > 0 {
				count[v[0]] -= d
				count[v[1]] -= d
				count[v[2]] += d
				ok = true
			}
		}
	}
}

func countSteps(steps []string) int {
	count := make(map[string]int)
	for _, v := range steps {
		count[v]++
	}
	fewestSteps(count)

	total := int(0)
	for _, v := range count {
		if v < 0 {
			return -1
		}
		total += v
	}
	return total
}

func countMaxDist(steps []string) int {
	count := make(map[string]int)
	maxDist := 0
	for _, v := range steps {
		count[v]++
		fewestSteps(count)

		total := int(0)
		for _, v := range count {
			if v < 0 {
				return -1
			}
			total += v
		}
		maxDist = max(maxDist, total)
	}
	return maxDist
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	steps := strings.Split(s, ",")

	// Part 1
	fmt.Println(countSteps(steps))

	// Part 2
	fmt.Println(countMaxDist(steps))
}
