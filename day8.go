package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	register   string
	value      int
	registerIf string
	compIf     string
	valueIf    int
}

func getMaxValue(reg map[string]int) int {
	maxValue := -123456789
	for _, v := range reg {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}

func largestValue(instr []Instruction) (int, int) {
	reg := make(map[string]int)
	highest := -123456789
	for _, i := range instr {
		if _, ok := reg[i.registerIf]; !ok {
			reg[i.registerIf] = 0
		}
		if _, ok := reg[i.register]; !ok {
			reg[i.register] = 0
		}
		ok := false
		switch i.compIf {
		case ">":
			if reg[i.registerIf] > i.valueIf {
				ok = true
			}
		case ">=":
			if reg[i.registerIf] >= i.valueIf {
				ok = true
			}
		case "<":
			if reg[i.registerIf] < i.valueIf {
				ok = true
			}
		case "<=":
			if reg[i.registerIf] <= i.valueIf {
				ok = true
			}
		case "==":
			if reg[i.registerIf] == i.valueIf {
				ok = true
			}
		case "!=":
			if reg[i.registerIf] != i.valueIf {
				ok = true
			}
		}
		if ok {
			reg[i.register] += i.value
		}
		if h := getMaxValue(reg); h > highest {
			highest = h
		}
	}
	return getMaxValue(reg), highest
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var instr []Instruction
	for scanner.Scan() {
		s := scanner.Text()
		strs := strings.Fields(s)
		i := Instruction{strs[0], 0, strs[4], strs[5], 0}
		i.value, _ = strconv.Atoi(strs[2])
		i.valueIf, _ = strconv.Atoi(strs[6])
		if strs[1] == "dec" {
			i.value *= -1
		}
		instr = append(instr, i)
	}

	// PART 1
	l, h := largestValue(instr)
	fmt.Println(l)

	// PART 2
	fmt.Println(h)
}
