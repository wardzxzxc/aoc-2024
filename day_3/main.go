package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/wardzxzxc/aoc-2024/commons"
)

func main() {
	part1()
}

func part1() {
	scanner := commons.GetInputFileScannerPtr()
	defer scanner.Close()

	re := regexp.MustCompile(`mul\(`)
	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		matches := re.FindAllStringIndex(line, -1)

		for _, match := range matches {
			firstIdx := match[1]

			var firstNumRunes []rune
			runes := []rune(line)
			validFirstNum := true

			// Max length including comma is 4
			for i := 0; i < 4; i++ {
				currentRune := runes[firstIdx+i]

				if i == 0 && !commons.IsNumeric(currentRune) {
					validFirstNum = false
					break
				}

				if i > 0 && !commons.IsNumeric(currentRune) && currentRune != ',' {
					validFirstNum = false
					break
				}

				if i > 0 && currentRune == ',' {
					break
				}

				firstNumRunes = append(firstNumRunes, currentRune)
			}

			if validFirstNum {
				firstNumStr := string(firstNumRunes)
				firstNum, _ := strconv.Atoi(firstNumStr)
				fmt.Println(firstNum)
				// secondIdx is firstIdx plus length of firstNum and comma
				secondIdx := firstIdx + len(firstNumRunes) + 1
				var secondNumRunes []rune
				validSecondNum := true

				for i := 0; i < 4; i++ {

					currentRune := runes[secondIdx+i]

					if i == 0 && !commons.IsNumeric(currentRune) {
						validSecondNum = false
						break
					}

					if i > 0 && !commons.IsNumeric(currentRune) && currentRune != ')' {
						validSecondNum = false
						break
					}

					if i > 0 && currentRune == ')' {
						break
					}

					secondNumRunes = append(secondNumRunes, currentRune)

				}

				if validSecondNum {
					secondNumStr := string(secondNumRunes)
					secondNum, _ := strconv.Atoi(secondNumStr)
					total += firstNum * secondNum
				}
			}
		}
	}
	fmt.Println(total)
}
