package day1

import (
	"fmt"
	"strconv"
	"strings"
)

func Solution1() {
	lines := strings.Split(problem1input, "\n")

	depths := make([]int, len(lines))
	for i, line := range lines {
		depth, err := strconv.Atoi(line)
		if err != nil {
			panic("could not convert a depth string to int: " + line)
		}

		depths[i] = depth
	}

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
