package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/wardzxzxc/aoc-2024/commons"
)

func main() {
	part1()
	part2()
}

func getOrderRulesArrAndPagesArr(scanner *commons.FileScanner) (*map[string][]string, *[][]string) {
	orderRules := make(map[string][]string)
	var pagesArr [][]string
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "|") {
			orderSeq := strings.Split(line, "|")
			firstPage, secondPage := orderSeq[0], orderSeq[1]
			orderRules[firstPage] = append(orderRules[firstPage], secondPage)
		} else if strings.Contains(line, ",") {
			pages := strings.Split(line, ",")
			pagesArr = append(pagesArr, pages)
		}

	}

	return &orderRules, &pagesArr
}

func checkIsValid(pagesPtr *[]string, orderRulesPtr *map[string][]string) bool {
	pages := *pagesPtr
	orderRules := *orderRulesPtr
	isValid := true
	length := len(pages)

	for i, pageNum := range pages {
		orders := orderRules[pageNum]
		// If no order for that page, consider sequence not valid
		if len(orders) == 0 {
			isValid = false
			break
		}

		// Go through each remaining pageNum and check
		for j := i + 1; j < length-1; j++ {
			if !slices.Contains(orders, pages[j]) {
				isValid = false
			}
		}
		if !isValid {
			return isValid
		}
	}
	return isValid
}

func part1() {
	scanner := commons.GetInputFileScannerPtr()
	defer scanner.Close()
	orderRulesPtr, pagesArrPtr := getOrderRulesArrAndPagesArr(scanner)
	pagesArr := *pagesArrPtr
	isValid := true
	totalSum := 0

	for _, pages := range pagesArr {
		length := len(pages)
		isValid = checkIsValid(&pages, orderRulesPtr)
		if isValid {
			var midIdx int
			if length%2 == 1 {
				midIdx = (length - 1) / 2
			} else {
				midIdx = length / 2
			}
			midValString := pages[midIdx]
			midVal, _ := strconv.Atoi(midValString)
			totalSum += midVal

		}

		isValid = true
	}

	fmt.Println(totalSum)
}

func getMiddleNumberFromUnordered(orderRulesPtr *map[string][]string, pagesPtr *[]string) int {
	pages := *pagesPtr
	gapToMiddle := len(pages) / 2
	orderRules := *orderRulesPtr
	pageNumInt := -1
	for _, pageNum := range pages {
		orders := orderRules[pageNum]
		intersect := commons.FindIntersectionStr(orders, pages)
		if len(intersect) == gapToMiddle {
			pageNumInt, _ = strconv.Atoi(pageNum)
			return pageNumInt
		}
	}
	return pageNumInt
}

func part2() {
	scanner := commons.GetInputFileScannerPtr()
	defer scanner.Close()
	orderRulesPtr, pagesArrPtr := getOrderRulesArrAndPagesArr(scanner)
	pagesArr := *pagesArrPtr
	isValid := true
	totalSum := 0

	for _, pages := range pagesArr {
		isValid = checkIsValid(&pages, orderRulesPtr)
		if !isValid {
			midNum := getMiddleNumberFromUnordered(orderRulesPtr, &pages)
			totalSum += midNum
		}

	}

	fmt.Println(totalSum)
}
