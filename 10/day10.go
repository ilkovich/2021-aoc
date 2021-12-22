package day10

import (
	"2021-aoc/utils"
	"fmt"
	"sort"
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

func RunA() {
	charSequences := read()

	opens := "[({(<"
	closes := "])})>"

	scores := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	score := 0

	for i, chars := range charSequences {
		stack := make([]string, 0)
		fmt.Print(i, " ")

		for _, char := range chars {
			if strings.Contains(opens, char) {
				// fmt.Println("opens", char)
				stack = append(stack, char)
			} else if idx := strings.Index(closes, char); idx != -1 {
				if stack[len(stack)-1] == string(opens[idx]) {
					// fmt.Print(stack)
					stack = stack[:len(stack)-1]
					// fmt.Println("->", stack)
				} else {
					fmt.Println("score", score, "increment", scores[char])
					score += scores[char]
					break
					// panic("stop")
				}
			} else {
				panic("Unknown char " + char)
			}
		}
	}

	fmt.Println("Score:", score)
}

func RunB() {
	charSequences := read()

	closingScores := map[string]int{
		"(": 1,
		"[": 2,
		"{": 3,
		"<": 4,
	}

	opens := "[({(<"
	closes := "])})>"
	scores := make([]int, 0)

	for i, chars := range charSequences {
		stack := make([]string, 0)
		score := 0
		failed := false
		fmt.Print(i, " ")

		for _, char := range chars {
			if strings.Contains(opens, char) {
				// fmt.Println("opens", char)
				stack = append(stack, char)
			} else if idx := strings.Index(closes, char); idx != -1 {
				if stack[len(stack)-1] == string(opens[idx]) {
					// fmt.Print(stack)
					stack = stack[:len(stack)-1]
					// fmt.Println("->", stack)
				} else {
					failed = true
					break
				}
			} else {
				panic("Unknown char " + char)
			}
		}

		if !failed {
			for j := len(stack) - 1; j >= 0; j-- {
				char := stack[j]
				score = (score * 5) + closingScores[char]
			}
			scores = append(scores, score)
		}

		fmt.Println("score", score, strings.Join(stack, ""))
	}

	sort.Slice(scores, func(a, b int) bool {
		return scores[a] < scores[b]
	})

	idx := len(scores) / 2
	fmt.Println(len(scores), idx, "final score", scores[idx])
}

func read() (charSequences [][]string) {
	var err error

	data, err := utils.ReadFile("./10/data.txt")
	handle(err)

	lines := strings.Split(data, "\n")

	charSequences = make([][]string, len(lines))

	for i, line := range lines {
		charSequences[i] = strings.Split(line, "")
	}

	return
}
