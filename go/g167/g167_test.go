package g167

import (
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				numbers: []int{2, 7, 11, 15},
				target:  9,
			},
			want: nil,
		},
		{
			name: "test2",
			args: args{
				numbers: []int{2, 3, 4},
				target:  6,
			},
			want: nil,
		},
		{
			name: "test3",
			args: args{
				numbers: []int{-1, 0},
				target:  -1,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSumV4(tt.args.numbers, tt.args.target)
			t.Errorf("twoSum() = %v, want %v", got, tt.want)

		})
	}
}
