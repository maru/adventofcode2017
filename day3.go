package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

func abs(a int) int {
  if a < 0 {
    return -a
  }
  return a
}

func getMatrixLen(n int) int {
  c := 1
  for ; n > c*c; c += 2 {
  }
  return c
}

func getNumberOfSteps(n int) int {
  if n == 1 {
    return 0
  }

  // First count steps to the "circle"
  c := getMatrixLen(n)
  steps := c / 2
  // Get side
  baseNum := (c - 2)*(c - 2) + 1
  lenSide := c - 1

  // side = { 0: right, 1: up, 2: left, 3: down}
  side := 0
  for ; side < 4; side++ {
    if baseNum + side*lenSide <= n && n < baseNum + (side + 1)*lenSide {
      break
    }
  }
  // Number in the center of the side
  num := baseNum + side*lenSide + lenSide/2 - 1

  // Then walk to n
  dist := abs(num - n)
  steps += dist
  return steps
}

func getFirstLargerValue(n int) int {
  var m [][]int
  lenMatrix := getMatrixLen(n)
  m = make([][]int, lenMatrix)
  for i := 0; i < lenMatrix; i++ {
    m[i] = make([]int, lenMatrix)
  }

  // Center of matrix
  row, col := lenMatrix/2, lenMatrix/2
  m[row][col] = 1
  ret := -1
  offsetPos := [][]int {{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
  for lenSide := 2; ret < 0 && lenSide <= lenMatrix; lenSide += 2 {
    row, col = row + 1, col + 1
    for i := 0; ret < 0 && i < 4; i++ {
      for c := 0; ret < 0 && c < lenSide; c++ {
        row += offsetPos[i][0]
        col += offsetPos[i][1]
        if row < 0 || col < 0 || row >= lenMatrix || col >= lenMatrix {
          continue
        }
        offsetSum := [][]int {{-1, 0}, {0, -1}, {1, 0}, {0, 1}, {-1, 1}, {-1, -1}, {1, -1}, {1, 1}}
        for j := 0; j < 8; j++ {
          r := row + offsetSum[j][0]
          c := col + offsetSum[j][1]
          if r < 0 || c < 0 || r >= lenMatrix || c >= lenMatrix {
            continue
          }
          m[row][col] += m[r][c]
        }
        if m[row][col] > n {
          ret = m[row][col]
          break
        }
      }
    }
  }
  return ret
}

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  n, _ := strconv.Atoi(scanner.Text())

  // PART 1
  fmt.Println(getNumberOfSteps(n))

  // PART 2
  fmt.Println(getFirstLargerValue(n))

}
