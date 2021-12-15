package day4

import (
	"fmt"
	"strconv"
	"strings"
)

type result struct {
	Sum      int
	LastBall int
	NumTurns int
}

func Solution1() {
	balls, boards := getBallsAndBoards()

	var best result
	for _, board := range boards {
		r, won := isWinner(board, balls)
		if won {
			if best.NumTurns == 0 {
				best = r
			} else if best.NumTurns > r.NumTurns {
				best = r
			}
		}
	}

	fmt.Printf("First Winner %v\n", best)
	fmt.Println("What will your final score be if you choose that board?")
	fmt.Printf("%d\n", best.Sum * best.LastBall)
}

func Solution2() {
	balls, boards := getBallsAndBoards()

	var worst result
	for _, board := range boards {
		r, won := isWinner(board, balls)
		if won {
			if r.NumTurns > worst.NumTurns {
				worst = r
			}
		}
	}

	fmt.Printf("Last Winner %v\n", worst)
	fmt.Println(" Once it wins, what would its final score be?")
	fmt.Printf("%d\n", worst.Sum * worst.LastBall)
}

func isWinner(board Board, balls []int) (result, bool) {
	for i, ball := range balls {
		board.Draw(ball)
		if sum, won := board.Won(); won {
			return result{
				Sum:      sum,
				LastBall: ball,
				NumTurns: i,
			}, true
		}
	}

	return result{}, false
}

func getBallsAndBoards() ([]int, []Board) {
	ballAndBoards := strings.SplitN(input, "\n", 3)
	ballsInput := ballAndBoards[0]
	boardsInput := ballAndBoards[2]

	var balls []int
	for _, ballInput := range strings.Split(ballsInput, ",") {
		ball, err := strconv.Atoi(ballInput)
		if err != nil {
			panic("ball was not a number " + ballInput)
		}
		balls = append(balls, ball)
	}

	var boards []Board
	for _, b := range strings.Split(boardsInput, "\n\n") {
		boards = append(boards, NewBoard(b))
	}

	return balls, boards
}
