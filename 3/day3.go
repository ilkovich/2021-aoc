package day3

import (
	"2021-aoc/utils"
	"errors"
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
	data, err := utils.ReadFile("./3/sample.txt")
	handle(err)
	fmt.Println(data)

	lines := strings.Split(data, "\n")
	counts := make([]int, len(lines[0]))

	for _, line := range lines {
		for i, char := range line {
			fmt.Println(i, char)
			switch char {
			case '0':
				counts[i]--
			case '1':
				counts[i]++
			default:
				panic(errors.New("invalid digit"))
			}
		}
	}

	gamma := ""
	epsilon := ""

	for i, count := range counts {
		if count > 0 {
			gamma += "1"
			epsilon += "0"
		} else if count < 0 {
			gamma += "0"
			epsilon += "1"
		} else {
			fmt.Println(counts, i, count)
			panic(errors.New("0 count"))
		}
	}

	gammaInt, err := strconv.ParseInt(gamma, 2, 64)
	handle(err)

	epsilonInt, err := strconv.ParseInt(epsilon, 2, 64)
	handle(err)

	fmt.Println(
		"Epsilon Str:", epsilon,
		"Epsilon Int: ", epsilonInt,
		"Gamma Str:", gamma,
		"Gamma Int: ", gammaInt,
		"gamma*epsilon: ", epsilonInt*gammaInt)
}

func RunB() {
	data, err := utils.ReadFile("./3/data.txt")
	handle(err)
	fmt.Println(data)

	ogr, csc := getRatings(strings.Split(data, "\n"), 0)
	ogrInt, err := strconv.ParseInt(ogr[0], 2, 64)
	handle(err)
	cscInt, err := strconv.ParseInt(csc[0], 2, 64)
	handle(err)

	fmt.Println(
		"ogr: ", ogr,
		"ogrInt: ", ogrInt,
		"csc: ", csc,
		"cscInt: ", cscInt,
		"ogr*csc: ", ogrInt*cscInt,
	)
}

func getRatings(lines []string, i int) ([]string, []string) {
	if len(lines) <= 1 {
		return lines, lines
	}

	temp := [2][]string{
		make([]string, 0),
		make([]string, 0),
	}

	for _, line := range lines {
		switch line[i] {
		case '0':
			temp[0] = append(temp[0], line)
		case '1':
			temp[1] = append(temp[1], line)
		default:
			panic(errors.New("invalid digit"))
		}
	}

	var ogr []string
	var csr []string

	if len(temp[0]) > len(temp[1]) {
		ogr, _ = getRatings(temp[0], i+1)
		_, csr = getRatings(temp[1], i+1)
	} else {
		ogr, _ = getRatings(temp[1], i+1)
		_, csr = getRatings(temp[0], i+1)
	}

	return ogr, csr
}
