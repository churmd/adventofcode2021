package day3

import (
	"fmt"
	"strconv"
	"strings"
)

type line = []rune

func Solution1() {
	lines := getLines()
	totalNums := len(lines)
	occurancesOfOne := numOfOnes(lines)

	gamma := make(line, len(lines[0]))
	for i, countOfOnes := range occurancesOfOne {
		if countOfOnes > totalNums / 2 {
			gamma[i] = '1'
		} else {
			gamma[i] = '0'
		}
	}

	epsilon := make(line, len(lines[0]))
	for i, countOfOnes := range occurancesOfOne {
		if countOfOnes > totalNums / 2 {
			epsilon[i] = '0'
		} else {
			epsilon[i] = '1'
		}
	}

	gammaInt, _ := strconv.ParseInt(string(gamma), 2, 64)
	epsilonInt, _ := strconv.ParseInt(string(epsilon), 2, 64)

	fmt.Println("What is the power consumption of the submarine?")
	fmt.Printf("Gamma: %d  Epsilon: %d\n", gammaInt, epsilonInt)
	fmt.Printf("%d\n", gammaInt * epsilonInt)
}

func numOfOnes(lines []line) []int {
	counts := make([]int, len(lines[0]))

	for _, l := range lines {
		for i, value := range l {
			if value == '1' {
				counts[i]++
			}
		}
	}

	return counts
}

func getLines() []line { 
	newLines := strings.Split(input, "\n")

	lines := make([]line, len(newLines))  
	for i, l := range newLines {
		lines[i] = line(l)
	}

	return lines
}