package _go

import (
	"testing"
)

func TestTrap(t *testing.T) {

	t.Log(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	t.Log(trap([]int{4, 2, 3}))
	t.Log(trap([]int{5, 2, 1, 2, 1, 5}))
	//t.Log(trap([]int{3, 2, 1, 2, 1}))
}
