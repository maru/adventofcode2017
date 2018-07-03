package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Operation struct {
	fn     int
	offset int
	p, q   uint
	c, d   byte
}

func spin(a uint64, pos []uint, offset int) uint64 {
	n := 16
	for i := 0; i < n; i++ {
		j := uint((i + offset) % n)
		c := (a >> (4 * uint(n-i-1))) & 0xF
		pos[c] = j
	}
	a = (a << uint(4*(n-offset))) + (a >> uint(4*offset))
	return a
}

func exchange(a uint64, pos []uint, p, q uint) uint64 {
	c := (a >> (4 * (16 - p - 1))) & 0xF
	d := (a >> (4 * (16 - q - 1))) & 0xF
	a = a ^ (c << (4 * (16 - p - 1))) ^ (d << (4 * (16 - p - 1)))
	a = a ^ (d << (4 * (16 - q - 1))) ^ (c << (4 * (16 - q - 1)))
	pos[c], pos[d] = q, p
	return a
}

func partner(a uint64, pos []uint, p, q byte) uint64 {
	i := p - byte('a')
	j := q - byte('a')
	return exchange(a, pos, pos[i], pos[j])
}

func printChars(a uint64) {
	for i := 15; i >= 0; i-- {
		m := int((a >> (4 * uint(i))) & 0xF)
		fmt.Printf("%c", m+int('a'))
	}
	fmt.Println()
}
func dance(ops []Operation, moves int, a uint64, pos []uint) uint64 {
	seen := make(map[uint64]uint64)
	seen[a] = 1
	seen[1] = a

	for i := 1; i <= moves; i++ {
		for _, op := range ops {
			if op.fn == 0 {
				a = spin(a, pos, op.offset)
			} else if op.fn == 1 {
				a = exchange(a, pos, op.p, op.q)
			} else {
				a = partner(a, pos, op.c, op.d)
			}
		}
		if seen[a] > 0 {
			period := uint64(i+1) - seen[a]
			f := uint64(moves) % period
			a = seen[f]
			break
		} else {
			seen[a] = uint64(i + 1)
			seen[uint64(i+1)] = a
		}
	}
	printChars(a)
	return a
}

func parseOperations(strs []string) []Operation {
	var ops []Operation
	for _, s := range strs {
		var op Operation
		if s[0] == 's' {
			op.fn = 0
			fmt.Sscanf(s[1:], "%d", &op.offset)
		} else if s[0] == 'x' {
			op.fn = 1
			fmt.Sscanf(s[1:], "%d/%d", &op.p, &op.q)
		} else {
			op.fn = 2
			fmt.Sscanf(s[1:], "%c/%c", &op.c, &op.d)
		}
		ops = append(ops, op)
	}
	return ops
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	strs := strings.Split(s, ",")
	ops := parseOperations(strs)

	// Create array
	var a uint64
	n := 16
	a = 0x0123456789abcdef
	pos := make([]uint, n)
	for i := 0; i < n; i++ {
		pos[i] = uint(i)
	}

	// Part 1
	a = dance(ops, 1, a, pos)

	// Part 2
	dance(ops, 1000000000, a, pos)
}
