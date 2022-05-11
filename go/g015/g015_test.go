package main

import (
	"testing"
)

func TestThreeSum(t *testing.T) {
	// -4 -1 -1 0 1 2
	t.Log(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	t.Log(threeSum([]int{0, 0, 0, 0}))
	// -1 0 0 1
	t.Log(threeSum([]int{-1, 0, 1, 0}))

}
