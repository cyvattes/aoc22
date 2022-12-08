package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	name     string
	size     int
	parent   *Node
	children []*Node
}

func main() {
	fmt.Println(pt1())
	fmt.Println(pt2())
}

func pt1() int {
	file := open("7/input")
	//file := open("7/test")
	scanner := bufio.NewScanner(file)

	var filesystem = Node{
		name:     "/",
		size:     0,
		parent:   nil,
		children: []*Node{},
	}

	var current = &filesystem
	for scanner.Scan() {
		text := scanner.Text()
		current = current.parse(text)
	}

	filesystem.resize()
	return filesystem.under(100000)
}

func pt2() int {
	file := open("7/input")
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

func (n *Node) parse(text string) *Node {
	cmd := strings.Split(text, " ")
	switch cmd[0] {
	case "$":
		if cmd[1] == "ls" {
			return n
		}

		if cmd[2] == ".." {
			if n.parent != nil {
				return n.parent
			}

			return n
		}

		return n.cd(cmd[2])
	case "dir":
		dir := Node{cmd[1], 0, n, []*Node{}}
		n.addChild(&dir)
		return n
	default:
		size, _ := strconv.Atoi(cmd[0])
		file := Node{cmd[1], size, n, nil}
		n.addChild(&file)
		return n
	}
}

func (n *Node) cd(dir string) *Node {
	for _, c := range n.children {
		if c.name == dir {
			return c
		}
	}

	return n
}

func (n *Node) addChild(c *Node) {
	n.children = append(n.children, c)
}

func (n *Node) resize() int {
	for _, c := range n.children {
		n.size += c.resize()
	}

	return n.size
}

func (n *Node) under(limit int) int {
	var size int
	if len(n.children) > 0 && n.size <= limit {
		size += n.size
	}

	for _, c := range n.children {
		size += c.under(limit)
	}

	return size
}

func (n *Node) print(prepend string) {
	fmt.Printf("%s-%s: %d\n", prepend, n.name, n.size)
	for _, c := range n.children {
		c.print(prepend + "\t")
	}
}
