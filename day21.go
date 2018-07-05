package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func flip(a string) string {
	n := len(a)
	var bytes []byte
	if n == 4 {
		for i := 0; i < n; i += 2 {
			bytes = append(bytes, a[i+1])
			bytes = append(bytes, a[i])
		}
	} else {
		for i := 0; i < n; i += 3 {
			bytes = append(bytes, a[i+2])
			bytes = append(bytes, a[i+1])
			bytes = append(bytes, a[i])
		}
	}
	b := string(bytes[:n])
	return b
}

func rotate(a string) string {
	n := len(a)
	var bytes []byte

	if n == 4 {
		for i := 0; i < 2; i++ {
			bytes = append(bytes, a[i+2])
			bytes = append(bytes, a[i])
		}
	} else {
		for i := 0; i < 3; i++ {
			bytes = append(bytes, a[6+i])
			bytes = append(bytes, a[3+i])
			bytes = append(bytes, a[i])
		}
	}

	b := string(bytes[:n])
	return b
}

func printGrid(g [][]byte, size int) {
	fmt.Printf("\n")
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("%c", g[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func countOn(p string, size int, d map[string]string, steps int) int {
	g := make([][]byte, 3)
	for i := 0; i < 3; i++ {
		g[i] = make([]byte, 3)
	}

	// Init
	for i, j, k := 0, 0, 0; k < len(p); k++ {
		g[i][j] = p[k]
		j++
		if j == size {
			j = 0
			i++
		}
	}

	on := 0
	for r := 0; r < steps; r++ {
		var newlen, len int
		if size%2 == 0 {
			newlen, len = 3, 2
		} else {
			newlen, len = 4, 3
		}

		newsize := size / len * newlen
		tmp := make([][]byte, newsize)
		for i := 0; i < newsize; i++ {
			tmp[i] = make([]byte, newsize)
		}

		for i := 0; i < size/len; i++ {
			for j := 0; j < size/len; j++ {
				var s string
				for k := i * len; k < (i+1)*len; k++ {
					for l := j * len; l < (j+1)*len; l++ {
						s += string(g[k][l])
					}
				}

				m := d[s]

				for ii, k := 0, i*newlen; k < (i+1)*newlen; k++ {
					for l := j * newlen; l < (j+1)*newlen; l++ {
						tmp[k][l] = m[ii]
						ii++
					}
				}
			}
		}
		size = newsize
		g = make([][]byte, newsize)
		for i := 0; i < size; i++ {
			g[i] = make([]byte, newsize)
			for j := 0; j < size; j++ {
				g[i][j] = tmp[i][j]
			}
		}
		// printGrid(g, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if g[i][j] == '#' {
				on += 1
			}
		}
	}
	return on
}

func main() {
	dict := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		var a, b string
		fmt.Sscanf(s, "%s => %s", &a, &b)
		a = strings.Replace(a, "/", "", -1)
		b = strings.Replace(b, "/", "", -1)

		for i := 0; i < 4; i++ {
			dict[a] = b
			a = rotate(a)
		}
		a = flip(a)
		for i := 0; i < 4; i++ {
			dict[a] = b
			a = rotate(a)
		}
	}

	p := ".#./..#/###"
	p = strings.Replace(p, "/", "", -1)

	// Part 1
	fmt.Println(countOn(p, 3, dict, 5))

	// Part 2
	fmt.Println(countOn(p, 3, dict, 18))
}
