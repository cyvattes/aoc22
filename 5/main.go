package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var initial = map[int][]string{
	1: {"H", "C", "R"},
	2: {"B", "J", "H", "L", "S", "F"},
	3: {"R", "M", "D", "H", "J", "T", "Q"},
	4: {"S", "G", "R", "H", "Z", "B", "J"},
	5: {"R", "P", "F", "Z", "T", "D", "C", "B"},
	6: {"T", "H", "C", "G"},
	7: {"S", "N", "V", "Z", "B", "P", "W", "L"},
	8: {"R", "J", "Q", "G", "C"},
	9: {"L", "D", "T", "R", "H", "P", "F", "S"},
}

func main() {
	fmt.Println(pt1())
	fmt.Println(pt2())
}

func pt1() string {
	file := open("5/input")
	scanner := bufio.NewScanner(file)
	var stacks = map[int][]string{
		1: {"H", "C", "R"},
		2: {"B", "J", "H", "L", "S", "F"},
		3: {"R", "M", "D", "H", "J", "T", "Q"},
		4: {"S", "G", "R", "H", "Z", "B", "J"},
		5: {"R", "P", "F", "Z", "T", "D", "C", "B"},
		6: {"T", "H", "C", "G"},
		7: {"S", "N", "V", "Z", "B", "P", "W", "L"},
		8: {"R", "J", "Q", "G", "C"},
		9: {"L", "D", "T", "R", "H", "P", "F", "S"},
	}

	for scanner.Scan() {
		text := scanner.Text()
		text = strings.ReplaceAll(text, "move ", "")
		text = strings.ReplaceAll(text, " from ", " ")
		text = strings.ReplaceAll(text, " to ", " ")
		cols := strings.Fields(text)

		num, _ := strconv.Atoi(cols[0])
		from, _ := strconv.Atoi(cols[1])
		to, _ := strconv.Atoi(cols[2])

		for i := num; i > 0; i-- {
			l := len(stacks[from]) - 1
			x := stacks[from][l]
			stacks[from] = stacks[from][:l]
			stacks[to] = append(stacks[to], x)
		}
	}

	output := make([]string, len(stacks))
	for k, v := range stacks {
		output[k-1] = v[len(v)-1]
	}

	return strings.Join(output, "")
}

func pt2() string {
	file := open("5/input")
	scanner := bufio.NewScanner(file)
	var stacks = map[int][]string{
		1: {"H", "C", "R"},
		2: {"B", "J", "H", "L", "S", "F"},
		3: {"R", "M", "D", "H", "J", "T", "Q"},
		4: {"S", "G", "R", "H", "Z", "B", "J"},
		5: {"R", "P", "F", "Z", "T", "D", "C", "B"},
		6: {"T", "H", "C", "G"},
		7: {"S", "N", "V", "Z", "B", "P", "W", "L"},
		8: {"R", "J", "Q", "G", "C"},
		9: {"L", "D", "T", "R", "H", "P", "F", "S"},
	}

	var n int
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.ReplaceAll(text, "move ", "")
		text = strings.ReplaceAll(text, " from ", " ")
		text = strings.ReplaceAll(text, " to ", " ")
		cols := strings.Fields(text)

		num, _ := strconv.Atoi(cols[0])
		from, _ := strconv.Atoi(cols[1])
		to, _ := strconv.Atoi(cols[2])

		l := len(stacks[from]) - num
		if l < 0 {
			l = 0
		}

		if n < 10 {
			fmt.Printf("%d:%v | %d:%v\n", from, stacks[from], to, stacks[to])
		}

		x := stacks[from][l:]
		stacks[from] = stacks[from][:l]
		stacks[to] = append(stacks[to], x...)

		if n < 10 {
			fmt.Printf("%d:%v | %d:%v\n", from, stacks[from], to, stacks[to])
			n++
		}
	}

	output := make([]string, len(stacks))
	for k, v := range stacks {
		output[k-1] = v[len(v)-1]
	}

	return strings.Join(output, "")
}

func open(name string) *os.File {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err.Error())
	}

	return file
}
