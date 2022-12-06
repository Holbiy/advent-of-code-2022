package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	var result []rune

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		halfLength := len(fileScanner.Text()) / 2

		fistHalf := fileScanner.Text()[:halfLength]
		secondHalf := fileScanner.Text()[halfLength:]

		for _, c := range fistHalf {
			if strings.ContainsRune(secondHalf, c) {
				result = append(result, c)
				break
			}
		}
	}

	score := calculateScore(result)

	fmt.Println(score)
	readFile.Close()
}

func part2() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	var rucksacks []string
	var result []rune

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		rucksacks = append(rucksacks, fileScanner.Text())
	}

	for i, _ := range rucksacks {
		if i%3 == 0 {
			for _, c := range rucksacks[i] {
				if strings.ContainsRune(rucksacks[i+1], c) && strings.ContainsRune(rucksacks[i+2], c) {
					result = append(result, c)
					break
				}
			}
		}
	}

	score := calculateScore(result)

	fmt.Println(score)
	readFile.Close()
}

func calculateScore(arr []rune) int {
	result := 0
	for _, c := range arr {
		if int(c)-96 > 0 {
			result += int(c) - 96
		} else {
			result += int(c) - 38
		}
	}

	return result
}
