package main

import (
	"testing"
)

func TestIsValidSudoku(t *testing.T) {
	t.Log(IsValidSudoku([][]byte{
		[]byte("534678912"),
		[]byte("672195348"),
		[]byte("198342567"),
		[]byte("859761423"),
		[]byte("426853791"),
		[]byte("713924856"),
		[]byte("961537284"),
		[]byte("287419635"),
		[]byte("345286179"),
	}))
}
