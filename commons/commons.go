package commons

import (
	"bufio"
	"io"
	"os"
)

type FileScanner struct {
	io.Closer
	*bufio.Scanner
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func GetInputFileScannerPtr() *FileScanner {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	inputFilePath := path + "/input.txt"
	file, err := os.Open(inputFilePath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	return &FileScanner{file, scanner}
}

func IsNumeric(r rune) bool {
	return r >= '0' && r <= '9'
}

func FindIntersectionStr(arr1 []string, arr2 []string) []string {
	var result []string

	set := make(map[string]struct{})

	for _, v := range arr2 {
		set[v] = struct{}{}
	}

	for _, v := range arr1 {
		if _, ok := set[v]; ok {
			result = append(result, v)
		}
	}

	return result
}

func FindIntersectionInt(arr1 []int, arr2 []int) []int {
	var result []int

	set := make(map[int]struct{})

	for _, v := range arr2 {
		set[v] = struct{}{}
	}

	for _, v := range arr1 {
		if _, ok := set[v]; ok {
			result = append(result, v)
		}
	}

	return result
}
