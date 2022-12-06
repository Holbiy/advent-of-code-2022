package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
}

func part1() {
	collumns := ReadCollumns()
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		if fileScanner.Text() != "" && fileScanner.Text()[:4] == "move" {
			sliced := strings.Split(fileScanner.Text(), " ")
			count, _ := strconv.Atoi(sliced[1])
			from, _ := strconv.Atoi(sliced[3])
			to, _ := strconv.Atoi(sliced[5])
			from--
			to--

			copy := collumns[from][len(collumns[from])-count:]

			collumns[from] = collumns[from][:len(collumns[from])-count]

			collumns[to] = append(collumns[to], copy...)

		}
	}

	for i := 0; i < len(collumns)-1; i++ {
		print(collumns[i][len(collumns[i])-1])
	}
}

func ReadCollumns() [][]string {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var input []string

	for fileScanner.Scan() {
		if fileScanner.Text() == "" {
			break
		}
		input = append(input, fileScanner.Text())
	}

	var cratesLists [9][]string

	lines := strings.Split(input[len(input)-1], " ")
	input = input[:len(input)-1]

	for _, v := range lines {
		if v == "" {
			continue
		}

		number, _ := strconv.Atoi(v)
		index := (number-1)*4 + 1
		for _, i := range input {
			if string(i[index]) != " " {
				cratesLists[number-1] = append(cratesLists[number-1], string(i[index]))
			}
		}
	}
	for _, list := range cratesLists {
		for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
			list[i], list[j] = list[j], list[i]
		}
	}

	return cratesLists[:]
}
