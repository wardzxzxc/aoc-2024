package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/wardzxzxc/aoc-2024/commons"
)

func main() {
	part1()
	part2()
}

func part1() {
	scanner := commons.GetInputFileScannerPtr()
	defer scanner.Close()

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			x, xErr := strconv.Atoi(match[1])
			y, yErr := strconv.Atoi(match[2])

			if xErr == nil && yErr == nil {
				total += x * y
			}
		}
	}

	fmt.Println(total)
}

func part2() {
	scanner := commons.GetInputFileScannerPtr()
	defer scanner.Close()

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	total := 0
	enable := true

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			fullMatch := match[0]

			if fullMatch == "do()" {
				enable = true
				continue
			} else if fullMatch == "don't()" {
				enable = false
				continue
			} else {
				if enable {
					xStr := match[1]
					yStr := match[2]
					x, xErr := strconv.Atoi(xStr)
					y, yErr := strconv.Atoi(yStr)
					if xErr == nil && yErr == nil {
						total += x * y
					}
				}
			}
		}
	}
	fmt.Println(total)
}
