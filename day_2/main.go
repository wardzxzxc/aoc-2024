package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/wardzxzxc/aoc-2024/commons"
)

type FileScanner struct {
	io.Closer
	*bufio.Scanner
}

func main() {
	part1()
	part2()
}

func part1() {
	scanner := getScannerPtr()
	defer scanner.Close()
	numOfReports := 0

	for scanner.Scan() {
		line := scanner.Text()
		numbersString := strings.Fields(line)
		var direction int // 0 for negative, 1 for positive
		for idx, numStr := range numbersString {
			if idx != 0 {
				num, _ := strconv.Atoi(numStr)
				prevNumStr := numbersString[idx-1]
				prevNum, _ := strconv.Atoi(prevNumStr)

				diff := prevNum - num

				if idx == 1 {
					if diff < 0 {
						direction = 0
					} else {
						direction = 1
					}
				} else {
					if diff < 0 && direction == 1 {
						break
					}
					if diff > 0 && direction == 0 {
						break
					}
				}

				if diff < 0 {
					diff = diff * -1
				}

				if diff > 3 || diff < 1 {
					break
				}

				if idx == len(numbersString)-1 {
					numOfReports += 1
				}
			}
		}
	}
	fmt.Println(numOfReports)
}

func isReportSafe(report []int) bool {
	isIncreasing := report[1] > report[0]

	for i := 0; i < len(report)-1; i++ {
		notOrdered := isIncreasing && report[i] > report[i+1] || !isIncreasing && report[i] < report[i+1]
		invalidDiff := report[i] == report[i+1] || commons.Abs(report[i]-report[i+1]) > 3

		if notOrdered || invalidDiff {
			return false
		}
	}

	return true
}

func part2() {
	scanner := getScannerPtr()
	defer scanner.Close()

	numReports := 0
	extraReports := 0

	for scanner.Scan() {
		line := scanner.Text()
		numbersString := strings.Fields(line)
		report := make([]int, 0, len(numbersString))
		for _, numStr := range numbersString {
			num, _ := strconv.Atoi(numStr)
			report = append(report, num)
		}

		if isReportSafe(report) {
			numReports += 1
		} else {
			for i := 0; i < len(report); i++ {
				copyReport := make([]int, len(report))
				copy(copyReport, report)

				if i == len(copyReport)-1 {
					copyReport = copyReport[:i]
				} else {
					copyReport = append(copyReport[:i], copyReport[i+1:]...)
				}
				if isReportSafe(copyReport) == true {
					extraReports += 1
					break
				}
			}
		}
	}
	fmt.Println(extraReports + numReports)
}

func getScannerPtr() *FileScanner {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading input")
		fmt.Println(err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	return &FileScanner{file, scanner}
}
