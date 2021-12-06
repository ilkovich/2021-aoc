package day1

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

func RunA() {
	data, err := utils.ReadFile("./1/data.txt")
	handle(err)

	count, err := getIncreases(strings.Split(strings.TrimSpace(data), "\n"))
	handle(err)

	fmt.Println("Result Part 1: ", count)
}

func RunB() {
	var err error

	data, err := utils.ReadFile("./1/data.txt")
	handle(err)

	entries := strings.Split(strings.TrimSpace(data), "\n")
	numEntries := make([]int, len(entries))
	grouped := make([]int, len(entries)-2)

	for i, entry := range entries {
		val, err := strconv.Atoi(entry)
		handle(err)

		numEntries[i] = val
	}

	for i := 0; i+2 < len(entries); i++ {
		grouped[i] = numEntries[i] + numEntries[i+1] + numEntries[i+2]
	}

	count, err := getIncreases2(grouped)
	handle(err)

	fmt.Println("Result Part 2: ", count)
}

func getIncreases2(entries []int) (int, error) {
	count := 0
	prev := 0

	for _, num := range entries {
		if prev > 0 && num > prev {
			count++
		}

		prev = num
	}

	return count, nil
}

func getIncreases(entries []string) (int, error) {
	count := 0
	prev := 0

	for _, entry := range entries {
		num, err := strconv.Atoi(entry)
		if err != nil {
			return -1, err
		}

		if prev > 0 && num > prev {
			count++
		}

		prev = num
	}

	return count, nil
}
