package day17

import (
	"2021-aoc/utils"
	"fmt"
	"regexp"
	"strconv"
)

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func Run() {
	RunA()
	// RunB()
}

func RunA() {
	x, y, grid := read()

	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)

	fmt.Print("  ")
	for i := range grid[0] {
		fmt.Print(i % 10)
	}

	fmt.Println()

	for i := range grid {
		fmt.Print(i%10, " ")
		for j := range grid[i] {
			fmt.Print(grid[i][j])
		}
		fmt.Println()
	}

	// velocity := [2]int{7, 2}
	// velocity := [2]int{6, 3}
	// velocity := [2]int{9, 0}
	// velocity := [2]int{17, -4}

	var minY int
	var landedCount int
	for j, possible := -y[1], true; possible; j++ {
		for i := 0; i <= x[1]; i++ {
			// fmt.Println("i", i, "j", j)
			landed, curr, velocity, _minY := run(x, y, [2]int{i, j})

			if landed {
				landedCount++
				fmt.Println("landedCount", landedCount)

				if _minY < minY {
					minY = _minY
					fmt.Printf("ogv: %v, velocity: %v, curr: %v, landed: %v, minY: %v\n", [2]int{i, j}, velocity, curr, landed, minY)
				}
			}
		}
	}

	fmt.Println("minY", minY)
}

func run(x [2]int, y [2]int, velocity [2]int) (landed bool, curr [2]int, velocityOut [2]int, minY int) {
	velocityOut = velocity
	for curr = [2]int{0, 0}; !landed && curr[0] <= x[1] && curr[1] <= y[1]; curr, velocityOut = step(curr, velocityOut) {
		landed = curr[0] >= x[0] && curr[0] <= x[1] && curr[1] >= y[0] && curr[1] <= y[1]
		if curr[1] < minY {
			minY = curr[1]
		}
		// fmt.Printf("velocity: %v, curr: %v, landed: %v\n", velocity, curr, landed)
	}
	return
}

func RunB() {
}

func step(curr [2]int, velocityIn [2]int) (next [2]int, velocityOut [2]int) {
	next = [2]int{curr[0] + velocityIn[0], curr[1] - velocityIn[1]}
	velocityOut[0] = velocityIn[0]
	velocityOut[1] = velocityIn[1] - 1
	if velocityIn[0] < 0 {
		velocityOut[0] = velocityIn[0] + 1
	} else if velocityIn[0] > 0 {
		velocityOut[0] = velocityIn[0] - 1
	}

	return
}

func read() ([2]int, [2]int, [][]string) {
	var err error

	data, err := utils.ReadFile("./17/data.txt")
	handle(err)

	re := regexp.MustCompile(`target area: x=(\d+)\.\.(\d+), y=-(\d+)\.\.-(\d+)`)
	matches := re.FindStringSubmatch(data)

	x1, x2, y2, y1 := parseInt(matches[1]), parseInt(matches[2]), parseInt(matches[3]), parseInt(matches[4])

	grid := make([][]string, y2)
	for i := range grid {
		grid[i] = make([]string, x2)
		for j := range grid[i] {
			if i >= y1 && i <= y2 && j >= x1 && j <= x2 {
				grid[i][j] = "T"
			} else {
				grid[i][j] = "."
			}
		}
	}
	return [2]int{x1, x2}, [2]int{y1, y2}, grid
}

func parseInt(match string) int {
	num, err := strconv.Atoi(match)
	handle(err)
	return num
}
