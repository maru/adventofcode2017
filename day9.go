package main

import (
	"bufio"
	"fmt"
	"os"
)

func getScoreAndGarbage(s string, n int, i *int, score int) (int, int) {
	sum := 0
	g := 0
	for *i < n {
		if s[*i] == '{' {
			*i++
			ss, gg := getScoreAndGarbage(s, n, i, score+1)
			sum += ss
			g += gg
		} else if s[*i] == '}' {
			*i++
			return score + sum, g
		} else if s[*i] == '<' {
			*i++
			for *i < n {
				if s[*i] == '>' {
					*i++
					break
				} else if s[*i] == '!' {
					*i += 2
				} else {
					*i++
					g++
				}
			}
		} else {
			*i++
		}
	}
	return score + sum, g
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()

		// Part 1
		pos := 0
		score, garbage := getScoreAndGarbage(s, len(s), &pos, 0)
		fmt.Println(score)

		// Part 2
		fmt.Println(garbage)
	}
}
