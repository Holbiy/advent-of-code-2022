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

type File struct {
	Name     string  `json:"name"`
	Size     int     `json:"size"`
	Children []*File `json:"children"`
	Parent   *File   `json:"parent"`
}

func part1() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	RootDir := &File{
		Name:     "/",
		Children: []*File{},
	}
	CurrentDir := RootDir

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		text := fileScanner.Text()
		if text == "$ ls" || text == "$ cd /" {
			continue
		} else if strings.Contains(text, " cd ") {
			splitted := strings.Split(text, " ")
			name := string(splitted[2])
			if name == ".." {
				CurrentDir = (CurrentDir.Parent)
				continue
			}

			newDir := File{
				Name:     name,
				Parent:   CurrentDir,
				Children: []*File{},
			}
			CurrentDir.Children = append(CurrentDir.Children, &newDir)
			CurrentDir = &newDir
		} else if strings.Contains(text, "dir") {
			continue
		} else {
			splitted := strings.Split(text, " ")
			name := string(splitted[1])
			size, _ := strconv.Atoi(splitted[0])
			newFile := File{
				Name: name,
				Size: size,
			}
			CurrentDir.Children = append(CurrentDir.Children, &newFile)
		}
	}

	sizee := GetSizeRecursive(RootDir)
	PrintRecursive(*RootDir, "")
	sum := strconv.Itoa(GetBigDirsRecursive(*RootDir, 0))
	bestDirSize := strconv.Itoa(GetBestDirToDelete(*RootDir, 0))
	fmt.Println("Sum: " + sum)
	fmt.Println("BestDir: " + bestDirSize)
	fmt.Println(sizee)
}

func GetBigDirsRecursive(file File, size int) int {
	if file.Children != nil {
		for _, v := range file.Children {
			size = GetBigDirsRecursive(*v, size)
		}
		if file.Size < 100000 {
			return size + file.Size
		} else {
			return size
		}
	} else {
		return size
	}
}

func GetBestDirToDelete(file File, size int) int {
	if file.Children != nil {
		for _, v := range file.Children {
			size = GetBestDirToDelete(*v, size)
		}
		if file.Size > 208860 && (file.Size < size || size == 0) {
			size = file.Size
		}
	}
	return size
}

func GetSizeRecursive(file *File) int {
	var size int
	for _, v := range file.Children {
		if v.Children != nil {
			size += GetSizeRecursive(v)
		} else {
			size += v.Size
		}
	}
	file.Size = size
	return size
}

func PrintRecursive(file File, tab string) {
	size := strconv.Itoa(file.Size)
	if file.Children != nil {
		fmt.Println(tab + "- " + file.Name + " (dir)" + " (" + size + ")")
		for _, v := range file.Children {
			PrintRecursive(*v, tab+"\t")
		}
	} else {
		fmt.Println(tab + "- " + file.Name + " (" + size + ")")
	}
}
