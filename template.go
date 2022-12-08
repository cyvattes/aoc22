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
	file := open("N/input")
	scanner := bufio.NewScanner(file)

	var output int
	for scanner.Scan() {
		// TASK
	}

	return output
}

func pt2() int {
	file := open("N/input")
	scanner := bufio.NewScanner(file)

	var output int
	for scanner.Scan() {
		// TASK
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
