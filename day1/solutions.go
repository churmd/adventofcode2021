package day1

import (
	"fmt"
	"strconv"
	"strings"
)

func Solution1() {
	depths := getDepths()
	increases := 0

	for i := 0; i < len(depths)-1; i++ {
		change := depths[i+1] - depths[i]
		if change > 0 {
			increases++
		}
	}

	fmt.Println("How many measurements are larger than the previous measurement?")
	fmt.Printf("%d\n", increases)
}

func Solution2() {
	depths := getDepths()
	windows := make([]int, len(depths)-2)

	for i := 0; i < len(windows); i++ {
		windows[i] = depths[i] + depths[i+1] + depths[i+2]
	}

	increases := countIncreases(windows)

	fmt.Println("How many sums are larger than the previous sum?")
	fmt.Printf("%d\n", increases)
}

func countIncreases(measurements []int) int {
	increases := 0
	for i := 0; i < len(measurements)-1; i++ {
		change := measurements[i+1] - measurements[i]
		if change > 0 {
			increases++
		}
	}
	return increases
}

func getDepths() []int {
	lines := strings.Split(problem1input, "\n")

	depths := make([]int, len(lines))
	for i, line := range lines {
		depth, err := strconv.Atoi(line)
		if err != nil {
			panic("could not convert a depth string to int: " + line)
		}

		depths[i] = depth
	}
	return depths
}
