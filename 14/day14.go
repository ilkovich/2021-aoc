package day14

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
	polymer, rules := read()

	newPolymer := ""

	occurrences := make([]int, 26)
	for step := 0; step < 10; step++ {
		for i := 0; i < len(polymer)-1; i++ {
			j := polymer[i : i+2]

			for _, rule := range rules {
				if rule.from == j {
					j = string(j[0]) + rule.to + string(j[1])
				}
			}

			if len(newPolymer) == 0 {
				newPolymer = j
			} else {
				newPolymer = newPolymer[0:len(newPolymer)-1] + j
			}
		}
		polymer = newPolymer
		newPolymer = ""
		fmt.Println(step, len(polymer))
	}

	for _, char := range polymer {
		occurrences[char-'A']++
	}

	fmt.Println(occurrences)
	sort.SliceStable(occurrences, func(i, j int) bool {
		return occurrences[i] < occurrences[j]
	})

	min := 0
	for i := 0; i < len(occurrences); i++ {
		if occurrences[i] > 0 {
			min = occurrences[i]
			break
		}
	}

	max := 0
	for i := len(occurrences) - 1; i >= 0; i-- {
		if occurrences[i] > 0 {
			max = occurrences[i]
			break
		}
	}

	fmt.Println(occurrences)

	fmt.Println("answer: ", max-min)
}

func RunB() {
	polymer, rules := read()
	var occurrences []int

	polymerMap := make(map[string]int, 0)
	for i := 0; i < len(polymer)-1; i++ {
		increment(polymerMap, polymer[i:i+2], 1)
	}

	var step int
	for step = 0; step < 40; step++ {
		newMap := make(map[string]int, 0)
		for _, rule := range rules {
			if v, e := polymerMap[rule.from]; e && v != 0 {
				newMap[rule.from] -= v
				increment(newMap, rule.from[0:1]+rule.to, v)
				increment(newMap, rule.to+rule.from[1:2], v)
			}
		}

		for k, v := range newMap {
			polymerMap[k] += v
		}
	}

	fmt.Println(step)
	fmt.Println("->", polymerMap)

	occurrences = make([]int, 26)
	occurrences[polymer[len(polymer)-1]-'A']++
	for k, v := range polymerMap {
		occurrences[k[0]-'A'] += v
	}
	sum := 0
	for i := range occurrences {
		sum += occurrences[i]
	}
	fmt.Println("-> length", sum)
	fmt.Println("->", occurrences)

	sort.SliceStable(occurrences, func(i, j int) bool {
		return occurrences[i] < occurrences[j]
	})

	min := 0
	for i := 0; i < len(occurrences); i++ {
		if occurrences[i] > 0 {
			min = occurrences[i]
			break
		}
	}

	max := 0
	for i := len(occurrences) - 1; i >= 0; i-- {
		if occurrences[i] > 0 {
			max = occurrences[i]
			break
		}
	}

	fmt.Println("answer: ", max-min)
}

func clone(polymerMap map[string]int) (b map[string]int) {
	b = make(map[string]int, 0)
	for k, v := range polymerMap {
		b[k] = v
	}

	return
}

func increment(polymerMap map[string]int, s string, v int) {
	polymerMap[s] += v
}

type Rule struct {
	from string
	to   string
}

func read() (polymer string, rules []Rule) {
	var err error

	data, err := utils.ReadFile("./14/data.txt")
	handle(err)

	sections := strings.Split(data, "\n\n")
	polymer = sections[0]
	_rules := strings.Split(sections[1], "\n")

	rules = make([]Rule, len(_rules))

	for i, _rule := range _rules {
		parts := strings.Split(_rule, " -> ")
		rules[i] = Rule{from: parts[0], to: parts[1]}
	}

	return
}
