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

func RunB() {
}

func RunA() {
	timers := read(true)
	days := 256

	for i := 0; i < days; i++ {
		timers = step(timers)
		fmt.Println(i)
	}

	fmt.Println(len(timers))
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

func read(skipDiags bool) []int {
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
