package g502

import "testing"

func TestFindMaximizedCapital(t *testing.T) {
	profits := []int{1, 2, 3}
	capitals := []int{0, 1, 2}

	t.Log(findMaximizedCapital(3, 0, profits, capitals))
}
