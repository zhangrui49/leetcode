package g046

import (
	"testing"
)

func Test_permute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permute(tt.args.nums)
			t.Errorf("permute() = %v, want %v", got, tt.want)
		})
	}
}
