package day2

import (
	"2021-aoc/utils"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func RunA() {
	data, err := utils.ReadFile("./2/data.txt")
	handle(err)

	horizontal := 0
	depth := 0

	for _, line := range strings.Split(data, "\n") {
		vals := strings.Split(line, " ")
		log.Println(line)
		dir := vals[0]
		inc, err := strconv.Atoi(vals[1])
		handle(err)

		switch dir {
		case "forward":
			horizontal += inc
		case "down":
			depth += inc
		case "up":
			depth -= inc
		default:
			panic(errors.New("invalid direction"))
		}
	}

	fmt.Println("Depth:", depth, "Horiz:", horizontal, "Product: ", depth*horizontal)
}

func RunB() {
	data, err := utils.ReadFile("./2/data.txt")
	handle(err)

	aim := 0
	horizontal := 0
	depth := 0

	for _, line := range strings.Split(data, "\n") {
		vals := strings.Split(line, " ")
		log.Println(line)
		dir := vals[0]
		inc, err := strconv.Atoi(vals[1])
		handle(err)

		switch dir {
		case "forward":
			horizontal += inc
			depth += aim * inc
		case "down":
			aim += inc
		case "up":
			aim -= inc
		default:
			panic(errors.New("invalid direction"))
		}
	}

	fmt.Println("Depth:", depth, "Horiz:", horizontal, "Product: ", depth*horizontal)
}
