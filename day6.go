package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toString(a []byte, n int) string {
	s := string(a[:n])
	return s
}

func maxSlice(a []byte, n int) int {
	maxElem := byte(0)
	maxPos := 0
	for i, e := range a {
		if e > maxElem {
			maxElem = e
			maxPos = i
		}
	}
	return maxPos
}

func reallocation(a []byte, n int) {
	if n < 1 {
		return
	}
	maxPos := maxSlice(a, n)
	maxCount := a[maxPos]
	a[maxPos] = 0
	for i := (maxPos + 1) % n; maxCount > 0; maxCount-- {
		a[i]++
		i = (i + 1) % n
	}
}

func countSteps(a []byte) int {
	n := len(a)

	// Compute
	m := make(map[string]bool)
	steps := 0
	for {
		s := toString(a, n)
		ok, _ := m[s]
		if ok {
			break
		}
		m[s] = true
		steps++
		reallocation(a, n)
	}
	return steps
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	strs := strings.Fields(scanner.Text())
	v := make([]byte, len(strs))
	for i, num := range strs {
		b, _ := strconv.Atoi(num)
		v[i] = byte(b)
	}

	// PART 1
	fmt.Println(countSteps(v))

	// PART 1
	fmt.Println(countSteps(v))
}
