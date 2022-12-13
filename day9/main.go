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

type Position struct {
	X int
	Y int
}

func part1() {
	positions := make(map[string]bool, 0)

	head := Position{
		X: 0,
		Y: 0,
	}

	var tails []Position

	tail := Position{
		X: 0,
		Y: 0,
	}

	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		text := fileScanner.Text()
		splittedText := strings.Split(text, " ")
		direction := splittedText[0]
		count, _ := strconv.Atoi(splittedText[1])
		for i := 0; i < count; i++ {
			switch direction {
			case "U":
				moveUp(&head, &tail)
			case "D":
				moveDown(&head, &tail)
			case "R":
				moveRight(&head, &tail)
			case "L":
				moveLeft(&head, &tail)
			}
			tailString := strconv.Itoa(tail.X) + "/" + strconv.Itoa(tail.Y)
			positions[tailString] = true
		}
	}

	fmt.Println(len(positions))
}

func moveUp(head *Position, tail *Position) {
	if head.X == tail.X && head.Y == tail.Y {
		head.Y++
		return
	}
	if head.X == tail.X && head.Y+1 == tail.Y {
		head.Y++
		return
	}
	if head.X == tail.X {
		head.Y++
		tail.Y++
		return
	}
	if head.Y == tail.Y {
		head.Y++
		return
	}
	if head.Y-1 == tail.Y && (head.X+1 == tail.X || head.X-1 == tail.X) {
		tail.X = head.X
		tail.Y = head.Y
	}
	head.Y++
}

func moveDown(head *Position, tail *Position) {
	if head.X == tail.X && head.Y == tail.Y {
		head.Y--
		return
	}
	if head.X == tail.X && head.Y-1 == tail.Y {
		head.Y--
		return
	}
	if head.X == tail.X {
		head.Y--
		tail.Y--
		return
	}
	if head.Y == tail.Y {
		head.Y--
		return
	}
	if head.Y+1 == tail.Y && (head.X+1 == tail.X || head.X-1 == tail.X) {
		tail.X = head.X
		tail.Y = head.Y
	}
	head.Y--
}

func moveRight(head *Position, tail *Position) {
	if head.X == tail.X && head.Y == tail.Y {
		head.X++
		return
	}
	if head.Y == tail.Y && head.X+1 == tail.X {
		head.X++
		return
	}
	if head.Y == tail.Y {
		head.X++
		tail.X++
		return
	}
	if head.X == tail.X {
		head.X++
		return
	}
	if head.X-1 == tail.X && (head.Y+1 == tail.Y || head.Y-1 == tail.Y) {
		tail.Y = head.Y
		tail.X = head.X
	}
	head.X++
}

func moveLeft(head *Position, tail *Position) {
	if head.X == tail.X && head.Y == tail.Y {
		head.X--
		return
	}
	if head.Y == tail.Y && head.X-1 == tail.X {
		head.X--
		return
	}
	if head.Y == tail.Y {
		head.X--
		tail.X--
		return
	}
	if head.X == tail.X {
		head.X--
		return
	}
	if head.X+1 == tail.X && (head.Y+1 == tail.Y || head.Y-1 == tail.Y) {
		tail.Y = head.Y
		tail.X = head.X
	}
	head.X--
}
