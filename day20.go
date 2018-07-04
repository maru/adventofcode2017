package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type d3 struct {
  x, y, z int
}

type Particle struct {
  p, v, a d3
}

func abs(a int) int {
  if a >= 0 {
    return a
  }
  return -a
}

func (P *Particle) read(s string, p *d3) {
  var c byte
  fmt.Sscanf(s, "%c=<%d,%d,%d>", &c, &p.x, &p.y, &p.z)
}

func (P *Particle) move() {
  P.v.x += P.a.x
  P.v.y += P.a.y
  P.v.z += P.a.z

  P.p.x += P.v.x
  P.p.y += P.v.y
  P.p.z += P.v.z
}

func (P *Particle) dist() int {
  return abs(P.p.x) + abs(P.p.y) + abs(P.p.z)
}

func simulate(part []Particle) (int, int) {
  n := len(part)
  dist := make([]int, n)
  pos := make(map[string]int)
  run := make([]bool, n)
  for i := 0; i < n; i++ {
    run[i] = true
  }

  closest := 0
  left := n
  rounds := 1 << 12
  for i := 0; i < rounds; i++ {
    for j := 0; j < n; j++ {
      if run[j] {
        p := part[j]
        p.move()
        dist[j] = p.dist()
        part[j] = p
        s := fmt.Sprintf("%d,%d,%d", p.p.x, p.p.y, p.p.z)
        if v, ok := pos[s]; ok {
          if run[v] {
            run[v] = false
            left--
          }
          run[j] = false
          left--
        } else {
          pos[s] = j
        }
      }
    }

    /* DEBUG
    cl := 0
    for k, v := range dist {
    	if v < dist[cl] {
    		cl = k
    	}
    }
    if closest != cl {
      closest = cl
      fmt.Printf("closest=%d -> %d\n", closest, dist[closest])
    }
    */
  }

  for k, v := range dist {
  	if v < dist[closest] {
  		closest = k
  	}
  }
	return closest, left
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
  var part []Particle
	for scanner.Scan() {
    s := scanner.Text()
    strs := strings.Split(s, ", ")
    var p Particle
    p.read(strs[0], &p.p)
    p.read(strs[1], &p.v)
    p.read(strs[2], &p.a)
    part = append(part, p)
  }

  closest, left := simulate(part)

	// Part 1
  fmt.Println(closest)

	// Part 2
  fmt.Println(left)
}
