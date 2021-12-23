package day12

import (
	"2021-aoc/utils"
	"fmt"
	"strings"
	"unicode"
)

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func Run() {
	RunA()
	// RunB()
}

func RunA() {
	graph := read()

	for k, v := range graph {
		fmt.Print(k, " ", v.big, " ")
		for _, e := range v.edges {
			fmt.Print(e.name, " ")
		}

		fmt.Print("\n")
	}

	curr := graph["start"]
	visited := make([]string, 0)

	fmt.Print("\n")
	fmt.Print("\n")
	fmt.Print("\n")
	total := findPaths(curr, visited, "", false, false)
	fmt.Println(total)
}

func RunB() {
	graph := read()

	for k, v := range graph {
		fmt.Print(k, " ", v.big, " ")
		for _, e := range v.edges {
			fmt.Print(e.name, " ")
		}

		fmt.Print("\n")
	}

	curr := graph["start"]
	visited := make([]string, 0)

	fmt.Print("\n")
	fmt.Print("\n")
	fmt.Print("\n")
	total := findPaths(curr, visited, "", false, true)
	fmt.Println(total)
}

func findPaths(curr *Node, visited []string, path string, usedRepeat bool, allowRepeats bool) int {
	// allow visit?

	if curr.big {
		visited = append(visited, curr.name)
	} else {
		e := exists(visited, curr.name)
		if (allowRepeats && !usedRepeat) || !e {
			visited = append(visited, curr.name)
			if e {
				usedRepeat = true
			}
		} else {
			return 0
		}
	}

	if curr.name == "end" {
		path += ",end"
		fmt.Println(path)
		return 1
	}

	sum := 0
	for _, edge := range curr.edges {
		sum += findPaths(edge, visited, path+","+curr.name, usedRepeat, allowRepeats)
	}

	return sum
}

func exists(visited []string, s string) bool {
	for _, s1 := range visited {
		if s1 == s {
			return true
		}
	}

	return false
}

type Node struct {
	big   bool
	name  string
	edges []*Node
}

func read() (index map[string]*Node) {
	var err error

	data, err := utils.ReadFile("./12/data.txt")
	handle(err)

	lines := strings.Split(data, "\n")

	index = make(map[string]*Node, 0)

	for _, line := range lines {
		parts := strings.Split(line, "-")
		name0 := parts[0]
		name1 := parts[1]

		var node0 *Node
		var node1 *Node

		if v, e := index[name0]; e {
			node0 = v
		} else {
			node0 = &Node{name: name0, edges: make([]*Node, 0), big: unicode.IsUpper(rune(name0[0]))}
			index[name0] = node0
		}

		if v, e := index[name1]; e {
			node1 = v
		} else {
			node1 = &Node{name: name1, edges: make([]*Node, 0), big: unicode.IsUpper(rune(name1[0]))}
			index[name1] = node1
		}

		if node1.name != "start" {
			node0.edges = append(node0.edges, node1)
		}

		if node0.name != "start" {
			node1.edges = append(node1.edges, node0)
		}
	}

	return
}
