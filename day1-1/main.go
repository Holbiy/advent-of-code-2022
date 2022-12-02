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
	//part1()
	part2()
}

func part1() {
	max := 0
	currentSum := 0

	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		currentNumber, err := strconv.Atoi(strings.TrimSpace(fileScanner.Text()))
		if err != nil {
			fmt.Println(err.Error())
			currentSum = 0
		}
		currentSum += currentNumber
		if currentSum > max {
			max = currentSum
		}
	}

	fmt.Print(max)
	readFile.Close()
}

func part2() {
	var sums []int
	currentSum := 0

	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		currentNumber, err := strconv.Atoi(strings.TrimSpace(fileScanner.Text()))
		if err != nil {
			sums = append(sums, currentSum)
			currentSum = 0
		}
		currentSum += currentNumber
	}

	sort.Ints(sums)
	result := sums[len(sums)-1] + sums[len(sums)-2] + sums[len(sums)-3]

	fmt.Print(result)
	readFile.Close()
}
