package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println(pt1())
	fmt.Println(pt2())
}

func pt1() int {
	file := open("6/input")
	scanner := bufio.NewScanner(file)

	var output int
	for scanner.Scan() {
		text := []rune(scanner.Text())
		for i := 3; i < len(text); i++ {
			if unique(text[i-3 : i+1]) {
				return i + 1
			}
		}
	}

	return output
}

func pt2() int {
	file := open("6/input")
	scanner := bufio.NewScanner(file)

	var output int
	for scanner.Scan() {
		text := []rune(scanner.Text())
		for i := 13; i < len(text); i++ {
			if unique(text[i-13 : i+1]) {
				return i + 1
			}
		}
	}

	return output
}

func open(name string) *os.File {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err.Error())
	}

	return file
}

func unique(text []rune) bool {
	m := make(map[rune]bool)
	for _, v := range text {
		if m[v] {
			return false
		}

		m[v] = true
	}

	return true
}
