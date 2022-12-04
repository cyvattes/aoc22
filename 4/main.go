package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(pt1())
	fmt.Println(pt2())
}

func pt1() int {
	var total int

	file := open("4/input")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), ",")
		a := strings.Split(text[0], "-")
		b := strings.Split(text[1], "-")

		if overlapAll(a, b) {
			total++
		}
	}

	return total
}

func pt2() int {
	var total int

	file := open("4/input")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), ",")
		a := strings.Split(text[0], "-")
		b := strings.Split(text[1], "-")

		if overlapAny(a, b) {
			total++
		}
	}

	return total
}

func open(name string) *os.File {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err.Error())
	}

	return file
}

func overlapAll(a, b []string) bool {
	a1, _ := strconv.Atoi(a[0])
	a2, _ := strconv.Atoi(a[1])
	b1, _ := strconv.Atoi(b[0])
	b2, _ := strconv.Atoi(b[1])

	return (a1 <= b1 && a2 >= b2) ||
		(a1 >= b1 && a2 <= b2)
}

func overlapAny(a, b []string) bool {
	a1, _ := strconv.Atoi(a[0])
	a2, _ := strconv.Atoi(a[1])
	b1, _ := strconv.Atoi(b[0])
	b2, _ := strconv.Atoi(b[1])

	return (a1 <= b1 && a2 >= b1) ||
		(a1 <= b2 && a2 >= b2) ||
		(b1 <= a1 && b2 >= a1) ||
		(b1 <= a2 && b2 >= a2)
}
