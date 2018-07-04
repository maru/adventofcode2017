package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	offset = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	steps  int
)

func next(d []string, r, c, i, j int, dir int, letters *[]byte) {
	switch d[i][j] {
	case '|':
	case '-':
	case '+':
		// change direction
		for k := 0; k < 4; k++ {
			if k/2 == dir/2 {
				continue
			}
			newi := i + offset[k][0]
			newj := j + offset[k][1]
			if 0 <= newi && newi < r && 0 <= newj && newj < c && d[newi][newj] != ' ' {
				dir = k
				break
			}
		}
	case ' ':
		return
	default:
		*letters = append(*letters, d[i][j])
	}
	steps++
	next(d, r, c, i+offset[dir][0], j+offset[dir][1], dir, letters)
}

func follow(d []string) string {
	var letters []byte
	r, c := len(d), len(d[0])

	var i, j int
	for i, j = 0, 0; d[i][j] != '|'; j++ {
	}
	next(d, r, c, i, j, 0, &letters)
	return string(letters[:len(letters)])
}

func main() {
	var diagram []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		diagram = append(diagram, s)
	}

	// Part 1
	fmt.Println(follow(diagram))

	// Part 2
	fmt.Println(steps)
}
