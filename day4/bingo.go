package day4

import (
	"strconv"
	"strings"
)

type Coord struct {
	X int
	Y int
}

type Cell struct {
	Coord Coord
	Value  int
	Marked bool
}

type Board struct {
	Cells []Cell
}

func NewBoard(input string) Board {
	var cells []Cell

	lines := strings.Split(input, "\n")

	for y := 0; y < 5; y++ {
		values := strings.Fields(lines[y])

		for x := 0; x < 5; x++ {
			value, err := strconv.Atoi(values[x])
			if err != nil {
				panic("could not convert bingo board value to int " + values[x])
			}
			cells = append(cells, Cell{
				Coord: Coord{
					X: x,
					Y: y,
				},
				Value:  value,
				Marked: false,
			})
		}
	}

	return Board{Cells: cells}
}

func (b *Board) Draw(ball int) {
	for i, cell := range b.Cells {
		if cell.Value == ball {
			cell.Marked = true
		}
		b.Cells[i] = cell
	}
}

func (b Board) Won() (int, bool) {
	for y := 0; y < 5; y++ {
		var row []Cell
		for _, cell := range b.Cells {
			if cell.Coord.Y == y {
				row = append(row, cell)
			}
		}

		if b.winningLine(row) {
			return b.sumOfUnmarked(), true
		}
	}

	for x := 0; x < 5; x++ {
		var column []Cell
		for _, cell := range b.Cells {
			if cell.Coord.X == x {
				column = append(column, cell)
			}
		}

		if b.winningLine(column) {
			return b.sumOfUnmarked(), true
		}
	}

	return 0, false
}

func (b Board) winningLine(cells []Cell) bool {
	for _, cell := range cells {
		if !cell.Marked {
			return false
		}
	}

	return true
}

func (b Board) sumOfUnmarked()  int {
	sum := 0
	for _, cell := range b.Cells {
		if !cell.Marked {
			sum += cell.Value
		}
	}
	return sum
}
