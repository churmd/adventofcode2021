package day4

import "strings"

type Board struct {
	Values [5][5]string
	Marked [5][5]bool
}

func NewBoard(input string) Board {
	cells := [5][5]string{}
	marked := [5][5]bool{}
	// for i := range cells {
	// 	cells[i] = make([]string, 5)
	// 	marked[i] = make([]bool, 5)
	// }

	lines := strings.Split(input, "\n")

	for i := 0; i < 5; i++ {
		values := strings.Fields(lines[i])

		for j := 0; j < 5; j++ {
			cells[i][j] = values[j]
		}
	}

	return Board{Values: cells, Marked: marked}
}
