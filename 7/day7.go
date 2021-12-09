package day7

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
	// fuel(1, 5)
}

func RunB() {
	positions := read()
	distanceArr := createDistanceArray(positions)
	sums := make([]int, len(distanceArr))

	for i, point := range distanceArr {
		for _, position := range positions {
			sums[i] += fuel(position, point)
		}
	}

	val, pos := min(sums)

	fmt.Println("val: ", val, "pos: ", pos)
}

func fuel(start int, end int) int {
	dist := float64(absInt(start - end))
	fuel := dist * ((dist + 1) / 2)
	// fmt.Println(start, end, fuel)
	return int(fuel)
}

func RunA() {
	positions := read()

	sums := make([]int, len(positions))

	for i, position := range positions {
		fmt.Println(i, positions)
		positionsPrime := make([]int, len(positions))
		copy(positionsPrime, positions)
		distanceArr := append(positionsPrime[:i], positionsPrime[i+1:]...)
		for _, point := range distanceArr {
			sums[i] += absInt(position - point)
		}
	}

	val, pos := min(sums)

	fmt.Println("val: ", val, "pos: ", pos)
}

func createDistanceArray(positions []int) []int {
	min, _ := min(positions)
	max, _ := max(positions)

	results := make([]int, max-min+1)

	for i, j := 0, min; i <= max; {
		results[i] = j
		i++
		j++
	}

	return results
}

func min(arr []int) (int, int) {
	var ptr *int
	var pos int

	for i, val := range arr {
		if ptr == nil || *ptr > val {
			nextVal := val
			ptr = &nextVal
			pos = i
		}
	}

	return *ptr, pos
}

func max(arr []int) (int, int) {
	var ptr *int
	var pos int

	for i, val := range arr {
		if ptr == nil || *ptr < val {
			nextVal := val
			ptr = &nextVal
			pos = i
		}
	}

	return *ptr, pos
}

func absInt(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

func read() []int {
	var err error

	data, err := utils.ReadFile("./7/data.txt")
	handle(err)

	timerStrings := strings.Split(data, ",")
	timers := make([]int, len(timerStrings))

	for i, timer := range timerStrings {
		timers[i], err = strconv.Atoi(timer)
		handle(err)
	}

	return timers
}
