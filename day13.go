package main

import (
	"bufio"
	"fmt"
	"os"
	// "strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func severity(m map[int]int, maxLayer int, delay int) (int, bool) {
	count := 0
	ok := true
	for i := 0; i < maxLayer; i++ {
		if m[i] == 0 {
			continue
		} else if m[i] == 1 {
			count += i
			ok = false
		} else if (delay+i)%((m[i]-1)*2) == 0 {
			count += i * m[i]
			ok = false
		}
		if delay > 0 && !ok {
			return -1, ok
		}
	}
	return count, ok
}

func main() {
	layers := make(map[int]int)
	maxLayer := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		var v, d int
		fmt.Sscanf(s, "%d: %d", &v, &d)
		layers[v] = d
		maxLayer = max(maxLayer, v) + 1
	}

	// Part 1
	sev, _ := severity(layers, maxLayer, 0)
	fmt.Println(sev)

	// Part 2
	for d := 0; ; d++ {
		sev, ok := severity(layers, maxLayer, d)
		if sev == 0 && ok {
			fmt.Println(d, sev)
			break
		}
	}
}
