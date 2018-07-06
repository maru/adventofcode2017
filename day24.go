package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
  "strconv"
)

type Port struct {
  a, b int
}

func getMaxW(w int, maxW *int, l int, maxL *int) {
  if w > *maxW {
    *maxW = w
  }
}

func getMaxL(w int, maxW *int, l int, maxL *int) {
  if l > *maxL || (l == *maxL && w > *maxW) {
    *maxL = l
    *maxW = w
  }
}

func nextPort(ports []Port, n int, visited *[]bool, value int, sum int, len int,
              getMax func(int,*int,int,*int)) (int, int) {
  maxW := sum
  maxL := len
  for i := 0; i < n; i++ {
    p := ports[i]
    if (*visited)[i] { continue }
    w, l := 0, 0
    if p.a == value || p.b == value {
      (*visited)[i] = true
      next := p.a
      if p.a == value { next = p.b }
      w, l = nextPort(ports, n, visited, next, sum + p.a + p.b, len + 1, getMax)
      (*visited)[i] = false
    }
    getMax(w, &maxW, l, &maxL)
  }
  return maxW, maxL
}

func getMaxWeight(ports []Port) int {
  n := len(ports)
  visited := make([]bool, n)
  w, _ := nextPort(ports, n, &visited, 0, 0, 0, getMaxW)
  return w
}

func getMaxLength(ports []Port) int {
  n := len(ports)
  visited := make([]bool, n)
  w, _ := nextPort(ports, n, &visited, 0, 0, 0, getMaxL)
  return w
}

func main() {
  var ports []Port
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
    s := scanner.Text()
  	strs := strings.Split(s, "/")
    var p Port
    p.a, _ = strconv.Atoi(strs[0])
    p.b, _ = strconv.Atoi(strs[1])
    if p.a > p.b {
      p.a, p.b = p.b, p.a
    }
    ports = append(ports, p)
  }

	// Part 1
  fmt.Println(getMaxWeight(ports))

	// Part 2
  fmt.Println(getMaxLength(ports))
}
