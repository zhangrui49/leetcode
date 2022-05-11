package main

import (
	"testing"
	//245137689
)

func TestSolveSudoku(t *testing.T) {
	solveSudoku([][]byte{
		[]byte("187654329"),
		[]byte("2........"),
		[]byte("3........"),
		[]byte("4........"),
		[]byte("5........"),
		[]byte("6........"),
		[]byte("7........"),
		[]byte("8........"),
		[]byte("9........"),
	})
}
