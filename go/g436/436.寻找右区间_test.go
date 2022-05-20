package g436

import (
	"testing"
)

func Test_findRightInterval(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{

				intervals: [][]int{{3, 4}, {2, 3}, {1, 2}},
			},
			want: nil,
		},
		{
			name: "test2",
			args: args{

				intervals: [][]int{{1, 4}, {2, 3}, {3, 4}},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findRightIntervalV3(tt.args.intervals)
			t.Errorf("findRightInterval() = %v, want %v", got, tt.want)

		})
	}
}
