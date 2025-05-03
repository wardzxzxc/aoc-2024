package main

import (
	"fmt"
	"strings"

	"github.com/wardzxzxc/aoc-2024/commons"
)

var offsets = []struct {
	x int
	y int
}{
	{1, 0},
	{1, 1},
	{0, 1},
	{-1, 1},
	{-1, 0},
	{-1, -1},
	{0, -1},
	{1, -1},
}

var XMAS = [4]string{"X", "M", "A", "S"}

func main() {
	part1()
	part2()
}

func checkStartingFrom(matrix [][]string, x int, y int) int {
	count := 0

	for _, dir := range offsets {
		char := 1
		for i := 1; i < 4; i++ {
			x := x + dir.x*i
			y := y + dir.y*i
			if x < 0 || y < 0 || x >= len(matrix[0]) || y >= len(matrix) {
				break
			}
			if matrix[x][y] == XMAS[i] {
				char += 1
			}
		}
		if char == 4 {
			count += 1
		}
	}
	return count
}

func checkAroundA(matrix [][]string, x int, y int) bool {
	if x == 0 || y == 0 || x == len(matrix[0])-1 || y == len(matrix[0])-1 {
		return false
	}

	topLeft := matrix[x-1][y-1]
	bottomRight := matrix[x+1][y+1]
	topRight := matrix[x+1][y-1]
	bottomLeft := matrix[x-1][y+1]

	if (topLeft == "M" && bottomRight == "S") || (topLeft == "S" && bottomRight == "M") {
		if (topRight == "M" && bottomLeft == "S") || (topRight == "S" && bottomLeft == "M") {
			return true
		}
	}
	return false
}

func part1() {
	scanner := commons.GetInputFileScannerPtr()
	defer scanner.Close()
	matrix := make([][]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, "")
		matrix = append(matrix, chars)
	}

	validCount := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == "X" {
				validCount += checkStartingFrom(matrix, i, j)
			}
		}
	}
	fmt.Println(validCount)
}

func part2() {
	scanner := commons.GetInputFileScannerPtr()
	defer scanner.Close()
	matrix := make([][]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, "")
		matrix = append(matrix, chars)
	}

	validCount := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == "A" && checkAroundA(matrix, i, j) {
				validCount += 1
			}
		}
	}
	fmt.Println(validCount)
}
