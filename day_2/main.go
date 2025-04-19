package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/wardzxzxc/aoc-2024/commons"
)

func main() {
	part1()
	part2()
}

func part1() {
	scanner := commons.GetInputFileScannerPtr()
	defer scanner.Close()
	numOfReports := 0

	for scanner.Scan() {
		line := scanner.Text()
		numbersString := strings.Fields(line)
		report := make([]int, 0, len(numbersString))
		for _, numStr := range numbersString {
			num, _ := strconv.Atoi(numStr)
			report = append(report, num)
		}

		if isReportSafe(report) {
			numOfReports += 1
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
	scanner := commons.GetInputFileScannerPtr()
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
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	inputFilePath := path + "/input.txt"
	file, err := os.Open(inputFilePath)
	if err != nil {
		fmt.Println("Error reading input")
		fmt.Println(err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	return &FileScanner{file, scanner}
}
