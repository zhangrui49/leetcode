package main

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	t.Log(removeDuplicates([]int{1, 1, 2, 2, 2, 3}))
}
