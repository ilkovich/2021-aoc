package day9

import (
	"2021-aoc/utils"
	"fmt"
	"sort"
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

type Tile struct {
	x         int
	y         int
	height    int
	adjacents []*Tile
}

func RunB() {
	grid := read2()
	lowPoints := make([]*Tile, 0)

	offsets := [4][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	sum := 0
	for i := range grid {
		for j := range grid[i] {
			adjacents := make([]*Tile, 1)
			vals := make([]int, 1)
			adjacents[0] = &grid[i][j]
			vals[0] = grid[i][j].height
			for _, offset := range offsets {
				if i+offset[0] < len(grid) && i+offset[0] >= 0 && j+offset[1] < len(grid[i]) && j+offset[1] >= 0 {
					adjacents = append(adjacents, &grid[i+offset[0]][j+offset[1]])
					vals = append(vals, grid[i+offset[0]][j+offset[1]].height)
				}
			}

			grid[i][j].adjacents = adjacents

			// fmt.Println(adjacents)
			if min, _ := utils.Min(vals[:]); min == vals[0] && min != vals[1] {
				// fmt.Println("low point: ", min, adjacents)
				lowPoints = append(lowPoints, &grid[i][j])
				sum += grid[i][j].height + 1
			}
		}
	}

	println("len(lowPoints): ", len(lowPoints), "sum: ", sum)

	basins := make([][]Tile, 0)
	visited := make(map[string]bool)
	for _, seed := range lowPoints {
		if _, exists := visited[hash(*seed)]; !exists {
			basin := traverse([]*Tile{seed}, visited, make([]Tile, 0))
			basins = append(basins, basin)
		}
	}

	// println(len(basins))

	sort.SliceStable(basins, func(i, j int) bool {
		return len(basins[i]) > len(basins[j])
	})

	println("total basins", len(basins), "basins: ", len(basins[0]), len(basins[1]), len(basins[2]), len(basins[3]), "product: ", len(basins[0])*len(basins[1])*len(basins[2]))

	for i := range grid {
		for j := range grid[i] {
			if _, exists := visited[hash(grid[i][j])]; exists {
				fmt.Print(string("\033[31m"), grid[i][j].height)
			} else {
				fmt.Print(string("\033[37m"), grid[i][j].height)
			}
		}
		fmt.Print("\n")
	}
}

func hash(seed Tile) string {
	return strconv.Itoa(seed.x) + "|" + strconv.Itoa(seed.y)
}

func traverse(tiles []*Tile, visited map[string]bool, basin []Tile) []Tile {
	for _, tile := range tiles {
		// if tile is a 9 or we've already seen it
		if _, exists := visited[hash(*tile)]; exists || tile.height == 9 {
			// fmt.Println(tile.height)
			continue
		}

		// visit current tile
		// println("Visiting: ", tile.x, tile.y, tile.height)
		visited[hash(*tile)] = true

		// add tile to the basin
		basin = append(basin, *tile)
		// fmt.Println(len(basin))

		// recurse on the adjacents
		basin = traverse(tile.adjacents, visited, basin)
		// fmt.Println(len(basin))
	}

	return basin
}

func RunA() {
	grid := read()

	offsets := [4][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	sum := 0
	for i := range grid {
		for j, col := range grid[i] {
			adjacents := make([]int, 1)
			adjacents[0] = col
			for _, offset := range offsets {
				if i+offset[0] < len(grid) && i+offset[0] >= 0 && j+offset[1] < len(grid[i]) && j+offset[1] >= 0 {
					adjacents = append(adjacents, grid[i+offset[0]][j+offset[1]])
				}
			}

			if min, _ := utils.Min(adjacents[:]); min == adjacents[0] && min != adjacents[1] {
				fmt.Println("low point: ", min, adjacents)
				sum += min + 1
			}
		}
	}

	fmt.Println(sum)
}

var grid [][]Tile

func read2() [][]Tile {
	var err error

	data, err := utils.ReadFile("./9/data.txt")
	handle(err)

	lines := strings.Split(data, "\n")

	grid = make([][]Tile, len(lines))
	for i := range grid {
		grid[i] = make([]Tile, len(lines[i]))
		for k, char := range strings.Split(lines[i], "") {
			height, err := strconv.Atoi(char)
			handle(err)

			grid[i][k] = Tile{height: height, x: i, y: k}
		}
	}

	return grid
}

func read() [][]int {
	var err error

	data, err := utils.ReadFile("./9/data.txt")
	handle(err)

	lines := strings.Split(data, "\n")

	grid := make([][]int, len(lines))
	for i := range grid {
		grid[i] = make([]int, len(lines[i]))
		for k, char := range strings.Split(lines[i], "") {
			grid[i][k], err = strconv.Atoi(char)
			handle(err)
		}
	}

	return grid
}
