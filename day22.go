package main

import (
	"bufio"
	"fmt"
	"os"
)

func printGrid(g [][]byte, x, y int) {
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			v := g[i][j]
			if i == x && y == j {
				fmt.Printf("%c]", v)
			} else {
				fmt.Printf("%c ", v)
			}
		}
		fmt.Println()
	}
}

func nextState1(b *byte, d *int) int {
	if *b == '#' {
		*b = '.'
		*d = (*d + 1) % 4
	} else {
		*b = '#'
		*d = (*d - 1 + 4) % 4
		return 1
	}
	return 0
}

func nextState2(b *byte, d *int) int {
	switch *b {
	case '.':
		*b = 'W'
		*d = (*d - 1 + 4) % 4
	case 'W':
		*b = '#'
		return 1
	case '#':
		*b = 'F'
		*d = (*d + 1) % 4
	case 'F':
		*b = '.'
		*d = (*d + 2) % 4
	}
	return 0
}

func countInfected(g [][]byte, steps int, nextState func(b *byte, d *int) int) int {
	n, m := len(g), len(g[0])

	offset := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	d := 0
	i, j := n/2, m/2
	count := 0
	for r := 0; r < steps; r++ {
		if i < 0 || j < 0 || i >= n || j >= m {
			println(r, "=", i, j, n, m)
			return -1
		}
		count += nextState(&g[i][j], &d)
		i += offset[d][0]
		j += offset[d][1]
	}
	// printGrid(g, i, j)
	return count
}

func prepareGrid(input []string, size int) [][]byte {
	grid := make([][]byte, size)
	for i := 0; i < size; i++ {
		grid[i] = make([]byte, size)
		for j := 0; j < size; j++ {
			grid[i][j] = '.'
		}
	}
	nn := len(input)
	i := (size - nn) / 2
	for _, s := range input {
		j := (size - nn) / 2
		for k := 0; k < len(s); k++ {
			grid[i][j+k] = s[k]
		}
		i++
	}
	return grid
}

func main() {
	var input []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		input = append(input, s)
	}

	// Part 1
	fmt.Println(countInfected(prepareGrid(input, 10001), 10000, nextState1))

	// Part 2
	fmt.Println(countInfected(prepareGrid(input, 101), 100, nextState2))
	fmt.Println(countInfected(prepareGrid(input, 501), 10000000, nextState2))
}
