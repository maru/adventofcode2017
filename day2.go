package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func checkSum(m [][]int) int {
	sum := 0
	for i, _ := range m {
		minNum := 100000000
		maxNum := 0
		for j, _ := range m[i] {
			minNum = min(minNum, m[i][j])
			maxNum = max(maxNum, m[i][j])
		}
		sum += maxNum - minNum
	}
	return sum
}

func checkSumDiv(m [][]int) int {
	sum := 0
	for _, v := range m {
		n := len(v)
		div := -1
		for i := 0; div < 0 && i < n; i++ {
			for j := 0; div < 0 && j < n; j++ {
				if i == j {
					continue
				}
				if v[i]%v[j] == 0 {
					div = v[i] / v[j]
				}
			}
		}
		sum += div
	}
	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var m [][]int
	for scanner.Scan() {
		s := scanner.Text()
		strs := strings.Fields(s)
		v := make([]int, len(strs))
		for i, num := range strs {
			v[i], _ = strconv.Atoi(num)
		}
		m = append(m, v)
	}

	// PART 1
	fmt.Println(checkSum(m))

	// PART 2
	fmt.Println(checkSumDiv(m))

}
