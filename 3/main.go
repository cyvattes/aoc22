package main

import (
	"bufio"
	"fmt"
	"os"
)

var priority = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
	"A": 26 + 1,
	"B": 26 + 2,
	"C": 26 + 3,
	"D": 26 + 4,
	"E": 26 + 5,
	"F": 26 + 6,
	"G": 26 + 7,
	"H": 26 + 8,
	"I": 26 + 9,
	"J": 26 + 10,
	"K": 26 + 11,
	"L": 26 + 12,
	"M": 26 + 13,
	"N": 26 + 14,
	"O": 26 + 15,
	"P": 26 + 16,
	"Q": 26 + 17,
	"R": 26 + 18,
	"S": 26 + 19,
	"T": 26 + 20,
	"U": 26 + 21,
	"V": 26 + 22,
	"W": 26 + 23,
	"X": 26 + 24,
	"Y": 26 + 25,
	"Z": 26 + 26,
}

func main() {
	fmt.Println(pt1())
	fmt.Println(pt2())
}

func pt1() int {
	var total int

	file := open("3/input")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := len(scanner.Text())
		a := toMap(scanner.Text()[:l/2])
		b := toMap(scanner.Text()[l/2:])

		for k, _ := range a {
			if b[k] {
				total += priority[k]
				break
			}
		}
	}

	return total
}

func pt2() int {
	var total int
	var groups []map[string]bool

	file := open("3/input")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		m := toMap(scanner.Text())
		groups = append(groups, m)
	}

	for i := 1; i <= len(groups); i++ {
		if i%3 == 0 {
			a := groups[i-3]
			b := groups[i-2]
			c := groups[i-1]

			for k, _ := range a {
				if b[k] && c[k] {
					total += priority[k]
					break
				}
			}
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

func toMap(word string) map[string]bool {
	m := make(map[string]bool)
	for _, v := range word {
		m[string(v)] = true
	}

	return m
}
