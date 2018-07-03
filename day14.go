package main

import (
	"bufio"
	"fmt"
	"os"
	// "strings"
	// "strconv"
)

func knotHash(lengths []int) []int {
	lengths = append(lengths, []int{17, 31, 73, 47, 23}...)

	pos := 0
	skipSize := 0
	size := 256
	arr := make([]int, size)
	for j := 0; j < size; j++ {
		arr[j] = j
	}
	numRounds := 64
	for i := 0; i < numRounds; i++ {
		for _, len := range lengths {
			if len > size {
				continue
			}
			// Reverse
			for i := 0; i < len/2; i++ {
				j := (i + pos) % size
				k := (pos + len - 1 - i + size) % size
				arr[j], arr[k] = arr[k], arr[j]
			}
			// Move pos
			pos = (pos + len + skipSize) % size
			// Increase skipSize
			skipSize++
		}
	}
	sizeOut := 16
	out := make([]int, sizeOut)
	for i := 0; i < size; i++ {
		out[i/(size/sizeOut)] ^= arr[i]
	}
	return out
}

func countFree(s string, grid [][]bool) int {
	count := 0
	gridSize := 128
	for i := 0; i < gridSize; i++ {
		key := fmt.Sprintf("%s%s%d", s, "-", i)
		lengths := make([]int, len(key))
		for j := 0; j < len(key); j++ {
			lengths[j] = int(key[j])
		}

		out := knotHash(lengths)
		for k := 0; k < len(out); k++ {
			bin := fmt.Sprintf("%.8b", out[k])
			for j := 0; j < len(bin); j++ {
				if bin[j] == '1' {
					grid[i][j+k*8] = true
					count++
				}
			}
		}
	}
	return count
}

func dfs(g [][]bool, visited [][]bool, i, j int) {
	visited[i][j] = true
	next := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	for k := 0; k < 4; k++ {
		r := i + next[k][0]
		c := j + next[k][1]
		if r < 0 || c < 0 || r >= len(g) || c >= len(g) || visited[r][c] || !g[r][c] {
			continue
		}
		dfs(g, visited, r, c)
	}
}
func countRegions(g [][]bool) int {
	size := len(g)
	visited := make([][]bool, size)
	for i := 0; i < size; i++ {
		visited[i] = make([]bool, size)
	}
	count := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if visited[i][j] || !g[i][j] {
				continue
			}
			count++
			dfs(g, visited, i, j)
		}
	}
	return count
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()

	// Part 1
	gridSize := 128
	grid := make([][]bool, gridSize)
	for i := 0; i < gridSize; i++ {
		grid[i] = make([]bool, gridSize)
	}
	fmt.Println(countFree(s, grid))

	// Part 2
	fmt.Println(countRegions(grid))
}
