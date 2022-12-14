package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type POS struct {
	X int
	Y int
}

var (
	Knots []POS
	M     map[POS]bool

	MaxX = 6
	MaxY = 6
)

func main() {
	fmt.Println(pt1())
	fmt.Println(pt2())
}

func pt1() int {
	file := open("9/input")
	scanner := bufio.NewScanner(file)

	Knots = []POS{
		{0, 0},
		{0, 0},
	}
	M = make(map[POS]bool)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), " ")
		d := text[0]
		n, _ := strconv.Atoi(text[1])

		move(d, n)
	}

	var output int
	for _, v := range M {
		if v {
			output++
		}
	}

	return output
}

func pt2() int {
	file := open("9/input")
	scanner := bufio.NewScanner(file)

	Knots = []POS{
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
	}
	M = make(map[POS]bool)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), " ")
		d := text[0]
		n, _ := strconv.Atoi(text[1])

		move(d, n)
	}

	var output int
	for _, v := range M {
		if v {
			output++
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

func move(d string, n int) {
	for n > 0 {
		switch d {
		case "U":
			update(0, 1)
		case "D":
			update(0, -1)
		case "R":
			update(1, 0)
		case "L":
			update(-1, 0)
		}

		n--
		last := len(Knots) - 1
		tail := Knots[last]
		M[tail] = true

		//print()
	}
}

func update(dx, dy int) {
	Knots[0].X += dx
	Knots[0].Y += dy

	//fmt.Printf("%v: <%d, %d>\n", Knots[0], dx, dy)

	for i := 1; i < len(Knots); i++ {
		k := follow(Knots[i-1], Knots[i])

		//fmt.Printf("%d: %v -> %v\n", i, Knots[i], k)
		Knots[i] = k
	}
}

func follow(previous, current POS) POS {
	p := POS{current.X, current.Y}
	dx := previous.X - current.X
	dy := previous.Y - current.Y

	//fmt.Printf("[%d, %d]\n", dx, dy)

	// INLINE DIFF
	if math.Abs(float64(dx)) > 1 && dy == 0 {
		p.X += reduce(dx)
		return p
	} else if math.Abs(float64(dy)) > 1 && math.Abs(float64(dx)) == 0 {
		p.Y += reduce(dy)
		return p
	}

	// DIAGONAL DIFF
	if math.Abs(float64(dx)) > 1 && math.Abs(float64(dy)) == 1 {
		p.X += reduce(dx)
		p.Y += dy
		return p
	} else if math.Abs(float64(dy)) > 1 && math.Abs(float64(dx)) == 1 {
		p.Y += reduce(dy)
		p.X += dx
		return p
	} else if math.Abs(float64(dy)) > 1 && math.Abs(float64(dx)) > 1 {
		p.Y += reduce(dy)
		p.X += reduce(dy)
		return p
	}

	return p
}

func reduce(i int) int {
	if i < -1 {
		return i + 1
	}

	if i > 1 {
		return i - 1
	}

	return i
}

func print() {
	arr := make([][]string, MaxY)
	for row := range arr {
		arr[row] = make([]string, MaxX)
		for col := range arr[row] {
			arr[row][col] = "."
		}
	}

	for i := len(Knots) - 1; i >= 0; i-- {
		k := Knots[i]
		v := strconv.Itoa(i)
		if v == "0" {
			v = "H"
		}

		arr[k.Y][k.X] = v
	}

	for y := MaxY - 1; y >= 0; y-- {
		fmt.Println(strings.Join(arr[y], " "))
	}

	fmt.Println()
}
