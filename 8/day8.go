package day8

import (
	"2021-aoc/utils"
	"errors"
	"fmt"
	"sort"
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
	inputLines, outputLines := read()
	sum := 0

	for i := range inputLines {
		inputs := inputLines[i]
		outputs := outputLines[i]

		counts := map[string]int{
			"a": 0,
			"b": 0,
			"c": 0,
			"d": 0,
			"e": 0,
			"f": 0,
			"g": 0,
		}

		// Determine characters for numbers 1,4,7,8
		// Store them in [10]int
		var signalPatterns [10]string

		for _, input := range inputs {
			num := getObviousNumber(input)
			if num != -1 {
				signalPatterns[num] = input
			}

			chars := strings.Split(input, "")
			for _, char := range chars {
				counts[char] = counts[char] + 1
			}
		}

		/**
		*  0
		* 5 1
		*  6
		* 4 2
		*  3
		**/
		var segments [7]string

		// Apply rules to get segment positions
		// 7 - 1 = position 0
		diff := Difference(
			strings.Split(signalPatterns[7], ""),
			strings.Split(signalPatterns[1], ""),
		)

		if len(diff) != 1 {
			panic(errors.New("could not determine 7 - 1"))
		}

		segments[0] = diff[0]

		segments[1] = signalPatterns[1]
		segments[2] = signalPatterns[1]

		// 4 - 1 = positions 5, 6
		diff = Difference(
			strings.Split(signalPatterns[4], ""),
			strings.Split(signalPatterns[1], ""),
		)

		if len(diff) != 2 {
			panic(errors.New("could not determine 4 - 1"))
		}

		segments[5] = strings.Join(diff, "")
		segments[6] = segments[5]

		/**
		0 => 8
		1 => 8
		2 => 9
		3 => 7
		4 => 4
		5 => 6
		6 => 7
		**/

		for char, count := range counts {
			switch count {
			case 8:
				if segments[0] != char {
					segments[1] = char
					segments[2] = strings.ReplaceAll(segments[2], segments[1], "")
				}
			case 4:
				segments[4] = char
			case 6:
				segments[5] = char
				segments[6] = strings.ReplaceAll(segments[6], segments[5], "")
			}
		}

		segments[3] = Difference(strings.Split("abcdefg", ""), segments[:])[0]

		// fmt.Println(segments)

		var result [4]string
		for i, output := range outputs {
			// println(output)
			num := getObviousNumber(output)
			var display [7]bool
			if num > -1 {
				result[i] = strconv.Itoa(num)
			} else {
				for i, seg := range segments {
					if strings.Contains(output, seg) {
						display[i] = true
					}
				}

				code := ""
				for i, on := range display {
					if on {
						code += strconv.Itoa(i)
					}
				}

				/**
				 * 0 => 6
				 * 2 => 6
				 * 3 => 6
				 * 5 => 6
				 * 6 => 7
				 * 9 => 7
				 **/
				switch code {
				case "012345":
					result[i] = "0"
				case "01346":
					result[i] = "2"
				case "01236":
					result[i] = "3"
				case "02356":
					result[i] = "5"
				case "023456":
					result[i] = "6"
				case "012356":
					result[i] = "9"
				default:
					panic("invalid code " + code)
				}
			}
		}

		val, err := strconv.Atoi(strings.Join(result[:], ""))
		handle(err)
		sum += val
		// fmt.Println(result)
	}

	fmt.Println(sum)
}

func getObviousNumber(in string) int {
	switch len(in) {
	case 2:
		return 1
	case 3:
		return 7
	case 4:
		return 4
	case 7:
		return 8
	default:
		return -1
	}
}

// Set Difference: A - B
func Difference(a, b []string) (diff []string) {
	m := make(map[string]bool)

	for _, item := range b {
		m[item] = true
	}

	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return
}

func RunA() {
	_, outputLines := read()

	count := 0
	for _, outputLine := range outputLines {
		for _, output := range outputLine {
			outputCharLength := len(output)
			knownLengths := []int{2, 3, 4, 7}

			/**
			* #1 => 2
			* #7 => 3
			* #4 => 4
			* #8 => 7
			**/

			idx := sort.Search(len(knownLengths), func(i int) bool {
				return knownLengths[i] >= outputCharLength
			})

			if idx < len(knownLengths) && knownLengths[idx] == outputCharLength {
				count++
			}
		}
	}

	fmt.Println("count: ", count)
}

func read() ([][]string, [][]string) {
	var err error

	data, err := utils.ReadFile("./8/data.txt")
	handle(err)

	lines := strings.Split(data, "\n")
	inputs := make([][]string, len(lines))
	outputs := make([][]string, len(lines))
	for i, line := range lines {
		inputAndOutput := strings.Split(line, " | ")
		inputs[i] = strings.Split(inputAndOutput[0], " ")
		outputs[i] = strings.Split(inputAndOutput[1], " ")
	}

	return inputs, outputs
}
