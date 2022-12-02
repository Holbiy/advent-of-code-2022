package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	score := 0

	valueOfShape := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	winsAgainst := map[string]string{
		"A": "B",
		"C": "A",
		"B": "C",
	}
	losesAgainst := map[string]string{
		"B": "A",
		"A": "C",
		"C": "B",
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		game := fileScanner.Text()
		player1 := game[0:1]
		player2 := ""
		gameResult := game[len(game)-1:]
		if gameResult == "Y" {
			player2 = player1
		}
		if gameResult == "X" {
			player2 = losesAgainst[player1]
		}
		if gameResult == "Z" {
			player2 = winsAgainst[player1]
		}
		score += valueOfShape[player2]
		score += getPoints(player1, player2)
	}

	fmt.Print(score)
	readFile.Close()
}

func getPoints(player1 string, player2 string) int {
	winsAgainst := map[string]string{
		"B": "A",
		"A": "C",
		"C": "B",
	}
	if player1 == player2 {
		return 3
	} else if winsAgainst[player2] == player1 {
		return 6
	} else {
		return 0
	}
}
