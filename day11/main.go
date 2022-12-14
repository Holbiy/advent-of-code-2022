package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	Items           []uint64
	Operation       string
	OperationNumber string
	Test            uint64
	IfTrue          uint64
	IfFalse         uint64
	inspected       uint64
}

func main() {
	divisorKgv := 1

	monkeys := make(map[uint64]*Monkey, 1)
	monkeys[0] = &Monkey{}

	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Print(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var id uint64 = 0
	for fileScanner.Scan() {
		text := fileScanner.Text()
		if text == "" {
			id++
			monkeys[id] = &Monkey{}
		}

		monkey := monkeys[id]
		if strings.Contains(text, "Starting") {
			index := strings.Index(text, ":")
			numbers := strings.Split(strings.Replace(text[index+1:], " ", "", -1), ",")
			for _, number := range numbers {
				item, _ := strconv.ParseUint(number, 10, 64)
				monkey.Items = append(monkey.Items, item)
			}
		}
		if strings.Contains(text, "Operation") {
			monkey.Operation = strings.Split(text, " ")[6]
			monkey.OperationNumber = strings.Split(text, " ")[7]
		}
		if strings.Contains(text, "Test") {
			monkey.Test, _ = strconv.ParseUint(strings.Split(text, " ")[5], 10, 64)
			divisorKgv *= int(monkey.Test)
		}
		if strings.Contains(text, "true") {
			monkey.IfTrue, _ = strconv.ParseUint(strings.Split(text, " ")[9], 10, 64)
		}
		if strings.Contains(text, "false") {
			monkey.IfFalse, _ = strconv.ParseUint(strings.Split(text, " ")[9], 10, 64)
		}
	}

	for i := 0; i < 10000; i++ {
		for y := 0; y < len(monkeys); y++ {
			monkey := monkeys[uint64(y)]
			for _, item := range monkey.Items {
				monkey.inspected++
				operationNumber, err := strconv.ParseUint(monkey.OperationNumber, 10, 64)
				if err != nil {
					operationNumber = item
				}

				var worry uint64
				switch monkey.Operation {
				case "*":
					worry = item * operationNumber
				case "+":
					worry = item + operationNumber
				}

				worry = worry % uint64(divisorKgv)

				if worry%monkey.Test != 0 {
					monkeys[monkey.IfFalse].Items = append(monkeys[monkey.IfFalse].Items, worry)
				} else {
					monkeys[monkey.IfTrue].Items = append(monkeys[monkey.IfTrue].Items, worry)
				}
				monkey.Items = []uint64{}
			}
		}
	}

	for i := 0; i < len(monkeys); i++ {
		fmt.Println(monkeys[uint64(i)].inspected)
	}
	fmt.Print(monkeys)
	readFile.Close()
}
