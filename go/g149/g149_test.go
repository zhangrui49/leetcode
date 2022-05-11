package main

import (
	"testing"
)

func TestMaxPoints(t *testing.T) {

	//t.Log(maxPoints([][]int{
	//	{1, 1}, {3, 2},
	//	{5, 3}, {4, 1},
	//	{2, 3}, {1, 4},
	//}))
	//t.Log(maxPoints([][]int{
	//	{1, 1}, {2, 2},
	//	{3, 3},
	//}))
	t.Log(maxPoints([][]int{
		{0, 0}, {1, 0},
		{2, 0}, {3, 0},
	}))
}
