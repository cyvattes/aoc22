package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(pt1())
	fmt.Println(pt2())
}

func pt1() int {
	var total int

	file := open("2/input")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		vals := strings.Split(scanner.Text(), " ")
		a, b := vals[0], vals[1]
		points := score(a, b)

		total += points
	}

	return total
}

func pt2() int {
	var total int

	file := open("2/input")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		vals := strings.Split(scanner.Text(), " ")
		a, out := vals[0], vals[1]
		b := response(a, out)
		points := score(a, b)

		total += points
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

func score(a, b string) int {
	points := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	return winner(a, b) + points[b]
}

func winner(a, b string) int {
	switch a + b {
	case "AY", "BZ", "CX":
		return 6
	case "AZ", "BX", "CY":
		return 0
	default:
		return 3
	}
}

func response(a, b string) string {
	if b == "Y" {
		switch a {
		case "A":
			return "X"
		case "B":
			return "Y"
		case "C":
			return "Z"
		}
	}

	switch a + b {
	case "BX", "CZ":
		return "X"
	case "CX", "AZ":
		return "Y"
	case "AX", "BZ":
		return "Z"
	}

	return ""
}
