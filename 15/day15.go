package day15

import (
	"2021-aoc/utils"
	"fmt"
	"strconv"
	"strings"
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

type Point struct {
	x    int
	y    int
	cost int
	risk int
	// path []*Point
}

func RunA() {
	grid := read()

	// for i := range grid {
	// 	for j := range grid[i] {
	// 		fmt.Print(grid[i][j].risk)
	// 	}
	// 	fmt.Print("\n\n")
	// }

	findShortestPaths(grid, nil, nil)
	score := grid[len(grid)-1][len(grid[0])-1].cost

	fmt.Println("bottom right", score)
}

func findShortestPaths(grid [][]*Point, curr *Point, prev *Point) {
	offsets := [4][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	// see if we have a lower cost
	if curr != nil {
		maybeLowestCost := curr.risk + prev.cost
		if curr.cost == -1 || maybeLowestCost < curr.cost {
			curr.cost = maybeLowestCost
		} else {
			return
		}
	} else {
		curr = grid[0][0]
		curr.cost = 0
	}

	for _, offset := range offsets {
		y := curr.y + offset[0]
		x := curr.x + offset[1]
		if y >= 0 && x >= 0 && y < len(grid) && x < len(grid[0]) {
			nextPoint := grid[y][x]
			findShortestPaths(grid, nextPoint, curr)
		}
	}
}

func RunB() {
}

func read() (grid [][]*Point) {
	var err error

	data, err := utils.ReadFile("./15/data.txt")
	handle(err)

	lines := strings.Split(data, "\n")
	grid = make([][]*Point, len(lines))

	for i := range lines {
		grid[i] = make([]*Point, len(lines[i]))
		for j := range lines[i] {
			v, err := strconv.Atoi(string(lines[i][j]))
			handle(err)
			grid[i][j] = &Point{y: i, x: j, risk: v, cost: -1}
		}
	}

	return
}
