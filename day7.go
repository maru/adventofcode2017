package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TreeNode struct {
	name           string
	weight         int
	childrenWeight int
	level          int
	children       []string
}

func getRoot(tree map[string]TreeNode) string {
	isChild := make(map[string]bool)
	for _, node := range tree {
		for _, child := range node.children {
			isChild[child] = true
		}
	}

	for _, node := range tree {
		if ok, _ := isChild[node.name]; !ok {
			return node.name
		}
	}
	return ""
}

func getWeight(tree map[string]TreeNode, s string) int {
  node := tree[s]
  if node.childrenWeight == 0 {
    for _, c := range node.children {
      node.childrenWeight += getWeight(tree, c)
    }
  }
  w := node.childrenWeight + node.weight
  return w
}

func getUnbalanced(tree map[string]TreeNode, s string) int {
  node := tree[s]
  if len(node.children) == 0 {
    return 0
  }
  for _, c := range node.children {
    if v := getUnbalanced(tree, c); v != 0 {
      return v
    }
  }

  weights := make(map[int][]string)
  for _, c := range node.children {
    w := getWeight(tree, c)
    weights[w] = append(weights[w], c)
  }
  if len(weights) == 1 {
    return 0
  }
  var w1, w2 int
  for w, count := range weights {
    if len(count) == 1 {
      w2 = w
    } else {
      w1 = w
    }
  }
  unb := weights[w2][0]
  w := tree[unb].weight
  if w1 < w2 {
    return w - (w2 - w1)
  }
  return w - (w1 - w2)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

  tree := make(map[string]TreeNode)
	for scanner.Scan() {
		s := scanner.Text()
		var node TreeNode
		strs := strings.Fields(s)
		node.name = strs[0]
		fmt.Sscanf(strs[1], "(%d)", &node.weight)
		for i := 3; i < len(strs); i++ {
			node.children = append(node.children, strings.TrimRight(strs[i], ","))
		}
    tree[node.name] = node
	}

	// PART 1
	root := getRoot(tree)
	fmt.Println(root)

	// PART 1
	fmt.Println(getUnbalanced(tree, root))
}
