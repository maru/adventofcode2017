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

func nextInst(i Instruction, freq *int, registers map[string]int, j *int) {
	switch i.cmd {
	case "snd":
		*freq = registers[i.reg]
	case "set":
		if val, err := strconv.Atoi(i.val); err == nil {
			registers[i.reg] = int(val)
		} else {
			registers[i.reg] = registers[i.val]
		}
	case "add":
		if val, err := strconv.Atoi(i.val); err == nil {
			registers[i.reg] += int(val)
		} else {
			registers[i.reg] += registers[i.val]
		}
	case "mul":
		if val, err := strconv.Atoi(i.val); err == nil {
			registers[i.reg] *= int(val)
		} else {
			registers[i.reg] *= registers[i.val]
		}
	case "mod":
		if val, err := strconv.Atoi(i.val); err == nil {
			registers[i.reg] %= int(val)
		} else {
			registers[i.reg] %= registers[i.val]
		}
	case "rcv":
		if registers[i.reg] > 0 {
			return
		}
	case "jgz":
		var y int
		if val, err := strconv.Atoi(i.val); err == nil {
			y = int(val)
		} else {
			y = registers[i.val]
		}
		if val, err := strconv.Atoi(i.reg); err == nil && val > 0 {
			*j += y - 1
		} else if registers[i.reg] > 0 {
			*j += y - 1
		}
	}
}

func run(inst []Instruction) int {
	registers := make(map[string]int)
	var j, freq int
	freq = 0
	for j = 0; j < int(len(inst)); j++ {
		i := inst[j]
		nextInst(i, &freq, registers, &j)
		if i.cmd == "rcv" && registers[i.reg] > 0 {
			break
		}
	}
	return freq
}

func runInst(pid int, i Instruction, reg map[string]int, j *int, qIn, qOut *[]int, lock *bool, count *int) {
	var val int
	nextInst(i, &val, reg, j)
	if i.cmd == "rcv" {
		if len(*qOut) > 0 {
			x := (*qOut)[0]
			*qOut = (*qOut)[1:]
			reg[i.reg] = x
			*lock = false
			*j++
		} else {
			*lock = true
		}
	} else {
		if i.cmd == "snd" {
			*count++
			*qIn = append(*qIn, val)
		}
		*lock = false
		*j++
	}
}

func run2(inst []Instruction) int {
	reg1 := make(map[string]int)
	reg2 := make(map[string]int)
	reg1["p"] = 0
	reg2["p"] = 1

	q1 := make([]int, 0)
	q2 := make([]int, 0)

	count1, count2 := 0, 0
	lock1, lock2 := false, false

	var j, k int
	for j, k = 0, 0; !(lock1 && lock2); {
		i := inst[j]
		runInst(0, i, reg1, &j, &q2, &q1, &lock1, &count1)

		i = inst[k]
		runInst(1, i, reg2, &k, &q1, &q2, &lock2, &count2)
	}
	return count2
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
		if len(strs) > 2 {
			i.val = strs[2]
		}
		inst = append(inst, i)
	}

	// Part 1
	fmt.Println(run(inst))

	// Part 2
	fmt.Println(run2(inst))
}
