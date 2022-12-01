package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	max, bags := ElfWithLargestBag()
	top := Top3Elves(bags)

	fmt.Println(max)
	fmt.Println(top)
}

func ElfWithLargestBag() (int, map[int]int) {
	file, err := os.Open("1/input")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(file)

	bags := make(map[int]int)
	i, max := 1, 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			if bags[i] > max {
				max = bags[i]
			}
			i++
			continue
		}

		weight, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("invalid weight")
		}

		bags[i] += weight
	}

	return max, bags
}

func Top3Elves(bags map[int]int) int {
	var weights []int
	for _, v := range bags {
		weights = append(weights, v)
	}

	sort.Ints(weights)
	l := len(weights) - 1
	top := weights[l] + weights[l-1] + weights[l-2]

	return top
}
