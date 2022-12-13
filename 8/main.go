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
	file := open("8/input")
	scanner := bufio.NewScanner(file)

	var trees [][]int
	for scanner.Scan() {
		text := scanner.Text()
		vals := strings.Split(text, "")

		var row []int
		for _, v := range vals {
			i, _ := strconv.Atoi(v)
			row = append(row, i)
		}

		trees = append(trees, row)
	}

	return visible(trees)
}

func pt2() int {
	file := open("8/input")
	scanner := bufio.NewScanner(file)

	var trees [][]int
	for scanner.Scan() {
		text := scanner.Text()
		vals := strings.Split(text, "")

		var row []int
		for _, v := range vals {
			i, _ := strconv.Atoi(v)
			row = append(row, i)
		}

		trees = append(trees, row)
	}

	return scenic(trees)
}

func open(name string) *os.File {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err.Error())
	}

	return file
}

func visible(trees [][]int) int {
	arr := make([][]bool, len(trees))
	for i, _ := range arr {
		arr[i] = make([]bool, len(trees[i]))
	}

	for i, row := range trees {
		offset, length := 0, len(row)-1
		treeLH, treeRH := -1, -1
		treeTS, treeBS := -1, -1

		for offset < length {
			// LEFT & RIGHT
			if treeLH < row[offset] {
				arr[i][offset] = true
				treeLH = row[offset]
			}

			if treeRH < row[length-offset] {
				arr[i][length-offset] = true
				treeRH = row[length-offset]
			}

			// TOP & BOTTOM
			if treeTS < trees[offset][i] {
				arr[offset][i] = true
				treeTS = trees[offset][i]
			}

			if treeBS < trees[length-offset][i] {
				arr[length-offset][i] = true
				treeBS = trees[length-offset][i]
			}

			offset++
		}
	}

	var output int
	for _, row := range arr {
		for _, v := range row {
			if v {
				output++
			}
		}
	}

	return output
}

func scenic(trees [][]int) int {
	arr := make([][]int, len(trees))
	for i, _ := range arr {
		arr[i] = make([]int, len(trees[i]))
	}

	var max int
	rows, cols := len(trees), len(trees[0])
	for row := 1; row < rows-1; row++ {
		for col := 1; col < cols-1; col++ {
			tree := trees[row][col]

			DL := 0
			for x := col - 1; x >= 0; x-- {
				DL++
				if trees[row][x] >= tree {
					break
				}
			}

			DU := 0
			for y := row - 1; y >= 0; y-- {
				DU++
				if trees[y][col] >= tree {
					break
				}
			}

			DR := 0
			for x := col + 1; x < cols; x++ {
				DR++
				if trees[row][x] >= tree {
					break
				}
			}

			DD := 0
			for y := row + 1; y < rows; y++ {
				DD++
				if trees[y][col] >= tree {
					break
				}
			}

			dist := DL * DU * DR * DD
			arr[row][col] = dist
			if dist > max {
				max = dist
			}
		}
	}

	return max
}
