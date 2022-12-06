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

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		text := fileScanner.Text()
		for i := 13; i < len(text); i++ {
			var chars []byte
			for y := 0; y < 14; y++ {
				chars = append(chars, text[i-y])
			}
			chars = removeDuplicateValues(chars)
			if len(chars) == 14 {
				println(i + 1)
				return
			}
		}
	}

	readFile.Close()
}

func removeDuplicateValues(intSlice []byte) []byte {
	keys := make(map[byte]bool)
	list := []byte{}

	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
