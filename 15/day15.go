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
	// RunA()
	RunB()
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

	end := grid[len(grid)-1][len(grid[0])-1]
	findShortestPaths(grid, nil, nil, end)
	score := end.cost

	fmt.Println("bottom right", score)
}

func RunB() {
	grid := read()

	largerGrid := make([][]*Point, len(grid)*5)
	for i := range largerGrid {
		largerGrid[i] = make([]*Point, len(grid[0])*5)
		for j := range largerGrid[i] {
			if i < len(grid) && j < len(grid[i]) {
				largerGrid[i][j] = grid[i][j]
			} else {
				i2 := i % len(grid)
				j2 := j % len(grid[i2])
				adjustment := i/len(grid) + j/len(grid[i2])
				risk := (grid[i2][j2].risk + adjustment) % 9
				if risk == 0 {
					risk = 9
				}
				largerGrid[i][j] = &Point{x: j, y: i, risk: risk, cost: -1}
			}
		}
	}

	// P(largerGrid)

	end := largerGrid[len(largerGrid)-1][len(largerGrid[0])-1]

	findShortestPaths(largerGrid, nil, nil, end)

	score := end.cost

	fmt.Println("bottom right", score)
}

func P(largerGrid [][]*Point) {
	for i := range largerGrid {
		if i == 0 {
			fmt.Print("  ")
			for k := 0; k < len(largerGrid[i]); k++ {
				fmt.Print(k % 10)
			}
			fmt.Print("\n\n")
		}

		fmt.Print(i%10, " ")
		for j := range largerGrid[i] {
			fmt.Print(largerGrid[i][j].risk)
		}
		fmt.Print("\n")
	}

	fmt.Print("\n")
}

var visited int

func findShortestPaths(grid [][]*Point, curr *Point, prev *Point, end *Point) {
	visited++
	if visited%100000000 == 0 {
		fmt.Println("visited: ", visited, "low: ", end.cost)
	}

	offsets := [4][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	// see if we have a lower cost
	if curr != nil {
		maybeLowestCost := curr.risk + prev.cost
		if curr.cost == -1 || maybeLowestCost < curr.cost {
			curr.cost = maybeLowestCost
			distanceToEnd := end.x - curr.x + end.y - curr.y

			if end.cost != -1 && curr.cost+distanceToEnd >= end.cost {
				return
			}
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
			findShortestPaths(grid, nextPoint, curr, end)
		}
	}
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
