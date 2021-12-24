package day13

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

func RunA() {
	grid, folds := read()

	// p(grid)

	for _, fold := range folds {
		p(grid, false)
		grid = makeFold(fold, grid)
		p(grid, false)
	}

	p(grid, true)
}

func p(grid [][]bool, render bool) {
	count := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] {
				count++
				if render {
					fmt.Print("#")
				}
			} else {
				if render {
					fmt.Print(".")
				}
			}
		}
		if render {
			fmt.Print("\n")
		}
	}

	fmt.Println("count", count)
}

func makeFold(fold Fold, _grid [][]bool) (grid [][]bool) {
	var gridPiece [][]bool
	if fold.dir == X {
		fmt.Println("Folding on x=", fold.coord)

		gridPiece = make([][]bool, len(_grid))
		grid = make([][]bool, len(_grid))

		for i := range _grid {
			gridPiece[i] = _grid[i][fold.coord+1:]
			grid[i] = _grid[i][:fold.coord]

			for j, k := 0, len(gridPiece[i])-1; j < k; j, k = j+1, k-1 {
				gridPiece[i][j], gridPiece[i][k] = gridPiece[i][k], gridPiece[i][j]
			}
		}

		fmt.Println("----------")
		fmt.Println("<--> grid <-->")
		p(grid, false)
		fmt.Println("<--> gridPiece <-->")
		p(gridPiece, false)
		fmt.Println("----------")

		offset := len(grid[0]) - len(gridPiece[0])
		if offset < 0 {
			grid, gridPiece = gridPiece, grid
			offset *= -1
		}

		fmt.Println("offset", offset, len(grid[0]), len(gridPiece[0]))

		for i := range grid {
			for j := offset; j < len(grid[i]); j++ {
				grid[i][j] = grid[i][j] || gridPiece[i][j-offset]
			}
		}

		fmt.Println("og length", len(_grid[0]), "left length", len(grid[0]), "right length", len(gridPiece[0]))
		// p(gridPiece)
	} else if fold.dir == Y {
		// fmt.Println("Folding on y=", fold.coord)
		gridPiece = _grid[fold.coord+1:]
		grid = _grid[0:fold.coord]

		if len(gridPiece) <= len(grid) {
			for j, h := len(grid)-len(gridPiece), len(gridPiece)-1; j < len(grid); j, h = j+1, h-1 {
				for k := 0; k < len(grid[j]); k++ {
					// fmt.Println("j", j, "k", k, "h", h, grid[j][k], gridPiece[h][k])
					grid[j][k] = gridPiece[h][k] || grid[j][k]
				}
			}

			// p(grid)
		} else {
			panic("not implemented 2")
		}
	} else {
		panic("Invalid dir")
	}

	return
}

func RunB() {
}

type Direction int64

const (
	X Direction = iota
	Y
)

type Fold struct {
	coord int
	dir   Direction
}

func read() (grid [][]bool, foldInstructions []Fold) {
	var err error

	data, err := utils.ReadFile("./13/data.txt")
	handle(err)

	sections := strings.Split(data, "\n\n")
	coords := strings.Split(sections[0], "\n")
	folds := strings.Split(sections[1], "\n")

	foldInstructions = make([]Fold, len(folds))

	var maxX, maxY int
	parsedCoords := make([][2]int, len(coords))
	for i, coordLine := range coords {
		coords := strings.Split(coordLine, ",")
		x, err := strconv.Atoi(coords[0])
		handle(err)

		y, err := strconv.Atoi(coords[1])
		handle(err)

		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}

		parsedCoords[i][0] = x
		parsedCoords[i][1] = y
	}

	fmt.Println("maxX", maxX, "maxY", maxY)
	grid = make([][]bool, maxY+1)
	for i := range grid {
		// fmt.Println(i)
		grid[i] = make([]bool, maxX+1)
	}

	for _, coord := range parsedCoords {
		// fmt.Println(coord[0], coord[1])
		grid[coord[1]][coord[0]] = true
	}

	for i, foldLine := range folds {
		parts := strings.Split(foldLine, " ")
		// fmt.Println(strings.Join(parts, ","))
		rawDir := string(parts[2][0])
		coord, err := strconv.Atoi(parts[2][2:])
		handle(err)

		var dir Direction
		switch rawDir {
		case "y":
			dir = Y
		case "x":
			dir = X
		default:
			panic("invalid direction " + rawDir)
		}

		foldInstructions[i] = Fold{coord: coord, dir: dir}
	}

	return
}
