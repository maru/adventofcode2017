package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func isValidNotAnagram(v []string) bool {
	m := make(map[string]bool)
	for _, w := range v {
		w1 := sortString(w)
		if m[w1] {
			return false
		}
		m[w1] = true
	}
	return true
}

func isValidDifferent(v []string) bool {
	m := make(map[string]bool)
	for _, w := range v {
		if m[w] {
			return false
		}
		m[w] = true
	}
	return true
}

func countValidPassphrases(t []string, isValidPassphrase func([]string) bool) int {
	count := 0
	for _, s := range t {
		v := strings.Fields(s)
		if isValidPassphrase(v) {
			count++
		}
	}
	return count
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var t []string
	for scanner.Scan() {
		t = append(t, scanner.Text())
	}

	// PART 1
	fmt.Println(countValidPassphrases(t, isValidDifferent))

	// PART 2
	fmt.Println(countValidPassphrases(t, isValidNotAnagram))
}
