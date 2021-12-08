package day4

import (
	"2021-aoc/utils"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

type Square struct {
	row    int
	col    int
	value  int
	marked bool
}

type Board struct {
	squares    map[int]Square
	rowMatches [5][]int
	colMatches [5][]int
}

func RunA() {
	var winner *Board
	var draw int

	boards, draws := read()

	for {
		draw = Go(boards, draws)
		draws = draws[1:]
		winner, _ = checkWinner(boards)
		fmt.Println("draw", draw, "winner", winner)
		if winner != nil {
			break
		}
	}

	if winner != nil {
		printWinner(winner, draw)
	} else {
		panic(errors.New("no winner"))
	}
}

func RunB() {
	var winner *Board
	var draw int

	boards, draws := read()

	for {
		if len(draws) == 0 {
			break
		}
		draw = Go(boards, draws)
		draws = draws[1:]
		for {
			var i int
			winner, i = checkWinner(boards)
			if winner == nil {
				break
			}
			printWinner(winner, draw)
			boards = append(boards[:i], boards[i+1:]...)
		}
	}
}

func printWinner(winner *Board, draw int) {
	sum := 0
	for _, square := range winner.squares {
		if !square.marked {
			sum += square.value
		}
	}

	if sum == 0 {
		fmt.Println(winner.squares)
	}

	fmt.Println("sum: ", sum, "draw: ", draw, "sum*draw: ", sum*draw)
}

func read() ([]Board, []string) {
	var err error

	data, err := utils.ReadFile("./4/data.txt")
	handle(err)

	lines := strings.Split(data, "\n\n")
	draws := strings.Split(lines[0], ",")
	boards := make(
		[]Board,
		len(lines)-1,
	)

	// generate map number => row int, col int, marked bool
	for i, board := range lines[1:] {
		boards[i] = generateBoard(board)
	}
	return boards, draws
}

func Go(boards []Board, draws []string) int {
	var draw int
	var err error
	// for each board push into row[x], col[x] when number is called
	draw, err = strconv.Atoi(draws[0])
	handle(err)

	for i := 0; i < len(boards); i++ {
		board := &boards[i]
		if square, prs := board.squares[draw]; prs {
			board.rowMatches[square.row] = append(board.rowMatches[square.row], square.value)
			board.colMatches[square.col] = append(board.colMatches[square.col], square.value)
			square.marked = true
			board.squares[draw] = square
		}
	}

	return draw
}

func checkWinner(boards []Board) (*Board, int) {
	for i, board := range boards {
		for _, col := range board.colMatches {
			if len(col) == 5 {
				return &board, i
			}
		}
		for _, row := range board.rowMatches {
			if len(row) == 5 {
				return &board, i
			}
		}
	}

	return nil, -1
}

func generateBoard(board string) Board {
	boardOut := make(map[int]Square)
	for row, str := range strings.Split(board, "\n") {
		squareValues := regexp.MustCompile(`\s+`).Split(strings.TrimSpace(str), -1)
		for col, numStr := range squareValues {
			num, err := strconv.Atoi(numStr)
			handle(err)

			boardOut[num] = Square{row: row, col: col, marked: false, value: num}
		}
	}

	return Board{squares: boardOut}
}
