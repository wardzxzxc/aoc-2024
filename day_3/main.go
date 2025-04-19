package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/wardzxzxc/aoc-2024/commons"
)

func main() {
	part1()
}

type DigitNumbers struct {
	X int
	Y int
}

func getDigitNumbers(startIdx int, runes []rune) (*DigitNumbers, error) {
	var firstNumRunes []rune
	validFirstNum := true
	for i := 0; i < 4; i++ {
		currentRune := runes[startIdx+i]

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

	if !validFirstNum {
		return nil, errors.New("Invalid First Digit Number")
	} else {
		secondIdx := startIdx + len(firstNumRunes) + 1
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
			firstNumStr := string(firstNumRunes)
			firstNum, _ := strconv.Atoi(firstNumStr)
			return &DigitNumbers{firstNum, secondNum}, nil
		} else {
			return nil, errors.New("Invalid Second Digit Number")
		}
	}
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
			runes := []rune(line)

			digitNumbersPtr, err := getDigitNumbers(firstIdx, runes)
			if err != nil {
				continue
			} else {
				total += digitNumbersPtr.X * digitNumbersPtr.Y
			}

		}
	}

	fmt.Println(total)
}
