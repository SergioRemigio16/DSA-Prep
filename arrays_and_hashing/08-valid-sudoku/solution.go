package main

import (
	"math"
)

func isValidSudoku(board [][]byte) bool {
	n := len(board)

	rows := make([]map[byte]int, n)
	for i := 0; i < n; i++ {
		rows[i] = make(map[byte]int)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sudokuSquare := board[i][j]
			if sudokuSquare == '.' {
				continue
			}
			rows[i][sudokuSquare]++
			if rows[i][sudokuSquare] > 1 {
				return false
			}
		}
	}

	columns := make([]map[byte]int, n)
	for i := 0; i < n; i++ {
		columns[i] = make(map[byte]int)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sudokuSquare := board[j][i]
			if sudokuSquare == '.' {
				continue
			}
			columns[i][sudokuSquare]++
			if columns[i][sudokuSquare] > 1 {
				return false
			}
		}
	}

	squareLength := int(math.Round(math.Sqrt(float64(n))))
	squares := make([]map[byte]int, n)
	for i := 0; i < n; i++ {
		squares[i] = make(map[byte]int)
	}
	row := 0
	for i := 0; i < n; i = i + squareLength {
		for j := 0; j < n; j = j + squareLength {
			for k := i; k < i+squareLength; k++ {
				for l := j; l < j+squareLength; l++ {
					sudokuSquare := board[k][l]
					sudokuSquareAsChar := string(board[k][l])
					_ = sudokuSquareAsChar
					if sudokuSquare == '.' {
						continue
					}
					squares[row][sudokuSquare]++
					if squares[row][sudokuSquare] > 1 {
						return false
					}
				}
			}
			row++
		}
	}

	return true
}

func main() {
	board := [][]byte{
		{'1', '2', '.', '.', '3', '.', '.', '.', '.'},
		{'4', '.', '.', '5', '.', '.', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '.', '3'},
		{'5', '.', '.', '.', '6', '.', '.', '.', '4'},
		{'.', '.', '.', '8', '.', '3', '.', '.', '5'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '.', '.', '.', '.', '.', '2', '.', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '8'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	boardTwo := [][]byte{

		{'1', '2', '.', '.', '3', '.', '.', '.', '.'},
		{'4', '.', '.', '5', '.', '.', '.', '.', '.'},
		{'.', '9', '1', '.', '.', '.', '.', '.', '3'},
		{'5', '.', '.', '.', '6', '.', '.', '.', '4'},
		{'.', '.', '.', '8', '.', '3', '.', '.', '5'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '.', '.', '.', '.', '.', '2', '.', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '8'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	println(isValidSudoku(board))
	println(isValidSudoku(boardTwo))
}
