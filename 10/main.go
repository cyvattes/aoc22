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
	file := open("10/input")
	scanner := bufio.NewScanner(file)

	var reg = []int{1}
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), " ")
		x := reg[len(reg)-1]

		switch text[0] {
		case "noop":
			reg = append(reg, x)
		case "addx":
			n, _ := strconv.Atoi(text[1])
			reg = append(reg, []int{x, x + n}...)
		}
	}

	cycles := []int{20, 60, 100, 140, 180, 220}
	output := signal(reg, cycles)
	return output
	return 0
}

func pt2() int {
	file := open("10/input")
	scanner := bufio.NewScanner(file)

	var reg = []int{1}
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), " ")
		x := reg[len(reg)-1]

		switch text[0] {
		case "noop":
			reg = append(reg, x)
		case "addx":
			n, _ := strconv.Atoi(text[1])
			reg = append(reg, []int{x, x + n}...)
		}
	}

	draw(reg)
	return 0
}

func open(name string) *os.File {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err.Error())
	}

	return file
}

func signal(reg, cycle []int) int {
	var sum int

	for _, c := range cycle {
		sum += c * reg[c-1]
	}

	return sum
}

func draw(reg []int) {
	for i := 0; i < len(reg)-1; i++ {
		c := (i) % 40
		x := reg[i]

		if c == x-1 || c == x || c == x+1 {
			fmt.Printf("#")
		} else {
			fmt.Printf(".")
		}

		if c == 39 {
			fmt.Println("")
		}
	}
}
