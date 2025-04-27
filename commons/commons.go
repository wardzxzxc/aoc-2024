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
