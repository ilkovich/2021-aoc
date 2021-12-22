package day11

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

type Octopus struct {
	energy int
	x      int
	y      int
}

func Run() {
	// RunA()
	RunB()
}

func RunA() {
	octopi := read()

	sum := 0
	for i := 0; i < 100; i++ {
		sum += step(&octopi)
		p(&octopi)
	}

	fmt.Println("Sum: ", sum)
}

func p(p_octopi *[][]Octopus) {
	octopi := *p_octopi
	for i := range octopi {
		fmt.Print(i, " ")
		for j := range octopi[i] {
			if octopi[i][j].energy == 0 {
				fmt.Print(string("\033[31m"), octopi[i][j].energy)
			} else {
				fmt.Print(string("\033[37m"), octopi[i][j].energy)
			}
		}
		fmt.Print(string("\033[37m"), "\n")
	}
	fmt.Print("\n")
}

func step(p_octopi *[][]Octopus) int {
	octopi := *p_octopi
	flashList := make([]*Octopus, 0)

	for i := range octopi {
		for j := range octopi[i] {
			octopi[i][j].energy++
			if octopi[i][j].energy == 10 {
				flashList = append(flashList, &octopi[i][j])
			}
		}
	}

	return flash(flashList, &octopi)
}

func flash(flashList []*Octopus, p_octopi *[][]Octopus) (total int) {
	if len(flashList) == 0 {
		return
	}

	octopi := *p_octopi

	offsets := [8][2]int{
		{-1, -1}, // NW
		{-1, 1},  // SW
		{-1, 0},  // W
		{0, -1},  // S
		{0, 1},   // N
		{1, -1},  // NE
		{1, 0},   // E
		{1, 1},   // SE
	}

	dimY := len(octopi)
	dimX := len(octopi[0])

	newFlashlist := make([]*Octopus, 0)
	for i := range flashList {
		total++
		octopus := flashList[i]
		octopus.energy = 0

		for _, offset := range offsets {
			x := octopus.x + offset[0]
			y := octopus.y + offset[1]

			if x >= 0 && y >= 0 && x < dimX && y < dimY {
				adjOctopus := &octopi[x][y]
				if adjOctopus.energy != 0 {
					adjOctopus.energy++
				}
				if adjOctopus.energy == 10 {
					// fmt.Println("Flashed", octopus.x, octopus.y, "Scheduling ", adjOctopus.x, adjOctopus.y)
					newFlashlist = append(newFlashlist, adjOctopus)
				}
			}
		}
	}

	// panic("done")
	return total + flash(newFlashlist, p_octopi)
}

func RunB() {
	octopi := read()
	count := len(octopi) * len(octopi[0])

	for i := 1; i < 1000; i++ {

		total := step(&octopi)

		if total == count {
			fmt.Println("step", i)
			break
		}
	}
}

func read() (octopi [][]Octopus) {
	var err error

	data, err := utils.ReadFile("./11/sample.txt")
	handle(err)

	lines := strings.Split(data, "\n")

	octopi = make([][]Octopus, len(lines))

	for i, line := range lines {
		octopi[i] = make([]Octopus, len(line))

		for j, char := range line {
			energy, err := strconv.Atoi(string(char))
			handle(err)

			octopi[i][j] = Octopus{energy: energy, x: i, y: j}
		}
	}

	return
}
