package sort

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "Test case 1",
			arr:  []int{4, 2, 1, 3},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "Test case 2",
			arr:  []int{5, 8, 6, 9, 7},
			want: []int{5, 6, 7, 8, 9},
		},
		{
			name: "Test case 3",
			arr:  []int{10, 20, 30, 40, 50},
			want: []int{10, 20, 30, 40, 50},
		},
		{
			name: "Test case 4",
			arr:  []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSort(tt.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
