package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := readInput()
	result := getCountOfVisableTrees(input)
	fmt.Println(result)

	result2 := getHighestScenicScore(input)
	fmt.Println(result2)

}

func readInput() [][]int {
	var result [][]int

	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	index := 0
	for fileScanner.Scan() {
		result = append(result, make([]int, 0))
		row := fileScanner.Text()
		for _, c := range row {
			number, _ := strconv.Atoi(string(c))
			result[index] = append(result[index], number)
		}
		index++
	}

	readFile.Close()

	return result
}

func getHighestScenicScore(input [][]int) int {
	result := 0
	for ri, r := range input {
		for ci, _ := range r {
			scenicScore := getScenicScore(input, ri, ci)
			if result == 0 || scenicScore > result {
				result = scenicScore
			}
		}
	}
	return result
}

func getScenicScore(input [][]int, r int, c int) int {

	colLength := len(input)
	rowLenth := len(input[0])
	heigth := input[r][c]

	scoreUp := 0
	scoreDown := 0
	scoreLeft := 0
	scoreRight := 0

	for i := r + 1; i < rowLenth; i++ {
		if r+1 == rowLenth {
			break
		}
		if input[i][c] < heigth {
			scoreRight++
		} else {
			scoreRight++
			break
		}
	}

	for i := c + 1; i < colLength; i++ {
		if c+1 == colLength {
			break
		}
		if input[r][i] < heigth {
			scoreDown++
		} else {
			scoreDown++
			break
		}
	}

	for i := r - 1; i >= 0; i-- {
		if r == 0 {
			break
		}
		if input[i][c] < heigth {
			scoreLeft++
		} else {
			scoreLeft++
			break
		}
	}

	for i := c - 1; i >= 0; i-- {
		if c+1 == colLength {
			break
		}
		if input[r][i] < heigth {
			scoreUp++
		} else {
			scoreUp++
			break
		}
	}

	return scoreDown * scoreUp * scoreLeft * scoreRight
}

func getCountOfVisableTrees(input [][]int) int {
	result := 0
	for ri, r := range input {
		for ci, _ := range r {
			if isTreeVisable(input, ri, ci) {
				result++
			}
		}
	}
	return result
}

func isTreeVisable(input [][]int, r int, c int) bool {

	colLength := len(input)
	rowLenth := len(input[0])
	heigth := input[r][c]

	if r == 0 || c == 0 || r == rowLenth-1 || c == colLength-1 {
		return true
	}

	for i := r + 1; i < rowLenth; i++ {
		if input[i][c] >= heigth {
			break
		}
		if i == rowLenth-1 {
			return true
		}
	}

	for i := c + 1; i < colLength; i++ {
		if input[r][i] >= heigth {
			break
		}
		if i == colLength-1 {
			return true
		}
	}

	for i := r - 1; i >= 0; i-- {
		if input[i][c] >= heigth {
			break
		}
		if i == 0 {
			return true
		}
	}

	for i := c - 1; i >= 0; i-- {
		if input[r][i] >= heigth {
			break
		}
		if i == 0 {
			return true
		}
	}

	return false
}
