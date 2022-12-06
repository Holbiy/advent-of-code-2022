package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part2()
}

func part1() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	counter := 0

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		pair := strings.Split(fileScanner.Text(), ",")
		sections1 := strings.Split(pair[0], "-")
		sections2 := strings.Split(pair[1], "-")
		section1_1, _ := strconv.Atoi(sections1[0])
		section1_2, _ := strconv.Atoi(sections1[1])
		section2_1, _ := strconv.Atoi(sections2[0])
		section2_2, _ := strconv.Atoi(sections2[1])
		if section1_1 >= section2_1 && section1_2 <= section2_2 {
			counter++
			continue
		}
		if section1_1 <= section2_1 && section1_2 >= section2_2 {
			counter++
			continue
		}
	}

	print(counter)
	readFile.Close()
}

func part2() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	counter := 0

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		pair := strings.Split(fileScanner.Text(), ",")
		sections1 := strings.Split(pair[0], "-")
		sections2 := strings.Split(pair[1], "-")
		section1_1, _ := strconv.Atoi(sections1[0])
		section1_2, _ := strconv.Atoi(sections1[1])
		section2_1, _ := strconv.Atoi(sections2[0])
		section2_2, _ := strconv.Atoi(sections2[1])

		if section1_1 >= section2_1 && section1_2 <= section2_2 {
			counter++
			continue
		}
		if section1_1 <= section2_1 && section1_2 >= section2_2 {
			counter++
			continue
		}
		if section1_1 >= section2_1 && section1_1 <= section2_2 {
			counter++
			continue
		}
		if section1_2 >= section2_1 && section1_2 <= section2_2 {
			counter++
			continue
		}
		if section1_1 <= section2_1 && section1_1 >= section2_2 {
			counter++
			continue
		}
		if section1_2 <= section2_1 && section1_2 >= section2_2 {
			counter++
			continue
		}
	}

	print(counter)
	readFile.Close()
}
