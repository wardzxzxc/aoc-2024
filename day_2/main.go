package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type FileScanner struct {
	io.Closer
	*bufio.Scanner
}

func main() {
	part1()
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
