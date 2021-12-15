package day4_test

import (
	"testing"

	"adventofcode2021/day4"

	"github.com/stretchr/testify/assert"
)

var (
	expectedCells = []day4.Cell{
		{Coord:  day4.Coord{X: 0, Y: 0}, Value: 17},
		{Coord:  day4.Coord{X: 1, Y: 0}, Value: 45},
		{Coord:  day4.Coord{X: 2, Y: 0}, Value: 62},
		{Coord:  day4.Coord{X: 3, Y: 0}, Value: 28},
		{Coord:  day4.Coord{X: 4, Y: 0}, Value: 73},
		{Coord:  day4.Coord{X: 0, Y: 1}, Value: 39},
		{Coord:  day4.Coord{X: 1, Y: 1}, Value: 12},
		{Coord:  day4.Coord{X: 2, Y: 1}, Value: 0},
		{Coord:  day4.Coord{X: 3, Y: 1}, Value: 52},
		{Coord:  day4.Coord{X: 4, Y: 1}, Value: 5},
		{Coord:  day4.Coord{X: 0, Y: 2}, Value: 87},
		{Coord:  day4.Coord{X: 1, Y: 2}, Value: 48},
		{Coord:  day4.Coord{X: 2, Y: 2}, Value: 50},
		{Coord:  day4.Coord{X: 3, Y: 2}, Value: 85},
		{Coord:  day4.Coord{X: 4, Y: 2}, Value: 44},
		{Coord:  day4.Coord{X: 0, Y: 3}, Value: 66},
		{Coord:  day4.Coord{X: 1, Y: 3}, Value: 57},
		{Coord:  day4.Coord{X: 2, Y: 3}, Value: 78},
		{Coord:  day4.Coord{X: 3, Y: 3}, Value: 94},
		{Coord:  day4.Coord{X: 4, Y: 3}, Value: 3},
		{Coord:  day4.Coord{X: 0, Y: 4}, Value: 91},
		{Coord:  day4.Coord{X: 1, Y: 4}, Value: 37},
		{Coord:  day4.Coord{X: 2, Y: 4}, Value: 69},
		{Coord:  day4.Coord{X: 3, Y: 4}, Value: 16},
		{Coord:  day4.Coord{X: 4, Y: 4}, Value: 1},
	}
	expectedBoard = day4.Board{
		Cells: expectedCells,
	}
)
func TestNewBoard(t *testing.T) {
	input :=
		`17 45 62 28 73
		39 12  0 52  5
		87 48 50 85 44
		66 57 78 94  3
		91 37 69 16  1`

	actualBoard := day4.NewBoard(input)

	assert.Equal(t, expectedBoard, actualBoard)
}

func TestWinningRow(t *testing.T) {
	board := expectedBoard
	board.Draw(17)
	board.Draw(45)
	board.Draw(62)
	board.Draw(28)
	board.Draw(73)

	sum, won := board.Won()

	assert.True(t, won)
	assert.Equal(t, 934, sum)
}
