package main

import (
	"bufio"
	"fmt"
	"os"
	// "strings"
)

type Instruction struct {
  value int
  write int
  move int
  next byte
}

type State struct {
  state byte
  inst map[int]Instruction
}

func printTape(tape []int, n int, pos int) {
  fmt.Printf("pos %d | ", pos)
  for i := 0; i < n; i++ {
    if i == pos {
      fmt.Printf("[%d] ", tape[i])
    } else {
      fmt.Printf("%d ", tape[i])
    }
  }
  fmt.Println()
}

func simulate(states map[byte]State, curr byte, steps int) int {
  n := 1
  tape := make([]int, n)

  pos := 0
  for i := 0; i < steps; i++ {
    s := states[curr]
    // dynamically enlarge tape if necessary
    if pos < 0 {
      more := make([]int, n + n)
      copy(more[n:], tape)
      tape = more
      pos += n
      n += n
    } else if pos >= n {
      more := make([]int, n + n)
      copy(more, tape)
      tape = more
      n += n
    }
    inst := s.inst[tape[pos]]
    tape[pos] = inst.write
    pos += inst.move
    curr = inst.next
  }

  checksum := 0
  for _, v := range tape {
    checksum += v
  }
  return checksum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

  var begin byte
  scanner.Scan()
  fmt.Sscanf(scanner.Text(), "Begin in state %c.", &begin)

  var steps int
  scanner.Scan()
  fmt.Sscanf(scanner.Text(), "Perform a diagnostic checksum after %d steps.", &steps)

  states := make(map[byte]State)
  for scanner.Scan() {
    var s State
    scanner.Scan()
    fmt.Sscanf(scanner.Text(), "In state %c:", &s.state)

    s.inst = make(map[int]Instruction)
    for i := 0; i < 2; i ++ {
      var inst Instruction
      scanner.Scan()
      fmt.Sscanf(scanner.Text(), "  If the current value is %d:", &inst.value)

      scanner.Scan()
      fmt.Sscanf(scanner.Text(), "    - Write the value %d.", &inst.write)

      scanner.Scan()
      var dir string
      fmt.Sscanf(scanner.Text(), "    - Move one slot to the %s", &dir)
      if dir == "left." {
        inst.move = -1
      } else {
        inst.move = 1
      }

      scanner.Scan()
      fmt.Sscanf(scanner.Text(), "    - Continue with state %c.", &inst.next)

      s.inst[inst.value] = inst
    }
    states[s.state] = s
  }

	// Part 1
  fmt.Println(simulate(states, begin, steps))
}
