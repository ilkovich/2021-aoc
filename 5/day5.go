package day5

import (
	"2021-aoc/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const SIZE = 1000

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

type Board struct {
	squares [SIZE][SIZE]int
}

func RunA() {
	board := read()

	count := 0
	for k := 0; k < SIZE; k++ {
		for j := 0; j < SIZE; j++ {
			if board.squares[k][j] >= 2 {
				count++
			}
		}
	}

	fmt.Println("count: ", count)
}

func read() Board {
	var err error

	data, err := utils.ReadFile("./5/data.txt")
	handle(err)

	lines := strings.Split(data, "\n")
	board := Board{}
	for _, line := range lines {
		coords := strings.Split(line, " -> ")
		start := coords[0]
		end := coords[1]

		coordsArray := [2][2]int{
			getCoords(start),
			getCoords(end),
		}

		sort.SliceStable(coordsArray[:], func(x, y int) bool {
			p1 := coordsArray[x]
			p2 := coordsArray[y]

			if p1[0] == p2[0] {
				return p1[1] < p2[1]
			} else {
				return p1[0] < p2[0]
			}
		})

		x1, y1 := coordsArray[0][0], coordsArray[0][1]
		x2, y2 := coordsArray[1][0], coordsArray[1][1]

		fmt.Println(x1, y1, x2, y2)

		if x1 != x2 && y1 != y2 {
			continue
		}

		for i := x1; i <= x2; i++ {
			for j := y1; j <= y2; j++ {
				board.squares[j][i]++
			}
		}
	}

	return board
}

func getCoords(coordIn string) [2]int {
	var err error
	var x, y int

	coordArr := strings.Split(coordIn, ",")

	x, err = strconv.Atoi(coordArr[0])
	handle(err)

	y, err = strconv.Atoi(coordArr[1])
	handle(err)

	return [2]int{x, y}
}
