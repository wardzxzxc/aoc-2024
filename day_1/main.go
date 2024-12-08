package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	firstSlicePtr, secondSlicePtr := getSlices()
	firstSlice := *firstSlicePtr
	secondSlice := *secondSlicePtr
	sort.Slice(firstSlice, func(i, j int) bool {
		return firstSlice[i] < firstSlice[j]
	})
	sort.Slice(secondSlice, func(i, j int) bool {
		return secondSlice[i] < secondSlice[j]
	})

	sum := 0
	for i := 0; i < len(firstSlice); i++ {
		diff := firstSlice[i] - secondSlice[i]
		if diff < 0 {
			diff = diff * -1
		}
		sum += diff
	}
	fmt.Println(sum)
}

func part2() {
	firstSlicePtr, secondSlicePtr := getSlices()

	buckets := make(map[int]int)

	for _, val := range *secondSlicePtr {
		freq, ok := buckets[val]
		if ok {
			buckets[val] = freq + 1
		} else {
			buckets[val] = 1
		}
	}

	sum := 0
	for _, val := range *firstSlicePtr {
		freq, ok := buckets[val]
		if ok {
			sum += val * freq
		} else {
			sum += 0
		}
	}

	fmt.Println(sum)
}

func getSlices() (*[]int, *[]int) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading input")
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var firstSlice []int
	var secondSlice []int

	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		firstNum, _ := strconv.Atoi(numbers[0])
		secondNum, _ := strconv.Atoi(numbers[1])
		firstSlice = append(firstSlice, firstNum)
		secondSlice = append(secondSlice, secondNum)
	}
	return &firstSlice, &secondSlice
}
