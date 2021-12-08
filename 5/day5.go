package day5

import (
	"2021-aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

const SIZE = 1000

// const SIZE = 10

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

type Board struct {
	squares [SIZE][SIZE]int
}

func RunB() {
	board := read(false)

	print(board)
}

func RunA() {
	board := read(true)

	print(board)
}

func print(board Board) {
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

func read(skipDiags bool) Board {
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

		x1, y1 := coordsArray[0][0], coordsArray[0][1]
		x2, y2 := coordsArray[1][0], coordsArray[1][1]

		var xinc func(i int) int
		var yinc func(i int) int

		if x1 > x2 {
			xinc = func(i int) int {
				return i - 1
			}
		} else {
			xinc = func(i int) int {
				return i + 1
			}
		}

		if y1 > y2 {
			yinc = func(i int) int {
				return i - 1
			}
		} else {
			yinc = func(i int) int {
				return i + 1
			}
		}

		// fmt.Println(x1, y1, x2, y2)

		if x1 != x2 && y1 != y2 {
			if skipDiags {
				continue
			}

			for i, j := x1, y1; i != xinc(x2) && j != yinc(y2); i, j = xinc(i), yinc(j) {
				board.squares[j][i]++
			}
		} else {
			for i := x1; i != xinc(x2); i = xinc(i) {
				for j := y1; j != yinc(y2); j = yinc(j) {
					board.squares[j][i]++
				}
			}
		}

		// for i := 0; i < SIZE; i++ {
		// 	fmt.Println(board.squares[i])
		// }
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
