package day5

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Coord struct {
	X int
	Y int
}

func (c *Coord) UnmarshalText(text []byte) error {
	values := strings.SplitN(string(text), ",", 2)

	x, err := strconv.Atoi(values[0])
	if err != nil {
		return errors.New("x value for coordinate not a number " + string(text))
	}

	y, err := strconv.Atoi(values[1])
	if err != nil {
		return errors.New("y value for coordinate not a number " + string(text))
	}

	c.X = x
	c.Y = y
	return nil
}

type Line struct {
	Start Coord
	End   Coord
}

func (l Line) isHorzOrVert() bool {
	if l.Start.X != l.End.X && l.Start.Y != l.End.Y {
		return false
	}

	return true
}

func (l Line) getAllCoords() []Coord {
	var coords []Coord

	xChange := 0
	if l.Start.X < l.End.X {
		xChange = 1
	}
	if l.Start.X > l.End.X {
		xChange = -1
	}

	yChange := 0
	if l.Start.Y < l.End.Y {
		yChange = 1
	}
	if l.Start.Y > l.End.Y {
		yChange = -1
	}

	current := l.Start
	for current != l.End {
		coords = append(coords, current)
		current.X += xChange
		current.Y += yChange
	}

	coords = append(coords, l.End)
	return coords
}

type Floor = map[Coord]int

func Solution1() {
	lines := getLines()

	var straightLines []Line
	for _, line := range lines {
		if line.isHorzOrVert() {
			straightLines = append(straightLines, line)
		}
	}

	floor := Floor{}

	for _, line := range straightLines {
		allCoords := line.getAllCoords()
		for _, coord := range allCoords {
			floor[coord] += 1
		}
	}

	atLeast2LinesCount := 0
	for _, count := range floor {
		if count >= 2 {
			atLeast2LinesCount += 1
		}
	}

	fmt.Println("At how many points do at least two lines overlap?")
	fmt.Printf("%d\n", atLeast2LinesCount)
}

func getLines() []Line{
	inputLines := strings.Split(input, "\n")

	var lines []Line
	for _, line := range inputLines {
		coords := strings.SplitN(line, " -> ", 2)
		start := Coord{}
		err := start.UnmarshalText([]byte(coords[0]))
		if err != nil {
			panic(err)
		}

		end := Coord{}
		err = end.UnmarshalText([]byte(coords[1]))
		if err != nil {
			panic(err)
		}

		lines = append(lines, Line{Start: start, End: end})
	}

	return lines
}
