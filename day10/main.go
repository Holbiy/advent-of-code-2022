package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(param1, param2) (int, error) {
	part1()
	part2()
}

func part1() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	test := make([]int, 0)

	test := 0
	cycles := 0
	x := 1

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		text := fileScanner.Text()
		if strings.Contains(text, "addx") {
			count, _ := strconv.Atoi(strings.Split(text, " ")[1])
			cycles++
			if cycles == 20 || cycles == 60 || cycles == 100 || cycles == 140 || cycles == 180 || cycles == 220 {
				result += x * cycles
			}
			cycles++
			if cycles == 20 || cycles == 60 || cycles == 100 || cycles == 140 || cycles == 180 || cycles == 220 {
				result += x * cycles
			}
			x += count
		} else {
			cycles++
			if cycles == 20 || cycles == 60 || cycles == 100 || cycles == 140 || cycles == 180 || cycles == 220 {
				result += x * cycles
			}

		}
	}
	//Comment
	/*
	   Long comment
	*/
	fmt.Print(result)
	readFile.Close()
}

func part2() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	result := make([]string, 1)
	cycles := 0
	x := 1

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		text := fileScanner.Text()
		if strings.Contains(text, "addx") {
			count, _ := strconv.Atoi(strings.Split(text, " ")[1])
			proccedCycle(&cycles, &result, &x)
			proccedCycle(&cycles, &result, &x)
			x += count
		} else {
			proccedCycle(&cycles, &result, &x)
		}
	}

	for _, s := range result {
		fmt.Println(s)
	}

	readFile.Close()
}

func proccedCycle(cylcle *int, res *[]string, x *int) {
	*cylcle++
	var newChar string
	lenOfNewestString := len((*res)[len((*res))-1])
	if lenOfNewestString > 39 {
		(*res) = append((*res), "")
		lenOfNewestString = 0
	}

	if lenOfNewestString+1 >= *x && lenOfNewestString+1 < *x+3 {
		newChar = "#"
	} else {
		newChar = "."
	}

	(*res)[len((*res))-1] += newChar

}
