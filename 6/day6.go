package day6

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

func RunB() {
	timers := read()
	days := 256
	vals := make(map[int]int, 8)
	newVals := make(map[int]int, 8)

	// initialize
	for _, timer := range timers {
		vals[timer]++
	}

	for i := 0; i < days; i++ {
		newVals[8] = vals[0]
		newVals[7] = vals[8]
		newVals[6] = vals[0] + vals[7]
		newVals[5] = vals[6]
		newVals[4] = vals[5]
		newVals[3] = vals[4]
		newVals[2] = vals[3]
		newVals[1] = vals[2]
		newVals[0] = vals[1]

		// fmt.Println(vals, " -> ", newVals)
		vals = newVals
		newVals = make(map[int]int, 8)
	}

	sum := 0
	for i := 0; i <= 8; i++ {
		sum += vals[i]
	}

	fmt.Println("Sum: ", sum)
}

func RunA() {
	timers := read()
	days := 24

	for i := 0; i < days; i++ {
		timers = step(timers)
		fmt.Println(i, timers)
	}

}

func step(timers []int) []int {
	for i, timer := range timers {
		if timer == 0 {
			timers = append(timers, 8)
			timers[i] = 6
		} else {
			timers[i]--
		}
	}

	return timers
}

func read() []int {
	var err error

	data, err := utils.ReadFile("./6/data.txt")
	handle(err)

	timerStrings := strings.Split(data, ",")
	timers := make([]int, len(timerStrings))

	for i, timer := range timerStrings {
		timers[i], err = strconv.Atoi(timer)
		handle(err)
	}

	return timers
}
