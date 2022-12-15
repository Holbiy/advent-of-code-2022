package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Position struct {
	X int
	Y int
}

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Print(err)
	}

	input := make([][]string, 0)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		row := strings.Split(fileScanner.Text(), "")
		input = append(input, row)
	}

	for y, row := range input {
		for x, cell := range row {

		}
	}

	print(input)
}
