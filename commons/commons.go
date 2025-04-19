package commons

import (
	"bufio"
	"fmt"
	"io"
	"log"
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
