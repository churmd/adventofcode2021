package day4_test

import (
	"adventofcode2021/day4"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBoard(t *testing.T) {
	input :=
		`17 45 62 28 73
		39 12  0 52  5
		87 48 50 85 44
		66 57 78 94  3
		91 37 69 16  1`

	expectedValues := [5][5]string{
		{"17", "45", "62", "28", "73"},
		{"39", "12", "0", "52", "5"},
		{"87", "48", "50", "85", "44"},
		{"66", "57", "78", "94", "3"},
		{"91", "37", "69", "16", "1"},
	}
	expectedBoard := day4.Board{Values: expectedValues}

	actualBoard := day4.NewBoard(input)

	assert.Equal(t, expectedBoard, actualBoard)
}
