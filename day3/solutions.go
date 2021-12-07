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

func Solution2() {
	lines := getLines()

	oxygen := oxygenGenRating(lines)
	co2 := co2ScrubberRating(lines)

	oxygenInt, _ := strconv.ParseInt(string(oxygen), 2, 64)
	co2Int, _ := strconv.ParseInt(string(co2), 2, 64)

	fmt.Println("What is the life support rating of the submarine?")
	fmt.Printf("Oxygen: %s %d  CO2: %s %d\n", string(oxygen), oxygenInt, string(co2),  co2Int)
	fmt.Printf("%d\n", oxygenInt * co2Int)
}

func oxygenGenRating(lines []line) line {
	return filterLines(lines, 0, mostCommonValueAt)
}

func co2ScrubberRating(lines []line) line {
	return filterLines(lines, 0, leastCommonValueAt)
}

func filterLines(lines []line, index int, getFilterValue func(lines []line, index int) rune) line {
	filteredLines := []line{}
	filterValue := getFilterValue(lines, index)

	for _, l := range lines {
		if l[index] == filterValue {
			filteredLines = append(filteredLines, l)
		}
	}

	if len(filteredLines) == 1 {
		return filteredLines[0]
	} else {
		return filterLines(filteredLines, index + 1, getFilterValue)
	}
}

func mostCommonValueAt(lines []line, index int) rune {
	oneCount := 0
	zeroCount := 0
	for _, r := range lines {
		if r[index] == '1' {
			oneCount++
		} else {
			zeroCount++
		}
	}

	if oneCount >= zeroCount {
		return '1'
	} else {
		return '0'
	}
}

func leastCommonValueAt(lines []line, index int) rune {
	oneCount := 0
	zeroCount := 0
	for _, r := range lines {
		if r[index] == '1' {
			oneCount++
		} else {
			zeroCount++
		}
	}

	if oneCount < zeroCount {
		return '1'
	} else {
		return '0'
	}
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