package main

import "fmt"

var (
	draughtMoves = [][]int{
		{-3, 0},
		{0, -3},
		{3, 0},
		{0, 3},
		{2, 2},
		{2, -2},
		{-2, 2},
		{-2, -2},
	}
)

type moves [][]int

func (m *moves) insert(move []int) {
	*m = append(*m, move)
}

func (m *moves) pop() []int {
	l := len(*m)
	move := (*m)[l-1]
	*m = (*m)[:l-1]
	return move
}

func (m *moves) printMoves() {
	var board [10][10]int
	var prev []int

	for _, move := range *m {
		board[move[0]][move[1]] = 3
		if prev != nil {
			board[prev[0]][prev[1]] = 2
		}

		printBoard(board)

		board[move[0]][move[1]] = 1
		if prev != nil {
			board[prev[0]][prev[1]] = 1
		}

		prev = move
	}
}

func printBoard(board [10][10]int) {
	for _, row := range board {
		for _, col := range row {
			switch col {
			case 3:
				fmt.Printf("\033[1;31m%s\033[0m", "[x]")
			case 2:
				fmt.Printf("\033[1;34m%s\033[0m", "[x]")
			case 1:
				fmt.Print("[x]")
			default:
				fmt.Print("[ ]")
			}
		}
		fmt.Println()
	}

	fmt.Println("-----------------------")
}

type draughts struct {
	board   [10][10]bool
	counter int
}

func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			d := new(draughts)
			result := new(moves)
			result.insert([]int{i, j})
			if solve(i, j, d, result) {
				fmt.Println("The result:")
				result.printMoves()

				return
			}
		}
	}

	fmt.Println("No solution")
}

func solve(i, j int, draught *draughts, result *moves) bool {
	draught.board[i][j] = true
	draught.counter++
	if draught.counter == 100 {
		return true
	}

	for _, move := range draughtMoves {
		newI, newJ := move[0]+i, move[1]+j

		if !pointsAreValid(newI, newJ) || draught.board[newI][newJ] {
			continue
		}

		result.insert([]int{newI, newJ})
		if solve(newI, newJ, draught, result) {
			return true
		}

		result.pop()
		draught.board[newI][newJ] = false
		draught.counter--
	}

	return false
}

func pointsAreValid(i, j int) bool {
	return i >= 0 && i < 10 && j >= 0 && j < 10
}

func allCovered(board [10][10]bool) bool {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if !board[i][j] {
				return false
			}
		}
	}

	return true
}
