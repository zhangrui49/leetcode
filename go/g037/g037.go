package main

import (
	"fmt"
	"strconv"
)

func solveSudoku(board [][]byte) {
	suduku(board, 0, 0)
	for k := 0; k < 9; k++ {
		fmt.Println(string(board[k]))
	}
}

func suduku(board [][]byte, i int, j int) bool {

	if j >= 9 {
		for k := 0; k < i; k++ {
			fmt.Println(strconv.Itoa(k) + ":" + string(board[k]))
		}
		fmt.Println("")
		fmt.Println("")
		return suduku(board, i+1, 0)
	}
	if i >= 9 {
		return true
	}
	ch := board[i][j]
	//fmt.Println("i is " + strconv.Itoa(i) + "  j is  " + strconv.Itoa(j) + "  ch is: " + string(ch))
	if ch != '.' {
		return suduku(board, i, j+1)
	} else {
		for k := 1; k <= 9; k++ {
			board[i][j] = byte(k + '0')
			//fmt.Println("ch new " + string(board[i][j]))
			if isValidSudoku(i, j, board) {
				if suduku(board, i, j+1) {
					return true
				}
			}
			board[i][j] = '.'
		}
	}
	return false
}

func isValidSudoku(i int, j int, board [][]byte) bool {

	for k := 1; k < 9; k++ {
		if k != j && board[i][k] == board[i][j] {
			return false
		}
	}
	for k := 1; k < 9; k++ {
		if k != j && board[i][k] == board[i][j] {
			return false
		}
	}
	for row := i / 3 * 3; row < i/3*3+3; row++ {
		for col := j / 3 * 3; col < j/3*3+3; col++ {
			if (row != i || j != col) && board[row][col] == board[i][j] {
				return false
			}
		}
	}
	return true
}
