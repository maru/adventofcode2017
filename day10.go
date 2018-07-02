package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func knotHash(arr []int, size int, lengths []int, pos *int, skipSize *int) {
	for _, len := range lengths {
		if len > size {
			continue
		}
		// Reverse
		for i := 0; i < len/2; i++ {
			j := (i + *pos) % size
			k := (*pos + len - 1 - i + size) % size
			arr[j], arr[k] = arr[k], arr[j]
		}
		// Move pos
		*pos = (*pos + len + *skipSize) % size
		// Increase skipSize
		*skipSize++
	}
}

func reduceHash(in []int, sizeIn int, out []int, sizeOut int) {
	for i := 0; i < sizeIn; i++ {
		out[i/(sizeIn/sizeOut)] ^= in[i]
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()

		// Part 1
		{
			size := 256
			arr := make([]int, size)
			for i := 0; i < size; i++ {
				arr[i] = i
			}

			var lengths []int
			for _, v := range strings.Split(s, ",") {
				i, _ := strconv.Atoi(v)
				lengths = append(lengths, i)
			}
			pos := 0
			skipSize := 0
			knotHash(arr, size, lengths, &pos, &skipSize)
			fmt.Println(arr[0], "*", arr[1], "=", arr[0]*arr[1])
		}

		// Part 2
		{
			size := 256
			arr := make([]int, size)
			for i := 0; i < size; i++ {
				arr[i] = i
			}

			var lengths []int
			for _, v := range strings.Split(s, "") {
				i := []byte(v)
				lengths = append(lengths, int(i[0]))
			}
			lengths = append(lengths, []int{17, 31, 73, 47, 23}...)
			numRounds := 64
			pos := 0
			skipSize := 0
			for i := 0; i < numRounds; i++ {
				knotHash(arr, size, lengths, &pos, &skipSize)
			}
			sizeDH := 16
			denseHash := make([]int, sizeDH)
			reduceHash(arr, size, denseHash, sizeDH)

			for _, v := range denseHash {
				fmt.Printf("%02x", v)
			}
			fmt.Println()
		}
	}
}
