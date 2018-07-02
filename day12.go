package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countCCNodes(g map[string][]string, from string) int {
	visited := make(map[string]bool, len(g))
	q := make([]string, 0)
	q = append(q, from)
	count := 0
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		if visited[x] {
			continue
		}
		visited[x] = true
		count++
		for _, v := range g[x] {
			if visited[v] {
				continue
			}
			q = append(q, v)
		}
	}
	return count
}

func countCC(g map[string][]string) int {
	visited := make(map[string]bool, len(g))
	q := make([]string, 0)
	count := 0

	for k, _ := range g {
		if visited[k] {
			continue
		}
		count++
		q = append(q, k)
		for len(q) > 0 {
			x := q[0]
			q = q[1:]
			if visited[x] {
				continue
			}
			visited[x] = true
			for _, v := range g[x] {
				if visited[v] {
					continue
				}
				q = append(q, v)
			}
		}
	}
	return count
}

func main() {
	graph := make(map[string][]string)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		t := scanner.Text()
		s := strings.Fields(t)
		u := s[0]
		for _, v := range s[2:] {
			v = strings.Trim(v, ",")
			graph[u] = append(graph[u], v)
			graph[v] = append(graph[v], u)
		}
	}

	// Part 1
	fmt.Println(countCCNodes(graph, "0"))

	// Part 2
	fmt.Println(countCC(graph))
}
