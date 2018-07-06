package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	cmd string
	reg string
	val string
}

func nextInst(i Instruction, registers map[string]int, j *int) {
	switch i.cmd {
	case "set":
		if val, err := strconv.Atoi(i.val); err == nil {
			registers[i.reg] = int(val)
		} else {
			registers[i.reg] = registers[i.val]
		}
	case "sub":
		if val, err := strconv.Atoi(i.val); err == nil {
			registers[i.reg] -= int(val)
		} else {
			registers[i.reg] -= registers[i.val]
		}
	case "mul":
		if val, err := strconv.Atoi(i.val); err == nil {
			registers[i.reg] *= int(val)
		} else {
			registers[i.reg] *= registers[i.val]
		}
	case "jnz":
		var y int
		if val, err := strconv.Atoi(i.val); err == nil {
			y = int(val)
		} else {
			y = registers[i.val]
		}
		if val, err := strconv.Atoi(i.reg); err == nil && val != 0 {
			*j += y - 1
		} else if registers[i.reg] != 0 {
			*j += y - 1
		}
	}
}

func countMul(inst []Instruction) int {
	count := 0
	reg := make(map[string]int)

	for j := 0; j < len(inst); j++ {
		i := inst[j]
		nextInst(i, reg, &j)
		if i.cmd == "mul" {
			count++
		}
	}
	return count
}

func getRegisterH() int {
  // "assembler" to code
	h := 0
	b := 84*100 + 100000
	c := b + 17000
	for ; b <= c; b += 17 {
		for i := 2; i*i <= b; i++ {
			if b%i == 0 {
				h++
				break
			}
		}
	}
	return h
}

func main() {
	var inst []Instruction
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		strs := strings.Fields(s)
		var i Instruction
		i.cmd = strs[0]
		i.reg = strs[1]
		i.val = strs[2]
		inst = append(inst, i)
	}

	// Part 1
	fmt.Println(countMul(inst))

	// Part 2
	fmt.Println(getRegisterH())
}
